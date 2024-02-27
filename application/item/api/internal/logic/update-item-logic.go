package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/util"
	pkgUtil "hmall/pkg/util"
	"strconv"
	"sync"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemLogic) UpdateItem(req *types.ItemReqAndResp) error {
	//1、构建参数
	item := util.ItemReqAndResp_To_ItemDTO(req)
	//2、调用数据库
	err := l.svcCtx.ItemModel.UpdateItemById(l.ctx, item)
	if err != nil {
		logx.Errorf("ItemModel.UpdateItemById: %v, error: %v", item, err)
		return err
	}
	//3、同步缓存 es
	pusherSearch := util.NewPusherSearchLogic(l.ctx, l.svcCtx)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//写缓存
	threading.GoSafe(func() {
		defer wg.Done()
		marshal, err := json.Marshal(item)
		if err != nil {
			logx.Errorf("json.Marshal: %v, error: %v", item, err)
			panic(err)
		}

		key := pkgUtil.CacheKey(types.CacheItemStockKey, strconv.Itoa(int(item.Id)))
		err = l.svcCtx.BizRedis.Set(key, string(marshal))
		if err != nil {
			logx.Errorf("BizRedis.Set: %v, error: %v", key, err)
			panic(err)
		}
	})

	//同步es
	threading.GoSafe(func() {
		defer wg.Done()
		err = pusherSearch.PusherSearch(types.KqUpdate, &item)
		if err != nil {
			logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", item, err)
			panic(err)
		}
	})
	wg.Wait()
	//4、返回响应
	return nil
}
