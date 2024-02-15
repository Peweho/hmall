package logic

import (
	"context"

	"hmall/application/user/api/internal/svc"
	"hmall/application/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeductMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductMoneyLogic {
	return &DeductMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeductMoneyLogic) DeductMoney(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
