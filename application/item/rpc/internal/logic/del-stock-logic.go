package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"strconv"
	"sync"
)

type DelStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStockLogic {
	return &DelStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var decut_lua = `
	local acquired = redis.call("exist",KEYS[2])
	if acquired == 1 then
		local stock = redis.call("HGET",KEYS[2],KEYS[3])
		redis.call("HSET",KEYS[2],KEYS[3],stock - ARGV[2])
		redis.call("EXPIRE",KEYS[2],ARGV[3])
	end
`

func (l *DelStockLogic) DelStock(in *pb.DelStockReq) (*pb.DelStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		wg := &sync.WaitGroup{}
		var res error = nil
		for _, val := range in.Detail {
			item, err := l.svcCtx.ItemModel.DecutStock(l.ctx, val.ItemId, val.Num)
			if err != nil {
				logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
				return status.Error(codes.Internal, err.Error())
			}

			//库存不足回滚
			if res == nil && item.Stock < 0 {
				res = status.Error(codes.Aborted, dtmcli.ResultFailure)
			}

			wg.Add(2)
			//更新缓存
			threading.GoSafe(func() {
				defer wg.Done()
				err := l.UpdateCache(val.ItemId, val.Num)
				if err != nil {
					panic(err)
				}
			})

			//同步es
			threading.GoSafe(func() {
				defer wg.Done()
				pusherSearch := NewPusherSearchLogic(l.ctx, l.svcCtx)
				err = pusherSearch.PusherSearch(types.KqUpdate, item)
				if err != nil {
					logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", *item, err)
					panic(err)
				}
			})
		}
		wg.Wait()
		return res
	}); err != nil {
		return nil, err
	}

	return &pb.DelStockResp{}, nil
}

func (l *DelStockLogic) UpdateCache(itemId string, stock int64) error {
	lockKey := util.CacheKey(types.CacheStockLock, itemId)
	key := util.CacheKey(types.CacheItemKey, itemId)

	// 执行Lua脚本
	_, err := l.svcCtx.BizRedis.Eval(decut_lua, []string{
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
