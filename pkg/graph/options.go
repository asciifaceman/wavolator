package graph

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/asciifaceman/wavolator/pkg/logging"
	"go.uber.org/zap"
)

// Options for wavolate
type Options struct {
	Filename string
	Logger   *zap.Logger
}

// Option functional API  for configuring Wavolate
type Option func(*Options)

// WithFile attaches a filepath/name to wavolate
func WithFile(file string) Option {
	return func(opt *Options) {
		rawName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		opt.Filename = fmt.Sprintf("%s.png", rawName)
	}
}

// WithLogger attaches a logger to wavolate
func WithLogger(logger *zap.Logger) Option {
	return func(opt *Options) {
		opt.Logger = logging.Child("graph")
	}
}
