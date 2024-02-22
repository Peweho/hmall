package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	BizRedis redis.RedisConf
	DB       struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
	AddressRPC zrpc.RpcClientConf
	ItemRPC    zrpc.RpcClientConf
}
