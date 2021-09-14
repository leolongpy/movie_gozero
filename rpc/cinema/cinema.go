package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"movie_gozero/rpc/cinema/internal/db"
	"time"

	"movie_gozero/rpc/cinema/internal/config"
	"movie_gozero/rpc/cinema/internal/server"
	"movie_gozero/rpc/cinema/internal/svc"
	"movie_gozero/rpc/cinema/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/cinema.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCinemaServiceExtServer(ctx)

	fileName := time.Now().Format("20060102")
	logc := logx.LogConf{
		ServiceName: "cinema",
		Mode:        "file",
		Path:        "..\\logs\\cinema\\" + fileName,
	}
	logx.MustSetup(logc)
	defer logx.Close()

	db.Init(c.DataSource)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCinemaServiceExtServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
