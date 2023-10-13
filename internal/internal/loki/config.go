package loki

import (
	"github.com/goexl/gox"
	"github.com/goexl/http"
	"github.com/goexl/simaqian/internal/config"
)

type Config struct {
	Url      string
	Batch    *config.Batch
	Labels   gox.Labels
	Username string
	Password string
	Http     *http.Client
}
