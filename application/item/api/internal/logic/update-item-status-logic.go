package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	utils "hmall/application/item/api/internal/util"
	"hmall/pkg/util"
	"strconv"
	"sync"
)

type UpdateItemStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemStatusLogic {
	return &UpdateItemStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemStatusLogic) UpdateItemStatus(req *types.UpdateItemStatusReq) error {

	//1、调用数据库
	newItem, err := l.svcCtx.ItemModel.UpdateItemStatusById(l.ctx, req.Id, req.Status)
	if err != nil {
		logx.Errorf("ItemModel.UpdateItemStatusById: %v, error: %v", req.Id, err)
		return err
	}
	//3、同步缓存 es
	pusherSearch := utils.NewPusherSearchLogic(l.ctx, l.svcCtx)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//写缓存
	threading.GoSafe(func() {
		defer wg.Done()

		key := util.CacheKey(types.CacheItemKey, strconv.Itoa(req.Id))
		err = l.svcCtx.BizRedis.Hset(key, types.CacheItemStatus, strconv.Itoa(req.Status))
		if err != nil {
			logx.Errorf("BizRedis.Set: %v, error: %v", key, err)

			//缓存失败，进行补偿
			cacheLogic := utils.NewPusherLogic(l.ctx, l.svcCtx)
			//构造对象
			msg := &utils.KqCacheMsg{
				Code:  types.KqCacheStatus,
				Stock: strconv.Itoa(req.Status),
				Key:   key,
			}

			if errKq := cacheLogic.Pusher(msg); errKq != nil {
				logx.Errorf("acheLogic.Pusher: %v, error: %v", msg, err)
				panic(errKq)
			}
		}
		_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime)
	})

	//同步es
	threading.GoSafe(func() {
		defer wg.Done()
		var err error
		switch req.Status {
		case types.ItemStatusNormal:
			err = pusherSearch.PusherSearch(types.KqUpdate, newItem)
		case types.ItemStatusRemove:
			err = pusherSearch.PusherSearch(types.KqDel, newItem)
		case types.ItemStatusDeleted:
			err = pusherSearch.PusherSearch(types.KqDel, newItem)
		default:
			logx.Errorf("switch req.Status: %v, error: %v", req.Status, "商品状态不正确")
			panic("商品状态不正确")
		}
		if err != nil {
			logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", *newItem, err)
			panic(err)
		}
	})
	wg.Wait()
	//3、返回响应
	return nil

}
