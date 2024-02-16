package svc

import (
	"hmall/application/item/rpc/internal/config"
	"hmall/application/item/rpc/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config    config.Config
	ItemModel *model.ItemModel
	Db        *orm.DB
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
		ItemModel: model.NewItemModel(db.DB),
	}
}
