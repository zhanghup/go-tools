package toolgin

import "github.com/gin-gonic/gin"

type Config struct {
	Port string `yarn:"port"`
}

func NewGin(cfg Config) *gin.Engine {
	return gin.Default()
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
}

func DoCustom(c *gin.Context, fn func(c *gin.Context) (interface{}, string)) {
	o, err := fn(c)
	if len(err) != 0 {
		c.JSON(200, map[string]interface{}{
			"code":     -1,
			"msg":      err,
			"response": o,
		})
	}
}