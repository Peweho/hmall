package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	utils "hmall/application/item/api/internal/util"
	"hmall/pkg/util"
	"strconv"
	"sync"
)

type AdditemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Bloom  *bloom.Filter
}

func NewAdditemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdditemLogic {
	return &AdditemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		Bloom:  bloom.New(svcCtx.BizRedis, types.ItemBloomKey, 20*svcCtx.Config.ItemNums),
	}
}

func (l *AdditemLogic) Additem(req *types.ItemReqAndResp) error {
	item := &model.ItemDTO{
		Brand:        req.Brand,
		Category:     req.Category,
		CommentCount: int64(req.CommentCount),
		Image:        req.Image,
		IsAD:         req.IsAD,
		Name:         req.Name,
		Price:        int64(req.Price),
		Sold:         int64(req.Sold),
		Spec:         req.Spec,
		Status:       int64(req.Status),
		Stock:        int64(req.Stock),
	}

	//操作数据库
	err := l.svcCtx.ItemModel.InserItem(l.ctx, item)
	if err != nil {
		logx.Errorf("ItemModel.InserItem: %v, error: %v", *item, err)
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	//添加到布隆过滤器
	threading.GoSafe(func() {
		defer wg.Done()
		err := l.Bloom.AddCtx(l.ctx, []byte(strconv.FormatInt(item.Id, 10)))
		panic(err)
	})

	//同步es
	threading.GoSafe(func() {
		defer wg.Done()
		pusherSearch := utils.NewPusherSearchLogic(l.ctx, l.svcCtx)
		if err := pusherSearch.PusherSearch(types.KqUpdate, item); err != nil {
			logx.Errorf(" pusherSearch.PusherSearch: %v, error: %v", item, err)
		}
	})

	//同步缓存
	threading.GoSafe(func() {
		defer wg.Done()
		key := util.CacheKey(types.CacheItemKey, strconv.FormatInt(item.Id, 10))
		marshal, err := json.Marshal(item)
		if err != nil {
			logx.Errorf(" pusherSearch.PusherSearch: %v, error: %v", item, err)
		}

		if err := l.svcCtx.BizRedis.Hmset(key, map[string]string{
			types.CacheItemFields: string(marshal),
			types.CacheItemStatus: strconv.FormatInt(item.Status, 10),
			types.CacheItemStock:  strconv.FormatInt(item.Stock, 10),
		}); err != nil {
			logx.Errorf("BizRedis.Hmset: %v, error: %v", key, err)

			//缓存失败，进行补偿
			cacheLogic := utils.NewPusherLogic(l.ctx, l.svcCtx)
			//构造对象
			msg := &utils.KqCacheMsg{
				Code:   types.KqCacheAll,
				Field:  string(marshal),
				Status: strconv.FormatInt(item.Status, 10),
				Stock:  strconv.FormatInt(item.Status, 10),
				Key:    key,
			}

			if errKq := cacheLogic.Pusher(msg); errKq != nil {
				logx.Errorf("acheLogic.Pusher: %v, error: %v", msg, err)
				panic(errKq)
			}
		}
		_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime)
	})

	wg.Wait()
	return nil
}
