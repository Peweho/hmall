package logic

import (
	"context"
	"fmt"
	"hmall/application/item/rpc/item"
	"hmall/pkg/xcode"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindItemsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindItemsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemsByIdLogic {
	return &FindItemsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindItemsByIdLogic) FindItemsById(req *types.FindItemsByIdReq) (resp *types.FindItemsByIdResp, err error) {
	//1、参数校验
	if len(req.Ids) == 0 {
		return nil, xcode.New(types.NotFound, "")
	}

	items := make([]types.Item, 0, len(req.Ids)) //响应结果

	//3、调用rpc方法
	res, err := l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: req.Ids})
	if err != nil {
		logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", req.Ids, err)
		return nil, xcode.New(types.NotFound, types.RpcError)
	}
	//4、类型转换
	for _, v := range res.Data {
		items = append(items, ItemPRC_To_Item(v))
	}
	//5、返回响应
	return &types.FindItemsByIdResp{
		Data: items,
	}, nil
}

func ItemPRC_To_Item(item *item.Items) types.Item {
	return types.Item{
		Id:           int(item.Id),
		Brand:        item.Brand,
		Category:     item.Category,
		CommentCount: int(item.CommentCount),
		Image:        item.Image,
		IsAD:         item.IsAD,
		Name:         item.Name,
		Price:        int(item.Price),
		Sold:         int(item.Sold),
		Spec:         item.Spec,
		Status:       int(item.Status),
		Stock:        int(item.Stock),
	}
}

// 构造对应的缓存键
func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheItemKey, id)
}
