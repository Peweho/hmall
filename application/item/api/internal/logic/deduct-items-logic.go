package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	"hmall/application/item/api/internal/util"
	pkgUtil "hmall/pkg/util"
	"strconv"
	"sync"
)

type DeductItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeductItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductItemsLogic {
	return &DeductItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Order struct {
	ItemId int
	Num    int
}

func (l *DeductItemsLogic) DeductItems(req *types.DeductItemsReq) error {
	// 1、校验参数
	if len(req.Order) == 0 {
		return nil
	}

	// 2、用于同步es
	pusherSearch := util.NewPusherSearchLogic(l.ctx, l.svcCtx)

	_, err := mr.MapReduce[Order, *model.ItemDTO, int](func(source chan<- Order) {
		//解析参数
		for _, val := range req.Order {
			source <- Order{
				ItemId: val.ItemId,
				Num:    val.Num,
			}
		}
	}, func(order Order, writer mr.Writer[*model.ItemDTO], cancel func(error)) {
		//2、扣减库存 （下单时会查询库存是否足够）
		newItem, err := l.svcCtx.ItemModel.DecutStock(l.ctx, order.ItemId, order.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: id=%v,num=%v, error: %v", order.ItemId, order.Num, err)
			cancel(err)
		}
		writer.Write(newItem)
	}, func(pipe <-chan *model.ItemDTO, writer mr.Writer[int], cancel func(error)) {
		//3、同步缓存
		for newItem := range pipe {
			wg := &sync.WaitGroup{}
			wg.Add(2)
			//写缓存
			threading.GoSafe(func() {
				defer wg.Done()
				marshal, err := json.Marshal(newItem)
				if err != nil {
					logx.Errorf("json.Marshal: %v, error: %v", newItem, err)
					cancel(err)
				}
				key := pkgUtil.CacheKey(types.CacheItemStockKey, strconv.Itoa(int(newItem.Id)))
				err = l.svcCtx.BizRedis.Set(key, string(marshal))
				if err != nil {
					logx.Errorf("BizRedis.Set: %v, error: %v", key, err)
					cancel(err)
				}
			})

			//同步es
			threading.GoSafe(func() {
				defer wg.Done()
				err := pusherSearch.PusherSearch(types.KqUpdate, newItem)
				if err != nil {
					logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", *newItem, err)
					cancel(err)
				}
			})
			wg.Wait()
		}
		//没有结果，所以随便输出
		writer.Write(-1)
	})
	if err != nil {
		return err
	}

	//4、返回
	return nil
}
