package logic

import (
	"context"
	"hmall/application/address/rpc/internal/svc"
	"hmall/application/address/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDefaultAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDefaultAddressLogic {
	return &GetUserDefaultAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDefaultAddressLogic) GetUserDefaultAddress(in *service.GetUserDefaultAddressReq) (*service.FindAdressByIdResp, error) {
	address, err := l.svcCtx.AddressModel.GetUserDefaultAddress(l.ctx, in.Uid)
	if err != nil {
		logx.Errorf("AddressModel.GetUserDefaultAddress: %v, error； %v", in.Uid, err)
		return nil, err
	}

	return &service.FindAdressByIdResp{
		Id:        int64(address.Id),
		Contact:   address.Contact,
		IsDefault: int64(address.IsDefault),
		Mobile:    address.Mobile,
		Town:      address.Town,
		Notes:     address.Notes,
		Province:  address.Province,
		Street:    address.Street,
		City:      address.City,
	}, nil
}
