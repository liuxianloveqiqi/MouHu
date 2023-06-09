// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"MouHu/service/app/like/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/follow/add",
					Handler: FollowUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/like/add",
					Handler: LikeAddHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/like"),
	)
}
