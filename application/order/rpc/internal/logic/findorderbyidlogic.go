package logic

import (
	"context"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderByIdLogic {
	return &FindOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOrderByIdLogic) FindOrderById(in *pb.FindOrderByIdReq) (*pb.FindOrderByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FindOrderByIdResp{}, nil
}
