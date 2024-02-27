package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hmall/application/item/api/internal/config"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/rpc/item"
	"hmall/pkg/interceptors"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config                    config.Config
	ItemRPC                   item.Item
	BizRedis                  *redis.Redis
	Db                        *orm.DB
	ItemModel                 *model.ItemModel
	KqPusherClientUpdateCache *kq.Pusher
	KqPusherSearch            *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	itemRPC := zrpc.MustNewClient(c.ItemRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                    c,
		ItemRPC:                   item.NewItem(itemRPC),
		BizRedis:                  rds,
		Db:                        db,
		ItemModel:                 model.NewItemModel(db.DB),
		KqPusherClientUpdateCache: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		KqPusherSearch:            kq.NewPusher(c.KqPusherSearch.Brokers, c.KqPusherSearch.Topic),
	}
}
