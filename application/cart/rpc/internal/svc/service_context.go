package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"hmall/application/cart/rpc/internal/config"
	"hmall/application/cart/rpc/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config    config.Config
	Db        *orm.DB
	CartModel *model.CartModel
	BizRedis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		Db:        db,
		CartModel: model.NewCartModel(db.DB),
		BizRedis:  rds,
	}
}
