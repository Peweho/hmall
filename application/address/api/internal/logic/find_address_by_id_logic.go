package logic

import (
	"context"
	"hmall/application/address/api/internal/utils"
	"hmall/pkg/xcode"

	"hmall/application/address/api/internal/svc"
	"hmall/application/address/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAddressByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAddressByIdLogic {
	return &FindAddressByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAddressByIdLogic) FindAddressById(req *types.FindAddressByIdReq) (resp *types.FindAddressByIdResp, err error) {
	address, err := l.svcCtx.AddressModel.QueryAddressFindById(l.ctx, req.AddressId)
	if err != nil {
		logx.Error("AddressModel.QueryAddressFindById: %v, error: %v", req.AddressId, err)
	}
	return &types.FindAddressByIdResp{
		*utils.AddressPO_to_AddresDTO(&address),
	}, xcode.New(types.OK, "")
}
