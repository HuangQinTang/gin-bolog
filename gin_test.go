package gin_test

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGin (t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message":"pong",
		})
	})
	r.Run()
}
