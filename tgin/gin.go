package tgin

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools/tog/logger"
)

type Config struct {
	Mode string `yarn:"mode"` // [debug,release,test]
	Port string `yarn:"port"`
}

func NewGin(cfg Config, fn func(g *gin.Engine) error) error {
	gin.SetMode(cfg.Mode)
	e := gin.New()

	logopt := logger.OptionStdout()
	logopt.ShowLine = false
	logopt.LevelKey = ""
	logopt.TimeKey = ""
	logopt.LineEnding = "\r"
	gin.DefaultWriter = logger.NewLogger(logopt)

	//e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	return fmt.Sprintf(`{"msg":"[GIN]","client_ip":"%s","method": "%s","path":"%s","http_code": %d}`, param.ClientIP, param.Method, param.Path, param.StatusCode)
	//}))

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
