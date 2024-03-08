package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"hmall/application/item/rpc/internal/config"
	"hmall/application/item/rpc/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config         config.Config
	ItemModel      *model.ItemModel
	Db             *orm.DB
	BizRedis       *redis.Redis
	KqPusherSearch *kq.Pusher
	KqPusherCache  *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:         c,
		Db:             db,
		ItemModel:      model.NewItemModel(db.DB),
		BizRedis:       rds,
		KqPusherSearch: kq.NewPusher(c.KqPusherSearch.Brokers, c.KqPusherSearch.Topic),
		KqPusherCache:  kq.NewPusher(c.KqPusherCache.Brokers, c.KqPusherCache.Topic),
	}
}
