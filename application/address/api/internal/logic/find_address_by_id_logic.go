package logic

import (
	"context"
	"hmall/application/address/api/internal/svc"
	"hmall/application/address/api/internal/types"
	"hmall/application/address/rpc/address"

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
	rpcResp, err := l.svcCtx.AddressRPC.FindAdressById(l.ctx, &address.FindAdressByIdReq{
		Id: int64(req.AddressId),
	})
	if err != nil {
		logx.Error("AddressModel.QueryAddressFindById: %v, error: %v", req.AddressId, err)
		return nil, err
	}
	return &types.FindAddressByIdResp{
		types.QueryAddressesDTO{
			Id:        int(rpcResp.Id),
			Contact:   rpcResp.Contact,
			IsDefault: int(rpcResp.IsDefault),
			Mobile:    rpcResp.Mobile,
			Town:      rpcResp.Town,
			Notes:     rpcResp.Notes,
			Province:  rpcResp.Province,
			Street:    rpcResp.Street,
			City:      rpcResp.City,
		},
	}, nil
}
