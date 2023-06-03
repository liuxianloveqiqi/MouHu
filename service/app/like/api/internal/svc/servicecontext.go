package svc

import (
	"MouHu/service/app/like/api/internal/config"
	"MouHu/service/app/like/api/internal/middleware"
	"MouHu/service/app/like/rpc/likeclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware
	Rpc    likeclient.Like
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
		Rpc:    likeclient.NewLike(zrpc.MustNewClient(c.Rpc)),
	}
}
