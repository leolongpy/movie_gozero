package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"movie_gozero/api/place/internal/config"
	"movie_gozero/api/place/internal/handler"
	"movie_gozero/api/place/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/place-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logc := logx.LogConf{
		ServiceName: "order",
		Mode:        "file",
		Path:        "..\\logs\\order\\",
	}
	logx.MustSetup(logc)
	defer logx.Close()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
