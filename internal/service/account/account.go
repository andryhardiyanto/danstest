package account

import (
	"context"
	"fmt"

	modelAccount "github.com/AndryHardiyanto/danstest/internal/model/account"
	repoAccount "github.com/AndryHardiyanto/danstest/internal/repository/account"
	serviceAuth "github.com/AndryHardiyanto/danstest/internal/service/auth"
	"github.com/AndryHardiyanto/danstest/lib/errors"
)

type service struct {
	accountRepo repoAccount.Repository
	authService serviceAuth.Service
}

type Service interface {
	Login(ctx context.Context, req *modelAccount.LoginRequest) (*modelAccount.LoginResponse, error)
}

func NewService(accountRepo repoAccount.Repository, authService serviceAuth.Service) Service {
	return &service{
		accountRepo: accountRepo,
		authService: authService,
	}
}

func (s *service) Login(ctx context.Context, req *modelAccount.LoginRequest) (*modelAccount.LoginResponse, error) {
	data, err := s.accountRepo.Select(ctx, req.Username)
	if err != nil {
		return nil, errors.NewWrapError(err, "error accountRepo Select")
	}
	if data == nil {
		return nil, errors.NewError("validation user not found").SetType(errors.TypeUserNotFound)
	}

	if data.Password != req.Password {
		return nil, errors.NewError("validation password not match").SetType(errors.TypePasswordNotMatch)
	}

	dataJwt, err := s.authService.GenerateJwt(ctx, fmt.Sprintf("%d", data.ID), data.Role)
	if err != nil {
		return nil, errors.NewWrapError(err, "error authService GenerateJwt")
	}

	return &modelAccount.LoginResponse{
		AccessToken:  dataJwt.AccessToken,
		AccessExp:    dataJwt.AccessExp,
		RefreshToken: dataJwt.RefreshToken,
		RefreshExp:   dataJwt.RefreshExp,
	}, nil
}
