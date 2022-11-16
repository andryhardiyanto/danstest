package http

import (
	"net/http"
	"strings"

	"github.com/AndryHardiyanto/danstest/internal/model/app"
	"github.com/AndryHardiyanto/danstest/internal/model/auth"
	"github.com/AndryHardiyanto/danstest/lib/errors"
	"github.com/gin-gonic/gin"
)

func Middleware(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getToken(c.Request)
		if token == "" {
			c.Error(errors.NewError("validate empty token").SetType(errors.TypeUnauthorized))
			c.Abort()
			return
		}

		err := app.Services.AuthService.Validate(c.Request.Context(), &auth.ValidateRequest{
			Token: token,
		})
		if err != nil {
			c.Error(errors.NewWrapError(err, "error validate token"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// getToken .
func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) < 2 {
		return ""
	}

	token = strings.Trim(splitToken[1], " ")
	return token
}
