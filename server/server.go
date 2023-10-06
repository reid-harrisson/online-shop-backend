package server

import (
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/db"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, OPTIONS")

			c.Request().Header.Set("Access-Control-Allow-Origin", "*")
			c.Request().Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Request().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, OPTIONS")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	})

	return &Server{
		Echo:   e,
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
