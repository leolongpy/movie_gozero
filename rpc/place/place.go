package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"movie_gozero/rpc/place/internal/db"

	"movie_gozero/rpc/place/internal/config"
	"movie_gozero/rpc/place/internal/server"
	"movie_gozero/rpc/place/internal/svc"
	"movie_gozero/rpc/place/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/place.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewPlaceServiceExtServer(ctx)

	logc := logx.LogConf{
		ServiceName: "place",
		Mode:        "file",
		Path:        "..\\logs\\place\\",
	}
	logx.MustSetup(logc)
	defer logx.Close()

	db.Init(c.DataSource)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPlaceServiceExtServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
