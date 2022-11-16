package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AndryHardiyanto/danstest/lib/errors"
)

type client struct {
	httpClient *http.Client
	headers    map[string]string
	url        string
}

//go:generate mockgen -destination=mock/client.go -package=mock_http . Client
type Client interface {
	SetUrl(url string) Client
	SetHeader(headers map[string]string) Client
	Get(ctx context.Context) (interface{}, error)
}

func NewClient(httpCLient *http.Client) Client {
	return &client{
		httpClient: httpCLient,
	}
}
func (c *client) SetUrl(url string) Client {
	c.url = url
	return c
}
func (c *client) SetHeader(headers map[string]string) Client {
	c.headers = headers
	return c
}
func (c *client) Get(ctx context.Context) (interface{}, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return nil, errors.NewWrapError(err, "error http request").SetType(errors.TypeInternalServerError)
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewWrapError(err, "error http do").SetType(errors.TypeInternalServerError)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewWrapError(err, "error ioutil ReadAll").SetType(errors.TypeInternalServerError)
	}

	var result interface{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, errors.NewWrapError(err, "error unmarshal").SetType(errors.TypeInternalServerError)
	}

	return result, nil
}
