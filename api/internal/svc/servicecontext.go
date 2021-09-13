package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/internal/config"
	"movie_gozero/rpc/cms/cmsservice"
	"movie_gozero/rpc/user/userservice"
)

type ServiceContext struct {
	Config config.Config
	User   userservice.UserService
	Cms    cmsservice.CMSService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userservice.NewUserService(zrpc.MustNewClient(c.User)),
		Cms:    cmsservice.NewCMSService(zrpc.MustNewClient(c.Cms)),
	}
}
