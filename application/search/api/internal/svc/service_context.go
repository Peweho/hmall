package svc

import (
	"github.com/elastic/go-elasticsearch/v8"
	"hmall/application/search/api/internal/config"
	"hmall/application/search/api/internal/model"
	"hmall/pkg/orm"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	Es        *elasticsearch.Client
	Db        *orm.DB
	ItemModel *model.ItemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	cfg := elasticsearch.Config{
		Addresses:              c.Es.Addresses,
		Username:               c.Es.Username,
		Password:               c.Es.Password,
		CertificateFingerprint: c.Es.Fingerprint,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("Something wrong with connection to Elasticsearch")
	}

	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	return &ServiceContext{
		Config:    c,
		Es:        es,
		Db:        db,
		ItemModel: model.NewItemModel(db.DB),
	}
}
