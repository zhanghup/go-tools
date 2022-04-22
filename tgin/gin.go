package tgin

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"strings"
)

//go:embed config-default.yml
var defaultConfig []byte

type Config struct {
	Mode    string `yaml:"mode"` // [debug,release,test]
	Port    string `yaml:"port"`
	Trusted string `yaml:"trusted"`
}

func InitGin(ymlData []byte, fn func(g *gin.Engine) error) error {

	cfg := struct {
		Web Config `json:"web" yaml:"web"`
	}{}

	err := tools.ConfOfByte(defaultConfig, &cfg)
	if err != nil {
		return err
	}

	if ymlData != nil && len(ymlData) > 0 {
		err := tools.ConfOfByte(ymlData, &cfg)
		if err != nil {
			return err
		}
	}

	return NewGin(cfg.Web, fn)
}

func NewGin(cfg Config, fn func(g *gin.Engine) error) error {
	gin.SetMode(cfg.Mode)
	gin.DefaultWriter = tog.Writer

	trusted := []string{"0.0.0.0"}
	if cfg.Trusted != "" {
		trusted = strings.Split(cfg.Trusted, ",")
	}

	e := gin.New()
	err := e.SetTrustedProxies(trusted)
	if err != nil {
		return err
	}
	e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(`[GIN]	%s | %s	 |	 %d |	 "%s"`+"\n", param.ClientIP, param.Method, param.StatusCode, param.Path)
	}))

	e.Use(gin.Recovery())

	err = fn(e)
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
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	Response   any    `json:"response"`
}

func (this ResponseEntity) Error() string {
	return fmt.Sprintf("GIN通用异常，Code：%d，Msg：%s", this.Code, this.Msg)
}

func (this ResponseEntity) SetStatusCode(code int) ResponseEntity {
	this.StatusCode = code
	return this
}

func NewResponseEntity(code int, msg string, response any) ResponseEntity {
	return ResponseEntity{
		StatusCode: 200,
		Code:       code,
		Msg:        msg,
		Response:   response,
	}
}

func Do(c *gin.Context, fn func(c *gin.Context) (any, error)) {
	o, err := fn(c)
	if err != nil {
		DoDirective(c, func(c *gin.Context) (any, error) {
			return o, err
		})
	} else {
		c.JSON(200, NewResponseEntity(200, "ok", o))
	}
}

func DoDirective(c *gin.Context, fn func(c *gin.Context) (any, error)) {
	o, err := fn(c)
	if err != nil {
		switch err.(type) {
		case ResponseEntity:
			e := err.(ResponseEntity)
			c.JSON(e.StatusCode, err)
			c.Abort()
		case *ResponseEntity:
			e := err.(*ResponseEntity)
			c.JSON(e.StatusCode, err)
			c.Abort()
		default:
			c.JSON(200, NewResponseEntity(200, err.Error(), o))
			c.Abort()
		}
	}
}
