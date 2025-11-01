package config

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/grafvonb/c8volt/toolx/logging"
)

type Log struct {
	Level           string `mapstructure:"level" json:"level" yaml:"level"`
	Format          string `mapstructure:"format" json:"format" yaml:"format"` // "text" or "json"
	WithSource      bool   `mapstructure:"with_source" json:"with_source" yaml:"with_source"`
	WithRequestBody bool   `mapstructure:"with_request_body" json:"with_request_body" yaml:"with_request_body"`
}

func (l *Log) Normalize() {
	l.Level = strings.ToLower(strings.TrimSpace(l.Level))
	l.Format = strings.ToLower(strings.TrimSpace(l.Format))
	if l.Level == "" {
		l.Level = "info"
	}
	switch l.Level {
	case "trace":
		l.Level = "debug" // collapse if your logger lacks trace
	case "warn", "warning":
		l.Level = "warn"
	case "err":
		l.Level = "error"
	}
	if l.Format == "" {
		l.Format = "text"
	}
	switch l.Format {
	case "console":
		l.Format = "text"
	}
}

func (l *Log) Validate() error {
	var errs []error
	switch l.Level {
	case "debug", "info", "warn", "error":
	default:
		errs = append(errs, fmt.Errorf("%w: %q (allowed: debug|info|warn|error)", ErrInvalidLogLevel, l.Level))
	}
	switch l.Format {
	case "text", "json", "plain":
	default:
		errs = append(errs, fmt.Errorf("%w: %q (allowed: text|json|plain)", ErrInvalidLogFormat, l.Format))
	}
	return errors.Join(errs...)
}

func (l *Log) NewLogger() *slog.Logger {
	return logging.New(logging.LoggerConfig{
		Level:           l.Level,
		Format:          l.Format,
		WithSource:      l.WithSource,
		WithRequestBody: l.WithRequestBody,
	})
}
