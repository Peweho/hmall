package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/cart/rpc/carts"
	"hmall/application/item/rpc/item"
	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"
	"hmall/application/order/rpc/order"
	"hmall/pkg/util"
	"strconv"
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
	dtmServer := "etcd://192.168.92.201:2379/dtmservice"
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}

	OrderRPCBuildTarget, err := l.svcCtx.Config.OrderRPC.BuildTarget()
	CartRPCBuildTarget, err := l.svcCtx.Config.CartRPC.BuildTarget()
	ItemRPCBuildTarget, err := l.svcCtx.Config.ItemRPC.BuildTarget()

	if err != nil {
		logx.Errorf("Config.OrderRPC.BuildTarget, error: %v", err)
		return err
	}
	itemNum := len(req.Details)
	itemDetail := make([]*item.ItemDetail, 0, itemNum)
	//构建请求 and 构建商品集合
	orderDetail := make([]*order.DetailDTO, 0, itemNum)
	itemIds := make([]string, 0, itemNum)

	for _, val := range req.Details {
		orderDetail = append(orderDetail, &order.DetailDTO{
			ItemId: int64(val.ItemId),
			Num:    int64(val.Num),
		})
		itemIds = append(itemIds, strconv.Itoa(val.ItemId))

		itemDetail = append(itemDetail, &item.ItemDetail{
			ItemId: strconv.Itoa(val.ItemId),
			Num:    int64(val.Num),
		})
	}
	//构建order请求
	orderReq := order.CreateOrderReq{
		UserId:      int64(usr),
		AddressId:   int64(req.AddressId),
		PaymentType: int64(req.PaymentType),
		Details:     orderDetail,
	}
	//构建cart请求
	cartReq := carts.DelCartsReq{
		Usr:    int64(usr),
		ItemId: itemIds,
	}

	//构建Item请求
	itemReq := item.DelStockReq{
		Detail: itemDetail,
	}

	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(OrderRPCBuildTarget+"/service.Order/CreateOrder", OrderRPCBuildTarget+"/service.Order/CreateOrderRollBack", &orderReq).
		Add(CartRPCBuildTarget+"/service.Carts/DelCarts", CartRPCBuildTarget+"/service.Carts/DelCartsRollBack", &cartReq).
		Add(ItemRPCBuildTarget+"/service.Item/DelStock", ItemRPCBuildTarget+"/service.Item/DelStockRollBack", &itemReq)

	err = saga.Submit()
	if err != nil {
		return err
	}

	return nil
}
