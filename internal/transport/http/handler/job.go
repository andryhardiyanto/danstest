package handler

import (
	"github.com/AndryHardiyanto/danstest/internal/model/app"
	modelJob "github.com/AndryHardiyanto/danstest/internal/model/job"
	"github.com/AndryHardiyanto/danstest/lib/errors"
	libGin "github.com/AndryHardiyanto/danstest/lib/gin"
	"github.com/AndryHardiyanto/danstest/lib/response"
	"github.com/gin-gonic/gin"
)

func ListJob(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *modelJob.ListRequest

		err := libGin.NewRequest(c).GetJsonQuery(&req)
		if err != nil {
			c.Error(errors.NewWrapError(err, "error GetBody ListJob"))
			c.Abort()
			return
		}

		resp, err := app.Services.JobService.List(c.Request.Context(), req)
		if err != nil {
			c.Error(errors.NewWrapError(err, "error ListJob"))
			c.Abort()
			return
		}

		response.New(c).Ok(resp)
	}
}
func GetJobByID(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := libGin.NewRequest(c).GetParams()

		resp, err := app.Services.JobService.GetById(c.Request.Context(), param.ByName("jobid"))
		if err != nil {
			c.Error(errors.NewWrapError(err, "error GetById"))
			c.Abort()
			return
		}

		response.New(c).Ok(resp)
	}
}
