package gin

import (
	"github.com/AndryHardiyanto/danstest/lib/errors"
	"github.com/gin-gonic/gin"
)

func CustomPageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Error(errors.NewError("error - page not found").SetType(errors.TypePageNotFound))
		c.Abort()
	}
}
