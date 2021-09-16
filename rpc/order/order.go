package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"movie_gozero/rpc/order/internal/db"

	"movie_gozero/rpc/order/internal/config"
	"movie_gozero/rpc/order/internal/server"
	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOrderServiceExtServer(ctx)

	logc := logx.LogConf{
		ServiceName: "order",
		Mode:        "file",
		Path:        "..\\logs\\order\\",
	}
	logx.MustSetup(logc)
	defer logx.Close()

	db.Init(c.DataSource)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServiceExtServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
