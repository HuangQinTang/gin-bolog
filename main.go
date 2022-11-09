package main

import (
	_ "blog/docs"
	"blog/models"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"blog/routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	setting.Setup()       //初始化配置
	logging.Setup()       //初始化日志配置
	defer logging.Close() //关闭日志句柄
	models.Setup()        //连接mysql

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout   //请求超时
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout //响应超时
	endless.DefaultMaxHeaderBytes = 1 << 20                          //最大header长度
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
