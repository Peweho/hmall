package svc

import (
	"hmall/application/order/mq/internal/config"
	"hmall/application/order/mq/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config     config.Config
	Db         *orm.DB
	OrderModel *model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	return &ServiceContext{
		Config:     c,
		Db:         db,
		OrderModel: model.NewOrderModel(db.DB),
	}
}
