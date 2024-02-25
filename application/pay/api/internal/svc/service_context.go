package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"hmall/application/order/rpc/order"
	"hmall/application/pay/api/internal/config"
	"hmall/application/pay/api/internal/model"
	"hmall/application/pay/rpc/pay"
	"hmall/application/user/rpc/user"
	"hmall/pkg/interceptors"
	"hmall/pkg/orm"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	Db       *orm.DB
	PayModel *model.PayModel
	OrderRPC order.Order
	UserRPC  user.User
	PayRPC   pay.Pay
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
	orderRPC := zrpc.MustNewClient(c.OrderRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	payRPC := zrpc.MustNewClient(c.PayRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:   c,
		BizRedis: rds,
		Db:       db,
		PayModel: model.NewPayModel(db.DB),
		OrderRPC: order.NewOrder(orderRPC),
		UserRPC:  user.NewUser(userRPC),
		PayRPC:   pay.NewPay(payRPC),
	}
}
