package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/util"

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
	//3、同步缓存
	group := threading.NewWorkerGroup(func() {
		_ = util.UpdateCache(l.ctx, l.svcCtx, req.Id)
	}, 1)
	group.Start()
	//4、返回响应
	return nil
}
