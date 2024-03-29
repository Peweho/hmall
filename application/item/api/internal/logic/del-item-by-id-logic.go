package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	utils "hmall/application/item/api/internal/util"
	"hmall/pkg/util"
	"strconv"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelItemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Bloom  *bloom.Filter
}

func NewDelItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelItemByIdLogic {
	return &DelItemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		Bloom:  bloom.New(svcCtx.BizRedis, types.ItemBloomKey, 20*svcCtx.Config.ItemNums),
	}
}

func (l *DelItemByIdLogic) DelItemById(req *types.DelItemByIdReq) error {
	err := l.svcCtx.ItemModel.DelItemById(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("ItemModel.DelItemById: %v,error: %v", req.Id, err)
		return err
	}

	//3、同步缓存 es
	pusherSearch := utils.NewPusherSearchLogic(l.ctx, l.svcCtx)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//删除缓存
	threading.GoSafe(func() {
		defer wg.Done()
		key := util.CacheKey(types.CacheItemKey, strconv.Itoa(req.Id))
		cacheLogic := utils.NewPusherLogic(l.ctx, l.svcCtx)
		//构造对象
		msg := &utils.KqCacheMsg{
			Code: types.KqCacheDel,
			Key:  key,
		}
		if errKq := cacheLogic.Pusher(msg); errKq != nil {
			logx.Errorf("acheLogic.Pusher: %v, error: %v", msg, err)
			panic(errKq)
		}
	})

	//同步es
	threading.GoSafe(func() {
		defer wg.Done()
		err = pusherSearch.PusherSearch(types.KqDel, &model.ItemDTO{Id: int64(req.Id)})
		if err != nil {
			logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", req.Id, err)
			panic(err)
		}
	})

	wg.Wait()
	return nil
}
