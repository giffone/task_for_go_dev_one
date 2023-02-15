package cli

import (
	"context"
	"fmt"
	"net/http"
	"user-reg/config"
	"user-reg/model"
	"user-reg/service"

	"github.com/go-resty/resty/v2"
)

func New(cfg *config.Cfg, cli *resty.Client) service.RestyCli {
	return &restyCli{
		cli:     cli,
		urlSalt: cfg.UrlSalt,
	}
}

type restyCli struct {
	cli     *resty.Client
	urlSalt string
}

func (c *restyCli) GetSalt(ctx context.Context) ([]byte, error) {
	resp, err := c.cli.R().
		SetContext(ctx).
		Post(c.urlSalt)
	if err != nil {
		return nil, fmt.Errorf("resty: request: %w", err)
	}
	if err := status(resp); err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func status(resp *resty.Response) error {
	if resp.IsError() {
		if resp.StatusCode() == http.StatusGatewayTimeout || resp.StatusCode() == http.StatusBadGateway {
			return model.ErrTimeOut
		}
		return fmt.Errorf("resty: response: %s", resp.String())
	}
	return nil
}
