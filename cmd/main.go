package main

import (
	application "PockitGolangBoilerplate"
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/logging"
)

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

	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
