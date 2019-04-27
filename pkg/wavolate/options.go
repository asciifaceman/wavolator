package wavolate

import (
	"github.com/asciifaceman/wavolator/pkg/logging"
	"github.com/asciifaceman/wavolator/pkg/reducer"
	"go.uber.org/zap"
)

// Options for wavolate
type Options struct {
	Filename   string
	Resolution uint
	Reducer    reducer.Reducer
	Logger     *zap.Logger
}

// Option functional API  for configuring Wavolate
type Option func(*Options)

// WithFile attaches a filepath/name to wavolate
func WithFile(file string) Option {
	return func(opt *Options) {
		opt.Filename = file
	}
}

// WithResolution attaches a resolution to wavolate
func WithResolution(resolution uint) Option {
	return func(opt *Options) {
		opt.Resolution = resolution
	}
}

// WithReducer attaches a reducer to wavolate
func WithReducer(red string) Option {
	return func(opt *Options) {
		opt.Reducer = reducer.NewReducer(red)
	}
}

// WithLogger attaches a logger to wavolate
func WithLogger() Option {
	return func(opt *Options) {
		opt.Logger = logging.Child("wavolate")
	}
}
