package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/film/internal/config"
	"movie_gozero/rpc/film/filmserviceext"
)

type ServiceContext struct {
	Config config.Config
	Film   filmserviceext.FilmServiceExt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Film:   filmserviceext.NewFilmServiceExt(zrpc.MustNewClient(c.Film)),
	}
}
