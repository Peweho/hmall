package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/order/rpc/internal/utils"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	threading.NewWorkerGroup(func() {
		key := utils.CacheKey(int(in.Id))
		if _, err := l.svcCtx.BizRedis.Del(key); err != nil {
			logx.Errorf("BizRedis.Del: %v, error: %v", key, err)
		}
	}, 1)
	// 调用mq服务修改
	pusherLogic := NewPusherLogic(l.ctx, l.svcCtx)
	if err := pusherLogic.UpdateStatus(int(in.Id)); err != nil {
		return nil, err
	}
	return &pb.UpdateOrderStatusResp{}, nil
}
