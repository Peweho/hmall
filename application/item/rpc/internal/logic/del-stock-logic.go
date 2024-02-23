package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *DelStockLogic) DelStock(in *pb.DelStockReq) (*pb.DelStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		for _, val := range in.Detail {
			key := util.CacheKey(types.CacheItemStockKey, val.ItemId)
			exists, _ := l.svcCtx.BizRedis.Exists(key)
			//判断缓存是否存在
			if exists {
				get, cacheErr := l.svcCtx.BizRedis.Get(key)
				cacheNum, atoiErr := strconv.Atoi(get)
				if cacheErr == nil && atoiErr == nil {
					//进行扣减库存和缓存续期
					_ = l.svcCtx.BizRedis.Set(key, strconv.Itoa(cacheNum-int(val.Num)))
					_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime)

					//不足返回错误，进行回滚
					if cacheNum < int(val.Num) {
						return dtmcli.ErrFailure
					}
					//无需查询数据库
					continue
				}
			}
			item, err := l.svcCtx.ItemModel.DecutStock(l.ctx, val.ItemId, val.Num)
			if err != nil {
				logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
				return err
			}
			if item.Stock < 0 {
				return dtmcli.ErrFailure
			}
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	return &pb.DelStockResp{}, nil
}
