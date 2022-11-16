package job

import (
	"context"
	"encoding/json"

	modelJob "github.com/AndryHardiyanto/danstest/internal/model/job"
	"github.com/AndryHardiyanto/danstest/lib/errors"
	libHttp "github.com/AndryHardiyanto/danstest/lib/http"
)

type client struct {
	httpClient libHttp.Client
	host       string
}

//go:generate mockgen -destination=mock/client.go -package=mock_apilayer . Client
type Client interface {
	List(ctx context.Context, req *modelJob.ListRequest) ([]*modelJob.Job, error)
	GetById(ctx context.Context, id string) (*modelJob.Job, error)
}

func NewClient(httpClient libHttp.Client, host string) Client {
	return &client{
		httpClient: httpClient,
		host:       host,
	}
}

func (c *client) List(ctx context.Context, req *modelJob.ListRequest) ([]*modelJob.Job, error) {
	query := ""
	if req.Description != "" {
		query += "?description=" + req.Description
	}
	if req.FullTime != "" {
		if query != "" {
			query += "?full_time=" + req.FullTime
		} else {
			query += "&full_time=" + req.FullTime
		}
	}
	if req.Location != "" {
		if query != "" {
			query += "?location=" + req.Location
		} else {
			query += "&location=" + req.Location
		}
	}
	if req.Page != "" {
		if query != "" {
			query += "?page=%d" + req.Page
		} else {
			query += "&page=%d" + req.Page
		}
	}
	url := c.host + "/api/recruitment/positions.json" + query

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	result, err := c.httpClient.SetUrl(url).SetHeader(headers).Get(ctx)
	if err != nil {
		return nil, errors.NewWrapError(err, "error httpClient Get").SetType(errors.TypeInternalServerError)
	}

	marshalByte, _ := json.Marshal(result)
	resp := []*modelJob.Job{}
	err = json.Unmarshal(marshalByte, &resp)
	if err != nil {
		return nil, errors.NewWrapError(err, "error unmarshal").SetType(errors.TypeInternalServerError)
	}

	return resp, nil
}
func (c *client) GetById(ctx context.Context, id string) (*modelJob.Job, error) {
	url := c.host + "/api/recruitment/positions/" + id

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	result, err := c.httpClient.SetUrl(url).SetHeader(headers).Get(ctx)
	if err != nil {
		return nil, errors.NewWrapError(err, "error httpClient Get").SetType(errors.TypeInternalServerError)
	}

	marshalByte, _ := json.Marshal(result)
	resp := &modelJob.Job{}
	err = json.Unmarshal(marshalByte, &resp)
	if err != nil {
		return nil, errors.NewWrapError(err, "error unmarshal").SetType(errors.TypeInternalServerError)
	}

	return resp, nil
}
