package main

import (
	"fmt"
	"os"

	"gin_zap/dao/mysql"
	"gin_zap/logger"
	"gin_zap/router"
	"gin_zap/setting"

	"go.uber.org/zap"
)

func main() {
	// 1，加载配置
	if len(os.Args) < 2 {
		panic("程序执行时必须通过命令行指定配置文件")
	}
	err := setting.Init(os.Args[1])
	if err != nil {
		panic(err)
	}


	// 2，初始化日志模块
	err = logger.Init()
	if err != nil {
		panic(err)
	}

	defer zap.L().Sync() //防止在程序panic挂掉，需要吧缓冲区的日志落盘

	// 3，数据库初始化
	// TODO 需要增加初始化代码
	err = mysql.Init()
	if err != nil {
		zap.L().Error("mysql.Init failed", zap.Error(err))
	}


	// 4，路由初始化，controller里面写业务，service里面写固定模式的东西
	r := router.Setup()

	// 5，程序启动
	fmt.Printf("\nhttp://127.0.0.1:%v\n",setting.Conf.Port)
	r.Run(fmt.Sprintf(":%v",setting.Conf.Port))
}
