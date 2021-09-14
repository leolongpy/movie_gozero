package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"time"

	"movie_gozero/api/cms/internal/config"
	"movie_gozero/api/cms/internal/handler"
	"movie_gozero/api/cms/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/cms-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	fileName := time.Now().Format("20060102")
	logc := logx.LogConf{
		ServiceName: "cms",
		Mode:        "file",
		Path:        "..\\logs\\cms\\" + fileName,
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
