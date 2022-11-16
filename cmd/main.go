package main

import (
	"net/http"

	"github.com/AndryHardiyanto/danstest/config"
	clientJob "github.com/AndryHardiyanto/danstest/internal/client/job"
	"github.com/AndryHardiyanto/danstest/internal/model/app"
	repoAccount "github.com/AndryHardiyanto/danstest/internal/repository/account"
	"github.com/AndryHardiyanto/danstest/internal/service/account"
	"github.com/AndryHardiyanto/danstest/internal/service/auth"
	"github.com/AndryHardiyanto/danstest/internal/service/job"
	transportHttp "github.com/AndryHardiyanto/danstest/internal/transport/http"
	libHttp "github.com/AndryHardiyanto/danstest/lib/http"
	libLog "github.com/AndryHardiyanto/danstest/lib/log"
	libPostgres "github.com/AndryHardiyanto/danstest/lib/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.RegisterConfig()
	libLog.RegisterLogger(config.Cfg.Logger.Debug)
	gin.SetMode(config.Cfg.Server.GinMode)

	db := libPostgres.NewPostgres(config.Cfg.Database.DatabaseConnection)

	authService := auth.NewService(
		config.Cfg.Jwt.SignedSecret,
		config.Cfg.Jwt.AccessExpDuration,
		config.Cfg.Jwt.RefreshExpDuration,
	)
	accountRepo := repoAccount.NewPostgres(db)
	app := &app.App{
		Services: &app.Services{
			AuthService: authService,
			AccountService: account.NewService(
				accountRepo,
				authService,
			),
			JobService: job.NewService(clientJob.NewClient(libHttp.NewClient(http.DefaultClient), config.Cfg.Client.Host)),
		},
	}

	transportHttp.RunServer(app,
		config.Cfg.Server.Port,
	)

}
