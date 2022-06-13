package main

import (
	conf2 "Coattails"
	"Coattails/bootstrap"
	"Coattails/core/internal/config"
	"Coattails/core/internal/handler"
	"Coattails/core/internal/svc"
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/coattails-api.yaml", "the config file")

var (
	showVersion   = flag.Bool("v", false, "展示版本信息")
	isUpdate      = flag.Bool("u", false, "是否是更新")
	delUpgradeBin = flag.Bool("delUpgradeBin", false, "是否删除升级程序")
)

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

	//flagInit() //暂不需要
	err := bootstrap.InitApp() //初始化 获得管理员权限  监听客户端事件
	if err != nil {
		panic(fmt.Sprintf("初始化应用失败:%v\n", err))
	}

	//初始化
	prophet := conf2.NewProphet()
	if err = prophet.Run(); err != nil {
		log.Fatal(err)
	}
	//	lolClientApi.InitCli() //
	//
	//	lolClientApi.InitWsClient() //初始化websocket  监听客户端事件
}

//func flagInit() {
//	flag.Parse()
//	if *showVersion {
//		log.Printf("当前版本:%s,commitID:%s,构建时间:%v\n", app.APPVersion,
//			app.Commit, app.BuildTime)
//		os.Exit(0)
//		return
//	}
//	if *isUpdate {
//		err := selfUpdate()
//		if err != nil {
//			log.Println("selfUpdate failed,", err)
//		}
//		return
//	} else {
//		_ = mustRunWithMain()
//	}
//	if *delUpgradeBin {
//		go func() {
//			_ = removeUpgradeBinFile()
//		}()
//	}
//}
