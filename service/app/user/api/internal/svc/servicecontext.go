package svc

import (
	"MouHu/service/app/user/api/internal/config"
	"MouHu/service/app/user/api/internal/middleware"
	"MouHu/service/app/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware
	Rpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
		Rpc:    userclient.NewUser(zrpc.MustNewClient(c.Rpc)),
	}
}
