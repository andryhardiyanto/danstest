package handler

import (
	"github.com/AndryHardiyanto/danstest/internal/model/app"
	"github.com/AndryHardiyanto/danstest/lib/response"
	"github.com/gin-gonic/gin"
)

func Validate(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.New(c).Ok(nil)
	}
}
