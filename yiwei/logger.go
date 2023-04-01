package main

import (
	"go.uber.org/zap"
)

func CreateLogger(prod bool) *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	if prod {
		cfg = zap.NewProductionConfig()
	}

	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return log
}
