package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/address/rpc/internal/svc"
	"hmall/application/address/rpc/internal/utils"
	"hmall/application/address/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAdressByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAdressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAdressByIdLogic {
	return &FindAdressByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAdressByIdLogic) FindAdressById(in *service.FindAdressByIdReq) (*service.FindAdressByIdResp, error) {
	//1、查缓存
	key := utils.CacheKey(in.Id)
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	if exists {
		//缓存续期
		threading.NewWorkerGroup(func() {
			_ = l.svcCtx.BizRedis.Expire(key, utils.CacheAddressTime)
		}, 1).Start()

		var res service.FindAdressByIdResp
		get, err := l.svcCtx.BizRedis.Get(key)
		if err != nil {
			logx.Error("BizRedis.Get: %v, error: %v", key, err)
			return nil, err
		}

		if err = json.Unmarshal([]byte(get), &res); err != nil {
			logx.Error("json.Unmarshal: %v, error: %v", get, err)
			return nil, err
		}
		return &res, nil
	}

	//2、查数据库
	address, err := l.svcCtx.AddressModel.QueryAddressFindById(l.ctx, in.Id)
	if err != nil {
		logx.Error("AddressModel.QueryAddressFindById: %v, error: %v", in.Id, err)
		return nil, err
	}
	//3、异步写缓存
	threading.NewWorkerGroup(func() {
		marshal, err := json.Marshal(address)
		if err != nil {
			logx.Errorf("json.Marshal: %v, error: %v", address, err)
			return
		}
		if err = l.svcCtx.BizRedis.Set(key, string(marshal)); err != nil {
			logx.Errorf("BizRedis.Set: %v, error: %v", key, err)
		}
		_ = l.svcCtx.BizRedis.Expire(key, utils.CacheAddressTime)
	}, 1).Start()

	res := &service.FindAdressByIdResp{
		Id:        int64(address.Id),
		Contact:   address.Contact,
		IsDefault: int64(address.IsDefault),
		Mobile:    address.Mobile,
		Town:      address.Town,
		Notes:     address.Notes,
		Province:  address.Province,
		Street:    address.Street,
		City:      address.City,
	}

	return res, nil
}
