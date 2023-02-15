package cli

import (
	"time"
	"user-reg/config"

	"github.com/go-resty/resty/v2"
)

func NewResty(cfg *config.Cfg) *resty.Client {
	r := resty.New()
	r.SetTimeout(30 * time.Second)
	r.SetHeader("Accept", "application/json")
	r.SetHeader("Content-Type", "application/json")

	return r
}
