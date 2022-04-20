package tgin_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools/tgin"
	"testing"
)

func TestStart(t *testing.T) {
	//tog.Enable(false)
	tgin.NewGin(tgin.Config{Port: "8888"}, func(g *gin.Engine) error {
		g.GET("test", func(c *gin.Context) {
			c.String(200, "123")
		})
		g.GET("test/2222", func(c *gin.Context) {
			c.String(200, "123")
		})
		return nil
	})
}

func TestErrorInfo(t *testing.T) {
	o := tgin.NewResponseEntity(200, "", nil).SetStatusCode(401)
	fmt.Println(o.StatusCode)
}
