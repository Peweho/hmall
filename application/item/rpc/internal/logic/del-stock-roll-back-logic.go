package logic

import (
	"bufio"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"io"
	"os"
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

func (l *DelStockRollBackLogic) DelStockRollBack(in *pb.DelStockReq) (*pb.DelStockResp, error) {
	for _, val := range in.Detail {
		item, err := l.svcCtx.ItemModel.AddStock(l.ctx, val.ItemId, val.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
			return nil, err
		}

		wg := &sync.WaitGroup{}
		wg.Add(2)
		//写缓存
		threading.GoSafe(func() {
			defer wg.Done()
			key := util.CacheKey(types.CacheItemKey, strconv.FormatInt(item.Id, 10))
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
	return &pb.DelStockResp{}, nil
}

func (l *DelStockRollBackLogic) UpdateCacheRollBack(itemId string, stock int64) error {
	lockKey := util.CacheKey(types.CacheStockLock, itemId)
	key := util.CacheKey(types.CacheItemKey, itemId)
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	//缓存存在，更新缓存
	if exists {
		//读取lua脚本
		luaScript, err := ReadLuaRollBack()
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

func ReadLuaRollBack() (string, error) {
	//打开脚本
	file, err := os.Open(types.LuapathRollBack)
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
