package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hmall/pkg/xcode"

	"hmall/application/item/rpc/internal/config"
	"hmall/application/item/rpc/internal/server"
	"hmall/application/item/rpc/internal/svc"
	svc1 "hmall/application/item/rpc/types/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/item.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		svc1.RegisterItemServer(grpcServer, server.NewItemServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//自定义错误码
	httpx.SetErrorHandler(xcode.ErrHandler)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
