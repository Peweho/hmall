// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"hmall/application/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/:orderId",
				Handler: MarkOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: FindOrderByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: CreateOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/flash",
				Handler: CreateFlashOrderHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/orders"),
	)
}
