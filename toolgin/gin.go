package toolgin

import "github.com/gin-gonic/gin"

type Config struct {
	Port string `yarn:"port"`
}

func NewGin(cfg Config) *gin.Engine {
	return gin.Default()
}
