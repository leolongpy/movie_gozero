package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"movie_gozero/rpc/cms/internal/db"
	"time"

	"movie_gozero/rpc/cms/internal/config"
	"movie_gozero/rpc/cms/internal/server"
	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/cms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCMSServiceServer(ctx)

	fileName := time.Now().Format("20060102")
	logc := logx.LogConf{
		ServiceName: "cms",
		Mode:        "file",
		Path:        "..\\logs\\cms\\" + fileName,
	}
	logx.MustSetup(logc)
	defer logx.Close()

	db.Init(c.DataSource)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCMSServiceServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
