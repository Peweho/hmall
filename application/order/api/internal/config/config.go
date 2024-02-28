package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	BizRedis redis.RedisConf
	DB       struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	AddressRPC zrpc.RpcClientConf
	ItemRPC    zrpc.RpcClientConf
	OrderRPC   zrpc.RpcClientConf
	CartRPC    zrpc.RpcClientConf
}
