package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hmall/application/cart/api/internal/config"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/rpc/carts"
	"hmall/application/item/rpc/item"
	"hmall/pkg/interceptors"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config         config.Config
	ItemRPC        item.Item
	CartsPRC       carts.Carts
	BizRedis       *redis.Redis
	Db             *orm.DB
	CartModel      *model.CartModel
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	itemRPC := zrpc.MustNewClient(c.ItemRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	cartsRPC := zrpc.MustNewClient(c.CartsRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:         c,
		ItemRPC:        item.NewItem(itemRPC),
		CartsPRC:       carts.NewCarts(cartsRPC),
		BizRedis:       rds,
		Db:             db,
		CartModel:      model.NewCartModel(db.DB),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
