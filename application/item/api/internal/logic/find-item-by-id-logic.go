package logic

import (
	"context"
	"hmall/application/item/rpc/item"
	"hmall/pkg/xcode"
	"strconv"

	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindItemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindItemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemByIdLogic {
	return &FindItemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindItemByIdLogic) FindItemById(req *types.FindItemByIdReq) (resp *types.ItemReqAndResp, err error) {

	//1、调用rpc方法
	res, err := l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: []string{strconv.Itoa(req.Id)}})
	if err != nil {
		logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", req.Id, err)
		return nil, xcode.New(types.NotFound, types.RpcError)
	}

	//2、返回响应
	if len(res.Data) == 0 {
		return nil, xcode.New(types.NotFound, "")
	}
	data := res.Data[0]
	return &types.ItemReqAndResp{
		Id:           int(data.Id),
		Brand:        data.Brand,
		Category:     data.Category,
		CommentCount: int(data.CommentCount),
		Image:        data.Image,
		IsAD:         data.IsAD,
		Name:         data.Name,
		Price:        int(data.Price),
		Sold:         int(data.Sold),
		Spec:         data.Spec,
		Status:       int(data.Status),
		Stock:        int(data.Stock),
	}, nil
}
