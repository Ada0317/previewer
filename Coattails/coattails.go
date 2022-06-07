package main

import (
	"Coattails/internal/config"
	"Coattails/internal/handler"
	"Coattails/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/coattails-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

	//	bootstrap.InitApp()    //初始化 获得管理员权限
	//	lolClientApi.InitCli() //
	//
	//	lolClientApi.InitWsClient() //初始化websocket  监听客户端事件
}
