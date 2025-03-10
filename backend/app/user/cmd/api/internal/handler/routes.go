// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "Table/app/user/cmd/api/internal/handler/admin"
	user "Table/app/user/cmd/api/internal/handler/user"
	"Table/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Coors},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: loginHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Coors},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/logout",
					Handler: logoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Coors},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/userinfo",
					Handler: user.UserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/wage",
					Handler: user.CalwageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/querywage",
					Handler: user.GetwageHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Coors},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/lookup",
					Handler: admin.LookupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: admin.UploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/wage",
					Handler: admin.AllwageHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin"),
	)
}
