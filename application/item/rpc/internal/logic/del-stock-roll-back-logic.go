package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"strconv"
	"sync"
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

var decut_rollback_lua = `
	local acquired = redis.call("exist",KEYS[2])
	if acquired == 1 then
		local stock = redis.call("HGET",KEYS[2],KEYS[3])
		redis.call("HSET",KEYS[2],KEYS[3],stock + ARGV[2])
		redis.call("EXPIRE",KEYS[2],ARGV[3])
	end
`

func (l *DelStockRollBackLogic) DelStockRollBack(in *pb.DelStockReq) (*pb.DelStockResp, error) {
	wg := &sync.WaitGroup{}
	for _, val := range in.Detail {
		item, err := l.svcCtx.ItemModel.AddStock(l.ctx, val.ItemId, val.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
			return nil, err
		}
		wg.Add(2)
		//写缓存
		threading.GoSafe(func() {
			defer wg.Done()
			key := util.CacheKey(types.CacheItemKey, strconv.FormatInt(item.Id, 10))
			err = l.UpdateCacheRollBack(val.ItemId, val.Num)
			if err != nil {
				logx.Errorf("UpdateCacheRollBack: %v, error: %v", key, err)
			}
		})

		//同步es
		threading.GoSafe(func() {
			defer wg.Done()
			pusherSearch := NewPusherSearchLogic(l.ctx, l.svcCtx)
			err = pusherSearch.PusherSearch(types.KqUpdate, item)
			if err != nil {
				logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", *item, err)
			}
		})
	}
	wg.Wait()
	return &pb.DelStockResp{}, nil
}

func (l *DelStockRollBackLogic) UpdateCacheRollBack(itemId string, stock int64) error {
	lockKey := util.CacheKey(types.CacheStockLock, itemId)
	key := util.CacheKey(types.CacheItemKey, itemId)

	_, err := l.svcCtx.BizRedis.Eval(decut_rollback_lua, []string{
		lockKey,
		key,
		types.CacheItemStock,
	}, itemId, strconv.FormatInt(stock, 10), types.CacheItemTime)
	if err != nil && err.Error() != "redis: nil" {
		logx.Errorf("BizRedis.Eval, error: %v", err)
		return err
	}

	return nil
}
