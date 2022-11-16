package handler

import (
	modelAccount "github.com/AndryHardiyanto/danstest/internal/model/account"
	"github.com/AndryHardiyanto/danstest/internal/model/app"
	"github.com/AndryHardiyanto/danstest/lib/errors"
	libGin "github.com/AndryHardiyanto/danstest/lib/gin"
	"github.com/AndryHardiyanto/danstest/lib/response"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func Login(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *modelAccount.LoginRequest

		err := libGin.NewRequest(c).GetBody(&req)
		if err != nil {
			c.Error(errors.NewWrapError(err, "error GetBody Login"))
			c.Abort()
			return
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			c.Error(errors.NewWrapError(err, "error validate struct").SetType(errors.TypeRequestCannotEmpty))
			c.Abort()
			return
		}

		resp, err := app.Services.AccountService.Login(c.Request.Context(), req)
		if err != nil {
			c.Error(errors.NewWrapError(err, "error Login"))
			c.Abort()
			return
		}

		response.New(c).Ok(resp)
	}
}
