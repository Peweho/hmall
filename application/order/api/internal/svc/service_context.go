package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"hmall/application/order/api/internal/config"
	"hmall/application/order/api/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config         config.Config
	BizRedis       *redis.Redis
	Db             *orm.DB
	AddressModel   *model.OrderModel
	KqPusherClient *kq.Pusher
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
	return &ServiceContext{
		Config:         c,
		BizRedis:       rds,
		Db:             db,
		AddressModel:   model.NewOrderModel(db.DB),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
