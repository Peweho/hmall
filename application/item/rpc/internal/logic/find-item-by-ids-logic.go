package logic

import (
	"context"
	"hmall/application/item/rpc/internal/model"
	"hmall/pkg/xcode"

	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/types/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindItemByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindItemByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemByIdsLogic {
	return &FindItemByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindItemByIdsLogic) FindItemByIds(in *service.FindItemByIdsReq) (*service.FindItemByIdsResp, error) {
	// 1、校验参数
	if len(in.Ids) == 0 {
		return nil, xcode.New(200, "ids为空")
	}

	//2、查询数据
	res, err := l.svcCtx.ItemModel.FindItemByIds(l.ctx, in.Ids)
	if err != nil {
		logx.Errorf("ItemModel.FindItemByIds: %v ,error: %v", in.Ids, err)
		return nil, err
	}
	//3、类型转换
	items := make([]*service.Items, 0, len(res))
	for _, v := range res {
		items = append(items, ItemDTO_To_Item(v))
	}
	//4、返回响应
	return &service.FindItemByIdsResp{Data: items}, nil
}

func ItemDTO_To_Item(item model.ItemDTO) *service.Items {
	return &service.Items{
		Id:           item.Id,
		Brand:        item.Brand,
		Category:     item.Category,
		CommentCount: item.CommentCount,
		Image:        item.Image,
		IsAD:         item.IsAD,
		Name:         item.Name,
		Price:        item.Price,
		Sold:         item.Sold,
		Spec:         item.Spec,
		Status:       item.Status,
		Stock:        item.Stock,
	}
}
