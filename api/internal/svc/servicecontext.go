package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/internal/config"
	"movie_gozero/rpc/user/userservice"
)

type ServiceContext struct {
	Config config.Config
	User   userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userservice.NewUserService(zrpc.MustNewClient(c.User)),
	}
}
