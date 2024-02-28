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
	pkgUtil "hmall/pkg/util"
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

func (l *DelStockLogic) DelStock(in *pb.DelStockReq) (*pb.DelStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		for _, val := range in.Detail {
			item, err := l.svcCtx.ItemModel.DecutStock(l.ctx, val.ItemId, val.Num)
			if err != nil {
				logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
				return err
			}
			if item.Stock < 0 {
				return dtmcli.ErrFailure
			}

			wg := &sync.WaitGroup{}
			wg.Add(2)
			//写缓存
			threading.GoSafe(func() {
				defer wg.Done()
				key := pkgUtil.CacheKey(types.CacheItemStockKey, strconv.Itoa(int(item.Id)))
				err = l.svcCtx.BizRedis.Set(key, strconv.Itoa(int(item.Stock)))
				if err != nil {
					logx.Errorf("BizRedis.Set: %v, error: %v", key, err)
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
			wg.Wait()
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	return &pb.DelStockResp{}, nil
}
