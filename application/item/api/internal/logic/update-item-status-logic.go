package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	"hmall/application/item/api/internal/util"
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
	err := l.svcCtx.ItemModel.UpdateItemStatusById(l.ctx, req.Id, req.Status)
	if err != nil {
		logx.Errorf("ItemModel.UpdateItemStatusById: %v, error: %v", req.Id, err)
		return err
	}
	//2、同步缓存
	group := threading.NewWorkerGroup(func() {
		_ = util.UpdateCache(l.ctx, l.svcCtx, req.Id)
	}, 1)
	group.Start()
	//3、返回响应
	return nil

}
