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
	groupShoppingCart := apiV1.Group("/cart")
	GroupShoppingCart(server, groupShoppingCart)
	groupProductReviews := apiV1.Group("/review")
	GroupProductReviews(server, groupProductReviews)
	groupOrderManagement := apiV1.Group("/order")
	GroupOrderManagement(server, groupOrderManagement)
	groupInventoryManagement := apiV1.Group("/store")
	GroupInventoryManagement(server, groupInventoryManagement)
	groupSalesMetrics := apiV1.Group("/analytic/sales")
	GroupSalesMetrices(server, groupSalesMetrics)
}

func GroupProductManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductManagement(server)
	e.POST("", handler.Create)
	e.POST("/attribute/:id", handler.CreateAttributes)
	e.POST("/tag/:id", handler.CreateTags)
	e.POST("/channel/:id", handler.CreateRelatedChannels)
	e.POST("/content/:id", handler.CreateRelatedContents)
	e.POST("/review/:id", handler.CreateReview)
	e.POST("/shipping/:id", handler.CreateShippingData)
	e.GET("", handler.ReadAll)
	e.GET("/:id", handler.Read)
	e.GET("/detail/:id", handler.ReadDetail)
	e.GET("/paging", handler.ReadPaging)
	e.PUT("/:id", handler.Update)
	e.PUT("/linked/:id", handler.UpdateLinkedProduct)
	e.PUT("/price/:id", handler.UpdatePrice)
	e.PUT("/quantity/:id", handler.UpdateStockQuantity)
	e.DELETE("/:id", handler.Delete)
}

func GroupShoppingCart(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersShoppingCart(server)
	e.POST("", handler.Create)
	e.GET("", handler.ReadAll)
	e.GET("/preview", handler.ReadPreview)
	e.PUT("/:id", handler.UpdateQuantity)
	e.DELETE("/:id", handler.Delete)
	e.DELETE("", handler.DeleteAll)
}

func GroupProductReviews(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductReviews(server)
	e.POST("/rate", handler.CreateRate)
	e.GET("/rate", handler.ReadRate)
	e.GET("", handler.ReadReview)
	e.GET("/publish", handler.ReadPublishReview)
	e.PUT("/publish/:id", handler.Update)
	e.DELETE("/:id", handler.Delete)
}

func GroupOrderManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersOrderManagement(server)
	e.POST("", handler.Create)
	e.PUT("/status/:id", handler.UpdateStatus)
}

func GroupInventoryManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersInventoryManagement(server)
	e.POST("", handler.Create)
	e.PUT("/backorder/:id", handler.UpdateBackOrder)
	e.PUT("/tracking/:id", handler.UpdateStockTracking)
}

func GroupSalesMetrices(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersSalesMetrics(server)
	e.GET("/revenue", handler.ReadRevenue)
	e.GET("/aov", handler.ReadAOV)
	e.GET("/product", handler.ReadSalesByProduct)
	e.GET("/category", handler.ReadSalesByCategory)
	e.GET("/clv", handler.ReadCLV)
}
