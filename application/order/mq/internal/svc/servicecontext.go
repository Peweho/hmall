package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"hmall/application/order/mq/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
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

	return &ServiceContext{
		Config:   c,
		BizRedis: rds,
	}
}
