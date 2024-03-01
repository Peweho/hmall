package logic

import (
	"bufio"
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
	"io"
	"os"
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
				return status.Error(codes.Internal, err.Error())
			}

			//库存不足回滚
			if item.Stock < 0 {
				return status.Error(codes.Aborted, dtmcli.ResultFailure)
			}

			wg := &sync.WaitGroup{}
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
			wg.Wait()
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.DelStockResp{}, nil
}

func (l *DelStockLogic) UpdateCache(itemId string, stock int64) error {
	lockKey := util.CacheKey(types.CacheStockLock, itemId)
	key := util.CacheKey(types.CacheItemKey, itemId)
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	//缓存存在，更新缓存
	if exists {
		//读取lua脚本
		luaScript, err := ReadLua()
		if err != nil {
			return err
		}
		// 执行Lua脚本
		_, err = l.svcCtx.BizRedis.Eval(luaScript, []string{
			lockKey,
			key,
			types.CacheItemStock,
		}, itemId, strconv.FormatInt(stock, 10), types.CacheItemTime)
		if err != nil && err.Error() != "redis: nil" {
			logx.Errorf("BizRedis.Eval, error: %v", err)
			return err
		}
	}
	return nil
}

func ReadLua() (string, error) {
	//打开脚本
	file, err := os.Open(types.Luapath)
	if err != nil {
		logx.Errorf("os.Open: %v,error: %v", types.Luapath, err)
		return "", err
	}
	reader := bufio.NewReader(file)
	var luaScript string
	//逐行读取
	for {
		line, err := reader.ReadString('\n')
		luaScript += line
		if err == io.EOF {
			break
		}
	}

	return luaScript, nil
}
