package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hmall/application/address/rpc/address"
	"hmall/application/cart/rpc/carts"
	"hmall/application/item/rpc/item"
	"hmall/application/order/api/internal/config"
	"hmall/application/order/api/internal/model"
	"hmall/application/order/rpc/order"
	"hmall/pkg/interceptors"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config     config.Config
	BizRedis   *redis.Redis
	Db         *orm.DB
	OrderModel *model.OrderModel
	AddressRPC address.Address
	ItemRPC    item.Item
	OrderRPC   order.Order
	CartRPC    carts.Carts
}

func NewServiceContext(c config.Config) *ServiceContext {
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
	addressRPC := zrpc.MustNewClient(c.AddressRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	itemRPC := zrpc.MustNewClient(c.ItemRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	orderRPC := zrpc.MustNewClient(c.OrderRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	cartRPC := zrpc.MustNewClient(c.CartRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:     c,
		BizRedis:   rds,
		Db:         db,
		OrderModel: model.NewOrderModel(db.DB),
		AddressRPC: address.NewAddress(addressRPC),
		ItemRPC:    item.NewItem(itemRPC),
		OrderRPC:   order.NewOrder(orderRPC),
		CartRPC:    carts.NewCarts(cartRPC),
	}
}
