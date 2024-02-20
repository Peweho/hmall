package svc

import (
	"hmall/application/cart/rpc/internal/config"
	"hmall/application/cart/rpc/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config    config.Config
	Db        *orm.DB
	CartModel *model.CartModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:    c,
		Db:        db,
		CartModel: model.NewCartModel(db.DB),
	}
}
