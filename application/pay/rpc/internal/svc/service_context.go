package svc

import (
	"hmall/application/pay/rpc/internal/config"
	"hmall/application/pay/rpc/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config   config.Config
	Db       *orm.DB
	PayModel *model.PayModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	return &ServiceContext{
		Db:       db,
		PayModel: model.NewPayModel(db.DB),
		Config:   c,
	}
}
