package logic

import (
	"context"

	"hmall/application/pay/api/internal/svc"
	"hmall/application/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePayLogic {
	return &CreatePayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePayLogic) CreatePay(req *types.CreatePayReq) error {
	// todo: add your logic here and delete this line

	return nil
}
