package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	"hmall/application/item/api/internal/util"
	pkgUtil "hmall/pkg/util"
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
	pusherSearch := util.NewPusherSearchLogic(l.ctx, l.svcCtx)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//写缓存
	threading.GoSafe(func() {
		defer wg.Done()
		marshal, err := json.Marshal(newItem)
		if err != nil {
			logx.Errorf("json.Marshal: %v, error: %v", *newItem, err)
			panic(err)
		}

		key := pkgUtil.CacheKey(types.CacheItemStockKey, strconv.Itoa(int(newItem.Id)))
		err = l.svcCtx.BizRedis.Set(key, string(marshal))
		if err != nil {
			logx.Errorf("BizRedis.Set: %v, error: %v", key, err)
			panic(err)
		}
	})

	//同步es
	threading.GoSafe(func() {
		defer wg.Done()
		err = pusherSearch.PusherSearch(types.KqUpdate, newItem)
		if err != nil {
			logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", *newItem, err)
			panic(err)
		}
	})
	wg.Wait()
	//3、返回响应
	return nil

}
