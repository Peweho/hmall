package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hmall/application/item/mq/internal/config"
	"hmall/application/item/rpc/item"
	"hmall/pkg/interceptors"
)

type ServiceContext struct {
	Config   config.Config
	ItemRPC  item.Item
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	itemRPC := zrpc.MustNewClient(c.ItemRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		ItemRPC:  item.NewItem(itemRPC),
		BizRedis: rds,
	}
}
