package app

import (
	"github.com/AndryHardiyanto/danstest/internal/service/account"
	"github.com/AndryHardiyanto/danstest/internal/service/auth"
	"github.com/AndryHardiyanto/danstest/internal/service/job"
)

type App struct {
	Services *Services
}

type Services struct {
	AuthService    auth.Service
	AccountService account.Service
	JobService     job.Service
}
