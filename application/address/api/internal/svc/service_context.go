package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hmall/application/address/api/internal/config"
	"hmall/application/address/api/internal/model"
	"hmall/application/address/rpc/address"
	"hmall/pkg/interceptors"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config       config.Config
	BizRedis     *redis.Redis
	Db           *orm.DB
	AddressModel *model.AddressModel
	AddressRPC   address.Address
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	addressRPC := zrpc.MustNewClient(c.AddressRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:       c,
		BizRedis:     rds,
		Db:           db,
		AddressModel: model.NewAddressModel(db.DB),
		AddressRPC:   address.NewAddress(addressRPC),
	}
}
