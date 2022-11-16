package http

import (
	"github.com/AndryHardiyanto/danstest/internal/model/app"
	"github.com/AndryHardiyanto/danstest/internal/transport/http/handler"
	libGin "github.com/AndryHardiyanto/danstest/lib/gin"
	"github.com/gin-gonic/gin"
)

func Router(engineGin *gin.Engine, app *app.App) {
	engineGin.NoRoute(libGin.CustomPageNotFound())

	engineGin.Use(libGin.RegisterGinLogger())
	engineGin.Use(libGin.RegisterErrorHandler())
	engineGin.Use(libGin.Recovery())

	v1 := engineGin.Group("/v1")
	{
		v1.POST("/login", handler.Login(app))
		v1.Use(Middleware(app))
		v1.GET("/job", handler.ListJob(app))
		v1.GET("/job/:jobid", handler.GetJobByID(app))
	}

}
