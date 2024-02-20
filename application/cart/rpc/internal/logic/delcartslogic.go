package logic

import (
	"context"

	"hmall/application/cart/rpc/internal/svc"
	"hmall/application/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartsLogic {
	return &DelCartsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartsLogic) DelCarts(in *pb.DelCartsReq) (*pb.DelCartsResp, error) {

	if err := l.svcCtx.CartModel.DelCartsByIds(l.ctx, in.Ids); err != nil {
		logx.Errorf("CartModel.DelCartById: %V, error: %v", in.Ids, err)
		return &pb.DelCartsResp{}, err
	}
	return &pb.DelCartsResp{}, nil
}
