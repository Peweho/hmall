package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/address/rpc/address"
	"hmall/application/item/rpc/item"
	"hmall/application/order/api/internal/model"
	"hmall/pkg/util"
	"strconv"
	"sync"

	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrdeReq) error {
	var (
		itemResp    *item.FindItemByIdsResp
		addressResp *address.FindAdressByIdResp
		usr         int
		err1        error
		err2        error
		err3        error
		err4        error
		err5        error
	)
	itemIds := make([]string, 0, len(req.Details))
	wg := &sync.WaitGroup{}
	// 1、查询地址
	wg.Add(3)
	threading.GoSafe(func() {
		defer wg.Done()
		addressResp, err1 = l.svcCtx.AddressRPC.FindAdressById(l.ctx, &address.FindAdressByIdReq{Id: int64(req.AddressId)})
		if err1 != nil {
			logx.Errorf("AddressRPC.FindAdressById: %v, error: %v", req.AddressId, err1)
		}
	})
	//2、查询商品
	threading.GoSafe(func() {
		defer wg.Done()
		for _, val := range req.Details {
			itemIds = append(itemIds, strconv.Itoa(val.ItemId))
		}
		itemResp, err2 = l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: itemIds})
		if err2 != nil {
			logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", itemIds, err2)
		}
	})
	//3、获得用户
	threading.GoSafe(func() {
		defer wg.Done()
		usr, err3 = util.GetUsr(l.ctx, types.JwtKey)
		if err3 != nil {
			logx.Errorf("util.GetUsr: %v, error: %v", itemIds, err3)
		}
	})
	wg.Wait()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if err3 != nil {
		return err3
	}

	//3、写入订单表 并 获得Id
	//计算总金额
	num := map[int64]int{}
	for _, val := range req.Details {
		num[int64(val.ItemId)] = val.Num
	}
	order := &model.OrderPO{
		Total_fee:    TotalFee(itemResp.Data, num),
		User_id:      int64(usr),
		Status:       types.NotPayment,
		Payment_type: req.PaymentType,
	}
	if err := l.svcCtx.OrderModel.AddOrder(l.ctx, order); err != nil {
		logx.Errorf("OrderModel.AddOrder: %v, error: %v", order, err)
		return err
	}

	wg.Add(2)
	//5、写入订单发货表
	threading.GoSafe(func() {
		defer wg.Done()
		err5 = l.WriteOrderLogistics(addressResp, order.Id)
	})

	//4、写入订单详情表

	threading.GoSafe(func() {
		defer wg.Done()
		err4 = l.WriteOrderDetail(itemResp.Data, num, order.Id)
	})
	wg.Wait()

	if err4 != nil {
		return err4
	}
	if err5 != nil {
		return err5
	}

	return nil
}

func TotalFee(items []*item.Items, num map[int64]int) int {
	var fee int
	for _, val := range items {
		fee += num[val.Id] * int(val.Price)
	}
	return fee
}

// 写入订单详情表
func (l *CreateOrderLogic) WriteOrderDetail(items []*item.Items, num map[int64]int, orderId int64) error {
	orderDetail := make([]map[string]any, 0, len(items))
	for _, val := range items {
		orderDetail = append(orderDetail, map[string]any{
			"order_id": orderId,
			"item_id":  val.Id,
			"num":      num[val.Id],
			"name":     val.Name,
			"spec":     val.Spec,
			"price":    int(val.Price),
			"image":    val.Image,
		})
	}
	if err := l.svcCtx.OrderModel.AddOrderDetail(l.ctx, orderDetail); err != nil {
		logx.Errorf("OrderModel.AddOrderDetail: %v, error: %v", orderDetail, err)
		return err
	}
	return nil
}

func (l *CreateOrderLogic) WriteOrderLogistics(addressResp *address.FindAdressByIdResp, orderId int64) error {
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
