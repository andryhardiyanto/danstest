package job

import (
	"context"

	clientJob "github.com/AndryHardiyanto/danstest/internal/client/job"
	modelJob "github.com/AndryHardiyanto/danstest/internal/model/job"
	"github.com/AndryHardiyanto/danstest/lib/errors"
)

type service struct {
	jobClient clientJob.Client
}

//go:generate mockgen -destination=mock/service.go -package=mock_job . Service
type Service interface {
	List(ctx context.Context, req *modelJob.ListRequest) ([]*modelJob.ListResponse, error)
	GetById(ctx context.Context, id string) (*modelJob.GetByIdResponse, error)
}

func NewService(jobClient clientJob.Client) Service {
	return &service{
		jobClient: jobClient,
	}
}
func (s *service) GetById(ctx context.Context, id string) (*modelJob.GetByIdResponse, error) {
	data, err := s.jobClient.GetById(ctx, id)
	if err != nil {
		return nil, errors.NewWrapError(err, "error jobClient GetById")
	}
	if data == nil {
		return nil, errors.NewError("job not found").SetType(errors.TypeNotFound)
	}

	return &modelJob.GetByIdResponse{
		Job: data,
	}, nil
}

func (s *service) List(ctx context.Context, req *modelJob.ListRequest) ([]*modelJob.ListResponse, error) {
	data, err := s.jobClient.List(ctx, req)
	if err != nil {
		return nil, errors.NewWrapError(err, "error jobClient List")
	}

	if data == nil {
		return []*modelJob.ListResponse{}, nil
	}

	result := make([]*modelJob.ListResponse, 0)

	for _, d := range data {
		result = append(result, &modelJob.ListResponse{
			Job: d,
		})
	}
	return result, nil
}
