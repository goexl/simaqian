package loki

import (
	"github.com/goexl/simaqian/internal/internal"
)

type Config struct {
	Url      string
	Batch    *internal.Batch
	Labels   map[string]string
	Username string
	Password string
}
