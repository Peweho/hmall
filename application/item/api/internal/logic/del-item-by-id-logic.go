package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/util"
	"strconv"
	"sync"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelItemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelItemByIdLogic {
	return &DelItemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelItemByIdLogic) DelItemById(req *types.DelItemByIdReq) error {
	err := l.svcCtx.ItemModel.DelItemById(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("ItemModel.DelItemById: %v,error: %v", req.Id, err)
		return err
	}

	//3、同步缓存 es
	pusherSearch := util.NewPusherSearchLogic(l.ctx, l.svcCtx)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//删除缓存
	threading.GoSafe(func() {
		defer wg.Done()
		key := util.CacheIds(strconv.Itoa(req.Id))
		_, err = l.svcCtx.BizRedis.Del(key)
		if err != nil {
			logx.Errorf("BizRedis.Del: %v,error: %v", key, err)
			panic(err)
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
