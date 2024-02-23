package logic

import (
	"context"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStockRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStockRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStockRollBackLogic {
	return &DelStockRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStockRollBackLogic) DelStockRollBack(in *pb.DelStockReq) (*pb.DelStockResp, error) {
	for _, val := range in.Detail {
		key := util.CacheKey(types.CacheItemStockKey, val.ItemId)
		exists, _ := l.svcCtx.BizRedis.Exists(key)
		if exists {
			get, _ := l.svcCtx.BizRedis.Get(key)
			cacheNum, _ := strconv.Atoi(get)
			_ = l.svcCtx.BizRedis.Set(key, strconv.Itoa(cacheNum+int(val.Num)))
			_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime)
			continue
		}
		err := l.svcCtx.ItemModel.AddStock(l.ctx, val.ItemId, val.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
			return nil, err
		}
	}
	return &pb.DelStockResp{}, nil
}
