package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/status"
	"hmall/application/address/rpc/address"
	"hmall/application/item/rpc/item"
	"hmall/application/order/api/internal/model"
	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"
	"hmall/pkg/util"
	"log"
	"strconv"
	"sync"
	"time"
)

type CreateFlashOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFlashOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFlashOrderLogic {
	return &CreateFlashOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFlashOrderLogic) CreateFlashOrder(req *types.CreateFlashOrdeReq) error {
	uid, _ := util.GetUsr(l.ctx, types.JwtKey)
	//检查是否已经参秒杀
	statusReq := &item.FlashUserStatusReq{
		Uid:    strconv.Itoa(uid),
		ItemId: req.ItemId,
	}
	userStatus, err := l.svcCtx.ItemRPC.FlashUserStatus(l.ctx, statusReq)
	if err != nil {
		logx.Errorf("ItemRPC.FlashUserStatus: %v, error: %v", statusReq, err)
	}
	if userStatus.Status == strconv.Itoa(types.FalshEndDecut) {
		return status.Error(200, "已经下单")
	}

	resp := status.Error(200, "下单失败")
	//构建扣减库存请求
	itemDelFlashReq := item.DelFlashItemStockReq{
		Uid:      int64(uid),
		ItemId:   req.ItemId,
		Num:      int64(req.Num),
		Duration: int64(req.Duration),
	}

	chFlash := make(chan error, 1)
	threading.GoSafe(func() {
		log.Println("调用rpc")
		_, err := l.svcCtx.ItemRPC.DelFlashItemStock(l.ctx, &itemDelFlashReq)
		log.Println("调用rpc结束")
		chFlash <- err
	})

	for {
		select {
		case err := <-chFlash:
			if err != nil {
				if !errors.Is(err, context.DeadlineExceeded) {
					//非超时错误，退出事务
					log.Println("退出事务1", err)
					return resp
				}
			} else {
				//创建订单
				err := l.cretaOrder(req.ItemId, uid, req.Num)
				return err
			}
		case <-time.After(time.Duration(l.svcCtx.Config.Timeout)):
			//超时检查
			userStatus, resp := l.svcCtx.ItemRPC.FlashUserStatus(l.ctx, &item.FlashUserStatusReq{
				Uid:    strconv.Itoa(uid),
				ItemId: req.ItemId,
			})
			if userStatus.Status == "" {
				log.Println("退出事务2")
				return resp
			}

			//根据状态，进行下一步处理
			statusCode, _ := strconv.Atoi(userStatus.Status)
			switch statusCode {
			//未开始，退出事务
			case types.FalshNotStart:
				log.Println("退出事务2")
				return resp
			//开始处理，继续等待
			case types.FalshStart:
				log.Println("循环等待")
			//扣减失败，退出事务
			case types.FalshEndNotDecut:
				log.Println("退出事务3")
				return resp
			//已经扣减，创建订单
			case types.FalshEndDecut:
				err := l.cretaOrder(req.ItemId, uid, req.Num)
				return err
			default:
				panic("unhandled default case")
			}
		}
	}
	return nil
}

func (l *CreateFlashOrderLogic) cretaOrder(itemId string, uid int, num int) error {
	log.Println("开始创建订单")
	//查询商品信息
	items, err := l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: []string{itemId}})
	if err != nil {
		logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", itemId, err)
		return err
	}
	//写入订单表，获得订单Id
	order := model.OrderPO{
		Total_fee:    int(items.Data[0].Price) * num,
		User_id:      int64(uid),
		Status:       types.NotPayment,
		Payment_type: 3,
	}
	err = l.svcCtx.OrderModel.AddOrder(l.ctx, &order)
	if err != nil {
		logx.Errorf("OrderModel.AddOrder: %v, error: %v", order, err)
		return err
	}

	//并发写入订单详情表和订单物流表，地址取用户第一条地址
	wg := &sync.WaitGroup{}
	wg.Add(2)
	chErr := make(chan error, 2)
	threading.GoSafe(func() {
		defer wg.Done()
		defaultAddress, err := l.svcCtx.AddressRPC.GetUserDefaultAddress(l.ctx, &address.GetUserDefaultAddressReq{Uid: int64(uid)})
		if err != nil {
			chErr <- err
			panic(err)
		}
		err = l.WriteOrderLogistics(defaultAddress, order.Id)
		chErr <- err
		panic(err)
	})

	threading.GoSafe(func() {
		defer wg.Done()
		err := l.WriteOrderDetail(items.Data[0], 1, order.Id)
		chErr <- err
		panic(err)
	})

	wg.Wait()
	close(chErr)

	for e := range chErr {
		if e != nil {
			return e
		}
	}
	return nil
}

func (l *CreateFlashOrderLogic) WriteOrderLogistics(addressResp *address.FindAdressByIdResp, orderId int64) error {
	orderLogisticsPO := map[string]any{
		"order_id": orderId,
		"mobile":   addressResp.Mobile,
		"province": addressResp.Province,
		"city":     addressResp.City,
		"town":     addressResp.Town,
		"street":   addressResp.Street,
		"contact":  addressResp.Contact,
	}
	if err := l.svcCtx.OrderModel.AddOrderLogistics(l.ctx, orderLogisticsPO); err != nil {
		logx.Errorf("OrderModel.AddOrderLogistics: %v, error: %v", orderLogisticsPO, err)
		return err
	}
	return nil
}

// 写入订单详情表
func (l *CreateFlashOrderLogic) WriteOrderDetail(val *item.Items, num int64, orderId int64) error {
	orderDetail := make([]map[string]any, 0, 1)
	orderDetail = append(orderDetail, map[string]any{
		"order_id": orderId,
		"item_id":  val.Id,
		"num":      num,
		"name":     val.Name,
		"spec":     val.Spec,
		"price":    int(val.Price),
		"image":    val.Image,
	})

	if err := l.svcCtx.OrderModel.AddOrderDetail(l.ctx, orderDetail); err != nil {
		logx.Errorf("OrderModel.AddOrderDetail: %v, error: %v", orderDetail, err)
		return err
	}
	return nil
}
