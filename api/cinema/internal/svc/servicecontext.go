package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/cinema/internal/config"
	"movie_gozero/rpc/cinema/cinemaserviceext"
)

type ServiceContext struct {
	Config config.Config
	Cinema cinemaserviceext.CinemaServiceExt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cinema: cinemaserviceext.NewCinemaServiceExt(zrpc.MustNewClient(c.Cinema)),
	}
}
