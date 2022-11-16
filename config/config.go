package config

import (
	"time"

	libEnv "github.com/AndryHardiyanto/danstest/lib/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   Server
	Jwt      Jwt
	Logger   Logger
	Database Database
	Client   Client
}
type Jwt struct {
	SignedSecret       string
	AccessExpDuration  time.Duration
	RefreshExpDuration time.Duration
}

type Server struct {
	Port    string
	GinMode string
}
type Logger struct {
	Name  string
	Debug bool
}

type Database struct {
	DatabaseConnection string
}
type Client struct {
	Host string
}

var Cfg Config

func RegisterConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	Cfg = Config{
		Server: Server{
			Port:    libEnv.GetStringOrDefault("DANS_SERVER_PORT", ""),
			GinMode: libEnv.GetStringOrDefault("DANS_SERVER_GIN_MODE", ""),
		},
		Database: Database{
			DatabaseConnection: libEnv.GetStringOrDefault("DANS_DATABASE_CONNECTION", ""),
		},
		Jwt: Jwt{
			SignedSecret:       libEnv.GetStringOrDefault("DANS_JWT_SIGNED_SECRET", ""),
			AccessExpDuration:  libEnv.GetTimeDurationInHourOrDefault("DANS_JWT_ACCESS_EXP_DURATION", 0),
			RefreshExpDuration: libEnv.GetTimeDurationInHourOrDefault("DANS_JWT_REFRESH_EXP_DURATION", 0),
		},
		Logger: Logger{
			Debug: libEnv.GetBoolOrDefault("DANS_LOGGER_DEBUD", false),
		},
		Client: Client{
			Host: libEnv.GetStringOrDefault("DANS_API_DANS_MULTI_PRO", ""),
		},
	}
}
