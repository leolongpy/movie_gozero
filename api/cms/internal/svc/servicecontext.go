package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/cms/internal/config"
	"movie_gozero/rpc/cms/cmsservice"
)

type ServiceContext struct {
	Config config.Config
	Cms    cmsservice.CMSService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cms:    cmsservice.NewCMSService(zrpc.MustNewClient(c.Cms)),
	}
}
