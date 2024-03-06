package logic

import (
	"context"
	"fmt"
	"hmall/application/item/rpc/types"

	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashUserStatusLogic {
	return &FlashUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashUserStatusLogic) FlashUserStatus(in *pb.FlashUserStatusReq) (*pb.FlashUserStatusResp, error) {
	key := fmt.Sprintf("%s#%d#%s", types.CacheFlashStatus, in.Uid, in.ItemId)
	get, err := l.svcCtx.BizRedis.Get(key)
	if err != nil {
		logx.Errorf("BizRedis.Get: %v, error: %v", key, err)
		return nil, err
	}

	return &pb.FlashUserStatusResp{Status: get}, nil
}
