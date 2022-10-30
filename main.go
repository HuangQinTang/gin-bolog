package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"blog/pkg/setting"
	"blog/routers"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout   //请求超时
	endless.DefaultWriteTimeOut = setting.WriteTimeout //响应超时
	endless.DefaultMaxHeaderBytes = 1 << 20            //最大header长度
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
