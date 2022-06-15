package main

import (
	"Coattails"
	"Coattails/bootstrap"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var configFile = flag.String("f", "etc/coattails-api.yaml", "the config file")

var (
	showVersion   = flag.Bool("v", false, "展示版本信息")
	isUpdate      = flag.Bool("u", false, "是否是更新")
	delUpgradeBin = flag.Bool("delUpgradeBin", false, "是否删除升级程序")
)

func flagInit() {
	flag.Parse()

	_ = mustRunWithMain()
}

func mustRunWithMain() error {
	binPath, err := os.Executable() //返回当前进程的工作目录
	if err != nil {
		return err
	}
	binFileName := filepath.Base(binPath)
	if binFileName == "hh-lol-prophet_new.exe" {
		os.Exit(-1) //当前进程以给出的状态码退出 状态码 0 表示成功，非 0 表示出错
	}
	return nil
}

func main() {
	flagInit()                 //暂不需要
	err := bootstrap.InitApp() //初始化 获得管理员权限  监听客户端事件
	if err != nil {
		panic(fmt.Sprintf("初始化应用失败:%v\n", err))
	}

	//初始化
	prophet := Coattails.NewProphet()
	if err = prophet.Run(); err != nil {
		log.Fatal(err)
	}
	//lolClientApi.InitCli() //
	//
	//lolClientApi.InitWsClient() //初始化websocket  监听客户端事件
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
