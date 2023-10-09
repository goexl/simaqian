package loki

import (
	"github.com/goexl/gox"
	"github.com/goexl/simaqian/internal/internal"
)

type Config struct {
	Url      string
	Batch    *internal.Batch
	Labels   gox.Labels
	Username string
	Password string
}
