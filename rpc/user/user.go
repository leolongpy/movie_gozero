package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"movie_gozero/rpc/user/internal/db"
	"time"

	"movie_gozero/rpc/user/internal/config"
	"movie_gozero/rpc/user/internal/server"
	"movie_gozero/rpc/user/internal/svc"
	"movie_gozero/rpc/user/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServiceServer(ctx)
	fileName := time.Now().Format("20060102")
	logc := logx.LogConf{
		ServiceName: "user",
		Mode:        "file",
		Path:        "..\\logs\\user\\" + fileName,
	}
	logx.MustSetup(logc)
	defer logx.Close()

	db.Init(c.DataSource)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServiceServer(grpcServer, srv)
	})
	defer s.Stop()

	logx.Info(fmt.Sprintf("Starting rpc server at %s...", c.ListenOn))
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
