package account

import (
	"context"

	modelRepository "github.com/AndryHardiyanto/danstest/internal/model/repository"
	"github.com/AndryHardiyanto/danstest/lib/errors"
	libPostgres "github.com/AndryHardiyanto/danstest/lib/postgres"
)

type postgres struct {
	repo libPostgres.Postgres
}

//NewPostgres ..
func NewPostgres(repo libPostgres.Postgres) Repository {
	return &postgres{
		repo: repo,
	}
}

type Repository interface {
	Select(ctx context.Context, email string) (*modelRepository.Account, error)
}

func (p *postgres) Select(ctx context.Context, username string) (*modelRepository.Account, error) {
	query := "select ddacc_id as id, ddacc_password as password,ddacc_role as role from dd_account where ddacc_username = :username"
	dest := &modelRepository.Account{}

	found, err := p.repo.Select(query, dest, "username", username).One(ctx)
	if err != nil {
		return nil, errors.NewWrapError(err, "error select").SetType(errors.TypeInternalServerError)
	}

	if !found {
		return nil, nil
	}

	return dest, nil
}
