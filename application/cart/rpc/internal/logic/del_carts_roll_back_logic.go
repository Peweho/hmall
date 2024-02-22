package logic

import (
	"context"

	"hmall/application/cart/rpc/internal/svc"
	"hmall/application/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartsRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartsRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartsRollBackLogic {
	return &DelCartsRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartsRollBackLogic) DelCartsRollBack(in *pb.DelCartsReq) (*pb.DelCartsResp, error) {
	if err := l.svcCtx.CartModel.SetCartsByUidItemId(l.ctx, in.Usr, in.ItemId); err != nil {
		logx.Errorf("CartModel.DelCartsByIds: %v, error: %v", in.Usr, err)
		return nil, err
	}

	return &pb.DelCartsResp{}, nil
}
