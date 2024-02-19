package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/address/api/internal/utils"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"sync"
	"time"

	"hmall/application/address/api/internal/svc"
	"hmall/application/address/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAddressesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAddressesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAddressesLogic {
	return &QueryAddressesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAddressesLogic) QueryAddresses() (resp *types.QueryAddressesResp, err error) {
	// 1、查询缓存
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return nil, xcode.New(types.Unauthorized, "")
	}
	key := utils.CacheIds(usr)
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	if exists {
		addresses, err := l.CacheFindAddressById(key)
		if err != nil {
			logx.Errorf("CacheFindAddressById, error: %v", err)
			return nil, err
		}
		return &types.QueryAddressesResp{Addresses: addresses}, xcode.New(types.OK, "")
	}
	//2、查询数据库
	addresses, err := l.svcCtx.AddressModel.QueryAddresses(l.ctx, usr)
	if err != nil {
		logx.Errorf("AddressModel.QueryAddresses: %v, error: %v", usr, err)
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, xcode.New(types.NotFound, "")
	}
	addressesDTO := make([]types.QueryAddressesDTO, 0, len(addresses))

	wg := &sync.WaitGroup{}
	for _, val := range addresses {
		temp := utils.AddressPO_to_AddresDTO(&val)
		addressesDTO = append(addressesDTO, *temp)
		//3、异步写缓存
		threading.GoSafe(func() {
			wg.Add(1)
			marshal, err := json.Marshal(temp)
			if err != nil {
				logx.Errorf("json.Marshal: %v, error: %v", temp, err)
				wg.Done()
				return
			}
			_, err = l.svcCtx.BizRedis.Zadd(key, time.Now().Unix(), string(marshal))
			if err != nil {
				logx.Errorf("BizRedis.Zadd: %v, error: %v", string(marshal), err)
				wg.Done()
				return
			}
			wg.Done()
		})
	}
	wg.Wait()
	return &types.QueryAddressesResp{Addresses: addressesDTO}, xcode.New(types.OK, "")
}

// 查询缓存
func (l *QueryAddressesLogic) CacheFindAddressById(key string) ([]types.QueryAddressesDTO, error) {
	zaddress, err := l.svcCtx.BizRedis.ZrangeWithScores(key, 0, time.Now().Unix())
	if err != nil {
		logx.Errorf("BizRedis.ZrangeWithScores: %v, error: %v", key, err)
		return nil, err
	}
	addressesDTO := make([]types.QueryAddressesDTO, len(zaddress))
	for i, v := range zaddress {
		if err := json.Unmarshal([]byte(v.Key), &addressesDTO[i]); err != nil {
			logx.Errorf("json.Unmarshal: %v, error: %v", v.Key, err)
			return nil, err
		}
	}
	return addressesDTO, nil
}
