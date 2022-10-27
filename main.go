package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"blog/pkg/setting"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,  //允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout, //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,              //请求头的最大字节数
	}

	s.ListenAndServe()
}
