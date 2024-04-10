package server

import (
	"OnlineStoreBackend/db"
	"OnlineStoreBackend/pkgs/config"
	"io"
	"net/http"
	"os"

	_ "OnlineStoreBackend/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
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
	serverCrt := "certificate/pockittv.com.crt"
	serverKey := "certificate/pockittv.com.key"

	if server.Config.Log.Server.Path != "" {
		logfile, _ := os.OpenFile(server.Config.Log.Server.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		multiWriter := io.MultiWriter(logfile, os.Stdout)
		server.Echo.Logger.SetOutput(multiWriter)
		server.Echo.Logger.SetLevel(log.INFO)
	}

	return server.Echo.StartTLS(":"+addr, serverCrt, serverKey)
}
