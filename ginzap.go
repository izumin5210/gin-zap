package ginzap

import (
	"github.com/gin-gonic/gin"
)

func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
