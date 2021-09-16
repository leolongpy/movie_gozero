package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/order/internal/config"
	"movie_gozero/rpc/order/orderserviceext"
)

type ServiceContext struct {
	Config config.Config
	Order  orderserviceext.OrderServiceExt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Order:  orderserviceext.NewOrderServiceExt(zrpc.MustNewClient(c.Order)),
	}
}
