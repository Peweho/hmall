// Code generated by goctl. DO NOT EDIT.
// Source: address.proto

package server

import (
	"context"

	"hmall/application/address/rpc/internal/logic"
	"hmall/application/address/rpc/internal/svc"
	"hmall/application/address/rpc/service"
)

type AddressServer struct {
	svcCtx *svc.ServiceContext
	service.UnimplementedAddressServer
}

func NewAddressServer(svcCtx *svc.ServiceContext) *AddressServer {
	return &AddressServer{
		svcCtx: svcCtx,
	}
}

func (s *AddressServer) FindAdressById(ctx context.Context, in *service.FindAdressByIdReq) (*service.FindAdressByIdResp, error) {
	l := logic.NewFindAdressByIdLogic(ctx, s.svcCtx)
	return l.FindAdressById(in)
}