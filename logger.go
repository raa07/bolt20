package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Level string `default:"warn"`
}

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
