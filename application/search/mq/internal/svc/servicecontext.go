package svc

import (
	"github.com/elastic/go-elasticsearch/v8"
	"hmall/application/search/mq/internal/config"
	"log"
)

type ServiceContext struct {
	Config config.Config
	Es     *elasticsearch.Client
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

	return &ServiceContext{
		Config: c,
		Es:     es,
	}
}
