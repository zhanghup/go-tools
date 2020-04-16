package tgin

import "github.com/gin-gonic/gin"

type Config struct {
	Port string `yarn:"port"`
}

func NewGin(cfg Config, fn func(g *gin.Engine) error) error {
	e := gin.Default()
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
