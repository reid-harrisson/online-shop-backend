package application

import (
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/server"
	"PockitGolangBoilerplate/server/routes"
	"log"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
