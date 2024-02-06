package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
	}

	return cfg.Build()
}

func main() {
	// logger, _ := zap.NewProduction()
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	url := "http://imooc.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3)
	sugar.Infof("failed to fetch URL: %s", url)

	// logger.Info("failed to fetch URL",
	// 	zap.String("url", url),
	// 	zap.Int("nums", 3))
}
