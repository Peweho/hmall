package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/order/rpc/internal/utils"
	"log"
	"sync"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRollBackLogic {
	return &CreateOrderRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderRollBackLogic) CreateOrderRollBack(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	//barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	//db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	//if err != nil {
	//	//!!!一般数据库不会错误不需要dtm回滚，就让他一直重试，这时候就不要返回codes.Aborted, dtmcli.ResultFailure 就可以了，具体自己把控!!!
	//	return nil, status.Error(codes.Internal, err.Error())
	//}
	//if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {

	// 1、查询用户最新创建的订单id
	orderId, err := l.svcCtx.OrderModel.FindNewOrderIdByUser(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("OrderModel.FindNewOrderIdByUser: %v, error: %v", in.UserId, err)
		return nil, err
	}
	log.Println("orderId:%v", orderId)
	//2、删除三张表
	chErr := make(chan error, 4) //接受错误

	wg := &sync.WaitGroup{}
	wg.Add(4)
	// 删除订单表
	threading.GoSafe(func() {
		defer wg.Done()
		err1 := l.svcCtx.OrderModel.DelOrderById(l.ctx, orderId)
		if err1 != nil {
			logx.Errorf("OrderModel.DelOrderById: %v, error: %v", orderId, err1)
		}
		chErr <- err1
	})
	//删缓存
	threading.GoSafe(func() {
		defer wg.Done()
		_, _ = l.svcCtx.BizRedis.Del(utils.CacheKey(int(orderId)))
		chErr <- nil
	})
	// 删除订单详情表
	threading.GoSafe(func() {
		defer wg.Done()
		var itemIds []int64
		for _, v := range in.Details {
			itemIds = append(itemIds, v.ItemId)
		}
		err2 := l.svcCtx.OrderModel.DelOrderDetailById(l.ctx, orderId, itemIds)
		if err2 != nil {
			logx.Errorf("OrderModel.DelOrderDetailById: %v, error: %v", orderId, err2)
		}
		chErr <- err2
	})
	//删除订单物流表
	threading.GoSafe(func() {
		defer wg.Done()
		err3 := l.svcCtx.OrderModel.DelOrderLogisticById(l.ctx, orderId)
		if err3 != nil {
			logx.Errorf("OrderModel.DelOrderLogisticById: %v, error: %v", orderId, err3)
		}
		chErr <- err3
	})
	wg.Wait()
	close(chErr)
	for e := range chErr {
		if e != nil {
			return nil, e
		}
	}
	//}); err != nil {
	//	return nil, err
	//}
	return &pb.CreateOrderResp{}, nil
}
