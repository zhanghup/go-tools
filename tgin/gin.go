package tgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools/tog"
)

type Config struct {
	Mode string `yaml:"mode"` // [debug,release,test]
	Port string `yaml:"port"`
}

func NewGin(cfg Config, fn func(g *gin.Engine) error) error {
	gin.SetMode(cfg.Mode)
	e := gin.New()

	gin.DefaultWriter = tog.Toginfo

	e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN]\t%s | %s\t |\t %d |\t \"%s\"", param.ClientIP, param.Method, param.StatusCode, param.Path)
	}))

	e.Use(gin.Recovery())

	err := fn(e)
	if err != nil {
		return err
	}
	err = e.Run(":" + cfg.Port)
	if err != nil {
		return err
	}
	return nil
}

type ResponseEntity struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Response interface{} `json:"response"`
}

func Do(c *gin.Context, fn func(c *gin.Context) (interface{}, string)) {
	o, err := fn(c)
	if len(err) != 0 {
		c.JSON(200, ResponseEntity{
			Code:     -1,
			Msg:      err,
			Response: o,
		})
	} else {
		c.JSON(200, ResponseEntity{
			Code:     0,
			Msg:      "ok",
			Response: o,
		})
	}
	c.Abort()
}

func DoCustom(c *gin.Context, fn func(c *gin.Context) (interface{}, string)) {
	o, err := fn(c)
	if len(err) != 0 {
		c.JSON(200, ResponseEntity{
			Code:     -1,
			Msg:      err,
			Response: o,
		})
		c.Abort()
	}
}
