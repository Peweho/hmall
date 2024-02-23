package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/address/rpc/address"
	"hmall/application/item/rpc/item"
	"hmall/application/order/rpc/internal/model"
	"hmall/application/order/rpc/internal/utils"
	"strconv"
	"sync"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	var (
		itemResp    *item.FindItemByIdsResp
		addressResp *address.FindAdressByIdResp
		usr         = in.UserId
		err1        error
		err2        error
	)
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		//!!!一般数据库不会错误不需要dtm回滚，就让他一直重试，这时候就不要返回codes.Aborted, dtmcli.ResultFailure 就可以了，具体自己把控!!!
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		itemIds := make([]string, 0, len(in.Details))
		wg := &sync.WaitGroup{}
		// 1、查询地址
		wg.Add(2)
		threading.GoSafe(func() {
			defer wg.Done()
			addressResp, err1 = l.svcCtx.AddressRPC.FindAdressById(l.ctx, &address.FindAdressByIdReq{Id: int64(in.AddressId)})
			if err1 != nil {
				logx.Errorf("AddressRPC.FindAdressById: %v, error: %v", in.AddressId, err1)
			}
		})
		//2、查询商品
		threading.GoSafe(func() {
			defer wg.Done()
			for _, val := range in.Details {
				itemIds = append(itemIds, strconv.Itoa(int(val.ItemId)))
			}
			itemResp, err2 = l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: itemIds})
			if err2 != nil {
				logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", itemIds, err2)
			}
		})

		wg.Wait()
		if err1 != nil || err2 != nil {
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		}

		//3、写入订单表 并 获得Id
		//计算总金额
		num := map[int64]int64{}
		for _, val := range in.Details {
			num[val.ItemId] = val.Num
		}
		order := &model.OrderPO{
			Total_fee:    TotalFee(itemResp.Data, num),
			User_id:      usr,
			Status:       utils.NotPayment,
			Payment_type: int(in.PaymentType),
		}
		if err := l.svcCtx.OrderModel.AddOrder(l.ctx, order); err != nil {
			logx.Errorf("OrderModel.AddOrder: %v, error: %v", order, err)
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		}

		wg.Add(3)
		threading.GoSafe(func() {
			//写缓存
			defer wg.Done()
			key := utils.CacheKey(int(order.Id))
			marshal, cacheErr := json.Marshal(order)
			if cacheErr != nil {
				logx.Errorf("json.Marshal: %v, error : %v", order, cacheErr)
			} else {
				_ = l.svcCtx.BizRedis.Set(key, string(marshal))
				_ = l.svcCtx.BizRedis.Expire(key, utils.CacheOrderTime)
			}
		})

		//4、写入订单发货表
		threading.GoSafe(func() {
			defer wg.Done()
			err1 = l.WriteOrderLogistics(addressResp, order.Id)
		})

		//5、写入订单详情表
		threading.GoSafe(func() {
			defer wg.Done()
			err2 = l.WriteOrderDetail(itemResp.Data, num, order.Id)
		})
		wg.Wait()
		if err1 != nil || err2 != nil {
			return dtmcli.ErrFailure
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	return nil, nil
}

func TotalFee(items []*item.Items, num map[int64]int64) int {
	var fee int64
	for _, val := range items {
		fee += num[val.Id] * val.Price
	}
	return int(fee)
}

// 写入订单详情表
func (l *CreateOrderLogic) WriteOrderDetail(items []*item.Items, num map[int64]int64, orderId int64) error {
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
