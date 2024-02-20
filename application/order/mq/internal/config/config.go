package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	KqConsumerConf kq.KqConf
	DB             struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
}
