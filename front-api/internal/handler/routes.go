// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "SimplePick-Mall-Server/front-api/internal/handler/auth"
	index "SimplePick-Mall-Server/front-api/internal/handler/index"
	"SimplePick-Mall-Server/front-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.MemberAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: auth.InfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/homeAdvertiseList",
				Handler: index.HomeAdvertiseListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/categoryList",
				Handler: index.CategoryListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/index"),
	)
}
