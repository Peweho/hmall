package logic

import (
	"context"

	"hmall/application/pay/api/internal/svc"
	"hmall/application/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return nil
}
