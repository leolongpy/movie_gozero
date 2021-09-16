package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/place/internal/config"
	"movie_gozero/rpc/place/placeserviceext"
)

type ServiceContext struct {
	Config config.Config
	Place  placeserviceext.PlaceServiceExt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Place:  placeserviceext.NewPlaceServiceExt(zrpc.MustNewClient(c.Place)),
	}
}
