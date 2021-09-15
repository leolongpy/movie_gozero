package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"movie_gozero/rpc/film/internal/db"

	"movie_gozero/rpc/film/internal/config"
	"movie_gozero/rpc/film/internal/server"
	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/film.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewFilmServiceExtServer(ctx)

	logc := logx.LogConf{
		ServiceName: "film",
		Mode:        "file",
		Path:        "..\\logs\\film\\",
	}
	logx.MustSetup(logc)
	defer logx.Close()

	db.Init(c.DataSource)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterFilmServiceExtServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
