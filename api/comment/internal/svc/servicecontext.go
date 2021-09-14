package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"movie_gozero/api/comment/internal/config"
	"movie_gozero/rpc/comment/commentserviceext"
)

type ServiceContext struct {
	Config  config.Config
	Comment commentserviceext.CommentServiceExt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Comment: commentserviceext.NewCommentServiceExt(zrpc.MustNewClient(c.Comment)),
	}
}
