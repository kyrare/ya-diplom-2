package services

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(development bool) (*Logger, error) {
	var c *zap.Logger
	var err error

	if development {
		c, err = zap.NewDevelopment()
	} else {
		c, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{c.Sugar()}, nil
}
