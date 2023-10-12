package loki

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/gox"
	"github.com/goexl/simaqian/internal/config"
)

type Config struct {
	Url      string
	Batch    *config.Batch
	Labels   gox.Labels
	Username string
	Password string
	Http     *resty.Client
}
