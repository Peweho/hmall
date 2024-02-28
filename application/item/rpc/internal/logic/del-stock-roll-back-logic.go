package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	pkgUtil "hmall/pkg/util"
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
	return &pb.DelStockResp{}, nil
}
