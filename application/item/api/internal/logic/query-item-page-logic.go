package logic

import (
	"context"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	"hmall/application/item/api/internal/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryItemPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryItemPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryItemPageLogic {
	return &QueryItemPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryItemPageLogic) QueryItemPage(req *types.QueryItemPageReq) (resp *types.QueryItemPageResp, err error) {
	// 0、参数处理
	var (
		sortBy   string
		isAsc    string
		page     int
		pageSize int
	)
	if req.SortBy == "" {
		sortBy = types.SortBy
	}
	if req.IsAsc == "" {
		isAsc = types.IsAsc
	}
	if req.PageNo == 0 {
		page = types.Page
	}
	if req.PageSize == 0 {
		pageSize = types.PageSize
	}
	// 1、查询数据库
	res, err := l.svcCtx.ItemModel.QueryItemPage(l.ctx, page, pageSize, sortBy, isAsc)
	if err != nil {
		logx.Errorf("ItemModel.QueryItemPage, error: %v", err)
		return nil, err
	}
	//2、返回响应
	ans := make([]types.Item, 0, len(res))
	for _, val := range res {
		ans = append(ans, util.ItemDTO_To_Item(val))
	}
	return &types.QueryItemPageResp{
		Pages: 1,
		Total: len(ans),
		List:  ans,
	}, nil
}
