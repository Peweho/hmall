package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/order/rpc/order"
	"hmall/application/pay/api/internal/svc"
	"hmall/application/pay/api/internal/types"
	"hmall/application/pay/rpc/pay"
	"hmall/application/user/rpc/user"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
)

type UseBalancePayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseBalancePayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseBalancePayLogic {
	return &UseBalancePayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseBalancePayLogic) UseBalancePay(req *types.UseBalancePayReq) error {
	dtmServer := "etcd://192.168.92.201:2379/dtmservice"
	// 1、判断订单是否可支付（状态，是否删除）
	res, err := l.svcCtx.PayModel.SelPayOrderStatusIsDel(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("PayModel.SelPayOrderStatusIsDel: %v ,error: %v", req.Id, err)
		return err
	}
	if res.Status != types.NotPay || string(res.Is_delete) != types.NotDelete {
		return xcode.New(200, "订单无效")
	}

	uid, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}
	//2、构造事务请求
	//扣减余额
	userBuildTarget, err := l.svcCtx.Config.UserRPC.BuildTarget()
	if err != nil {
		logx.Errorf("UserRPC.BuildTarget, error: %v", err)
		return err
	}
	decutMoneyReq := user.DecutMoneyReq{
		Uid:    int64(uid),
		Amount: int64(res.Amount),
		Pwd:    req.Pw,
	}
	//更新订单状态
	orderBuildTarget, err := l.svcCtx.Config.OrderRPC.BuildTarget()
	if err != nil {
		logx.Errorf("OrderRPC.BuildTarget, error: %v", err)
		return err
	}
	updateOder := order.UpdateOrderStatusReq{
		Id: int64(req.Id),
	}
	//更新支付单流水
	payBuildTarget, err := l.svcCtx.Config.PayRPC.BuildTarget()
	if err != nil {
		logx.Errorf("Config.PayRPC, error: %v", err)
		return err
	}
	updatePay := pay.UpdatePayOrderReq{
		PayType: 5,
	}
	//3、开启分布式事务
	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(userBuildTarget+"/service.User/DecutMoney", userBuildTarget+"/service.User/DecutMoneyRollBack", &decutMoneyReq).
		Add(payBuildTarget+"/service.Pay/UpdatePayOrder", payBuildTarget+"/service.Pay/UpdatePayOrderRollBack", &updatePay).
		Add(orderBuildTarget+"/service.Order/UpdateOrderStatus", orderBuildTarget+"/service.Order/UpdateOrderStatusRollBack", &updateOder)
	err = saga.Submit()
	if err != nil {
		return err
	}

	return nil
}
