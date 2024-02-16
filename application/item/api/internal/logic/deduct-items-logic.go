package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	"hmall/application/item/api/internal/util"

	"github.com/zeromicro/go-zero/core/logx"
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

	_, err := mr.MapReduce[Order, int, int](func(source chan<- Order) {
		//解析参数
		for _, val := range req.Order {
			source <- Order{
				ItemId: val.ItemId,
				Num:    val.Num,
			}
		}
	}, func(order Order, writer mr.Writer[int], cancel func(error)) {
		//2、扣减库存 （下单时会查询库存是否足够）
		err := l.svcCtx.ItemModel.DecutStock(l.ctx, order.ItemId, order.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: id=%v,num=%v, error: %v", order.ItemId, order.Num, err)
			cancel(err)
		}
		writer.Write(order.ItemId)
	}, func(pipe <-chan int, writer mr.Writer[int], cancel func(error)) {
		//3、同步缓存
		for id := range pipe {
			err := util.UpdateCache(l.ctx, l.svcCtx, id)
			if err != nil {
				cancel(err)
			}
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
