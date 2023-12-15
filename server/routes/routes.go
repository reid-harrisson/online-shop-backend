package routes

import (
	s "OnlineStoreBackend/server"
	"OnlineStoreBackend/server/handlers"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func AuthMiddleware(server *s.Server) echo.MiddlewareFunc {
	authMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(server.Config.Auth.AccessSecret),
		SuccessHandler: func(c echo.Context) {
			secretKey := []byte(server.Config.Auth.AccessSecret)
			claims := jwt.MapClaims{}
			jwt.ParseWithClaims(strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer "), claims, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			id := int(uint64(claims["id"].(float64)))

			c.Request().Header.Set("id", strconv.Itoa(id))
		},
	})

	return authMiddleware
}
func ConfigureRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/docs/*", echoSwagger.WrapHandler)

	apiV1 := server.Echo.Group("/api/v1")

	apiV1.Use(middleware.Logger())
	apiV1.Use(middleware.Recover())

	groupProductManagement := apiV1.Group("/product")
	GroupProductManagement(server, groupProductManagement)
}

func GroupProductManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProduct(server)
	e.POST("", handler.Create)
	e.GET("", handler.ReadAll)
	e.GET("/:id", handler.ReadOne)
	e.GET("/active", handler.ReadActive)
	e.GET("/paging", handler.ReadPaging)
	e.GET("/search", handler.ReadSearch)
	e.PUT("/:id", handler.Update)
	e.DELETE("/:id", handler.Delete)
}
