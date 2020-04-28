package tgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools/tog/logger"
	"time"
)

type Config struct {
	Port string `yarn:"port"`
}

func NewGin(cfg Config, fn func(g *gin.Engine) error) error {
	e := gin.Default()
	logopt := logger.OptionStdout()
	gin.DefaultWriter = logger.NewLogger(logopt)

	e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

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

func Do(c *gin.Context, fn func(c *gin.Context) (interface{}, string)) {
	o, err := fn(c)
	if len(err) != 0 {
		c.JSON(200, map[string]interface{}{
			"code":     -1,
			"msg":      err,
			"response": o,
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"code":     0,
			"msg":      "ok",
			"response": o,
		})
	}
	c.Abort()
}

func DoCustom(c *gin.Context, fn func(c *gin.Context) (interface{}, string)) {
	o, err := fn(c)
	if len(err) != 0 {
		c.JSON(200, map[string]interface{}{
			"code":     -1,
			"msg":      err,
			"response": o,
		})
		c.Abort()
	}
}
