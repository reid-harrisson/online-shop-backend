package main

import (
	application "OnlineStoreBackend"
	"OnlineStoreBackend/pkgs/config"
	"OnlineStoreBackend/pkgs/logging"
)

// @Title PockitTV Online Store
// @Version 1.0
// @SecurityDefinitions.apikey ApiKeyAuth
// @Type apiKey
// @In header
// @Name Authorization

func main() {
	cfg, err := config.Load([]string{"config.yaml"}, true, nil)
	if err != nil {
		panic(err)
	}

	logger, err := logging.Configure(&cfg.Log)
	if err != nil {
		panic(err)
	}

	logger.Info("server is starting")

	application.Start(cfg)
}
