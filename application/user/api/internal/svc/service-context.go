package svc

import (
	"hmall/application/user/api/internal/config"
	"hmall/application/user/api/internal/model"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
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
		UserModel: model.NewEmployeeGormModel(db.DB),
	}
}
