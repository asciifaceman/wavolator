package generator

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
	Label    string
	Logger   *zap.Logger
}

// Option functional API  for configuring Wavolate
type Option func(*Options)

// WithLabel attaches a label to a generator
func WithLabel(label string) Option {
	return func(opt *Options) {
		opt.Label = label
	}
}

// WithFile attaches a filepath/name to wavolate
func WithFile(file string) Option {
	return func(opt *Options) {
		rawName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		opt.Filename = fmt.Sprintf("%s.csv", rawName)
	}
}

// WithLogger attaches a child logger
func WithLogger() Option {
	return func(opt *Options) {
		// This isn't really needed  this way, I was originally
		// going to inject the logger from above
		// but  didn't and just too lazy to move this into New(), rn
		opt.Logger = logging.Child("generator")
	}
}
