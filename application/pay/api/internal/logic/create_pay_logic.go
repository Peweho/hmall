package logic

import (
	"context"
	"hmall/application/pay/api/internal/model"
	"hmall/pkg/util"
	"time"

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
	uid, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Error("util.GetUsr: %v, error: %v", uid, err)
		return err
	}
	// 构建对象
	pay := &model.PayPO{
		Biz_order_no:     int64(req.BizOrderNo),
		Biz_user_id:      int64(uid),
		Pay_channel_code: req.PayChannelCode,
		Amount:           req.Amount,
		Pay_type:         req.PayType,
		Status:           types.NotPay,
		Pay_success_time: nil,
		Pay_over_time:    time.Unix(time.Now().Unix()+types.OverTime, 0),
		Is_delete:        types.NotDelete,
	}

	//创建
	if err = l.svcCtx.PayModel.CreatePayOrder(l.ctx, pay); err != nil {
		logx.Error("PayModel.CreatePayOrder: %v, error : %v", pay, err)
		return err
	}

	return nil
}
