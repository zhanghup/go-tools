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
	if cfg.Port != "" {
		err = e.Run(":" + cfg.Port)
		if err != nil {
			return err
		}
	}

	return nil
}

type ResponseEntity struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Response interface{} `json:"response"`
}

func (this ResponseEntity) Error() string {
	return fmt.Sprintf("GIN通用异常，Code：%d，Msg：%s", this.Code, this.Msg)
}

func NewResponseEntity(code int, msg string, response interface{}) ResponseEntity {
	return ResponseEntity{
		Code:     code,
		Msg:      msg,
		Response: response,
	}
}

func Do(c *gin.Context, fn func(c *gin.Context) (interface{}, error)) {
	o, err := fn(c)
	if err != nil {
		DoDirective(c, func(c *gin.Context) (interface{}, error) {
			return o, err
		})
	} else {
		c.JSON(200, NewResponseEntity(200, "ok", o))
	}
}

func DoDirective(c *gin.Context, fn func(c *gin.Context) (interface{}, error)) {
	o, err := fn(c)
	if err != nil {
		switch err.(type) {
		case ResponseEntity:
			c.JSON(200, err)
			c.Abort()
		default:
			c.JSON(200, NewResponseEntity(200, err.Error(), o))
			c.Abort()
		}
	}
}
