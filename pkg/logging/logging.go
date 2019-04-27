// Package logging ...
// courtesy https://github.com/catherinetcai/gsuite-aws-sso/blob/master/pkg/logging/logging.go
// modified
package logging

import (
	"go.uber.org/zap"
)

var (
	// The global logger
	logger *zap.Logger
)

// Configure configures the zap logger for Stackdriver
func Configure() (err error) {
	config := zap.NewDevelopmentConfig()
	// This breaks my terminal for some reason
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err = config.Build()
	return err
}

// Logger returns the global instance of a logger
func Logger() *zap.Logger {
	if logger == nil {
		Configure()
	}
	return logger
}

// Child returns a child logger
func Child(name string) *zap.Logger {
	return logger.Named(name)
}

// SetLogger allows the overriding of the global logger, really
// only recommended for testing.
func SetLogger(l *zap.Logger) {
	logger = l
}
