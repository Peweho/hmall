// Code generated by goctl. DO NOT EDIT.
// Source: item.proto

package server

import (
	"context"

	"hmall/application/item/rpc/internal/logic"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/types/service"
)

type ItemServer struct {
	svcCtx *svc.ServiceContext
	service.UnimplementedItemServer
}

func NewItemServer(svcCtx *svc.ServiceContext) *ItemServer {
	return &ItemServer{
		svcCtx: svcCtx,
	}
}

func (s *ItemServer) FindItemByIds(ctx context.Context, in *service.FindItemByIdsReq) (*service.FindItemByIdsResp, error) {
	l := logic.NewFindItemByIdsLogic(ctx, s.svcCtx)
	return l.FindItemByIds(in)
}
