package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// LoggerConfig represents available configuration options for logger.
type LoggerConfig struct {
	Level string `default:"warn"`
}

// NewLogger returns logger instance with .
func NewLogger(config LoggerConfig) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, errors.Wrap(err, "parse level")
	}

	logger.SetLevel(level)

	return logger, nil
}
