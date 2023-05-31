package svc

import (
	"MouHu/service/app/qa/api/internal/config"
	"MouHu/service/app/qa/api/internal/middleware"
	"MouHu/service/app/qa/rpc/qaclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware
	Rpc    qaclient.Qa
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
		Rpc:    qaclient.NewQa(zrpc.MustNewClient(c.Rpc)),
	}
}
