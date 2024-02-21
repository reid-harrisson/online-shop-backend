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
	storeServer := server.Echo.Group("/store")

	server.Echo.Use(middleware.Logger())

	storeServer.GET("/docs/*", echoSwagger.WrapHandler)

	apiV1 := storeServer.Group("/api/v1")

	apiV1.Use(middleware.Logger())
	apiV1.Use(middleware.Recover())

	groupProductManagement := apiV1.Group("/product")
	GroupProductManagement(server, groupProductManagement)

	groupShoppingCart := apiV1.Group("/cart")
	GroupShoppingCart(server, groupShoppingCart)

	groupCategory := apiV1.Group("/category")
	GroupCategory(server, groupCategory)

	groupProductReviews := apiV1.Group("/review")
	GroupProductReviews(server, groupProductReviews)

	groupOrderManagement := apiV1.Group("/order")
	GroupOrderManagement(server, groupOrderManagement)

	groupInventoryManagement := apiV1.Group("/inventory")
	GroupInventoryManagement(server, groupInventoryManagement)

	groupGeneralStoreOffering := apiV1.Group("/store")
	GroupGeneralStoreOffering(server, groupGeneralStoreOffering)

	groupSalesMetrics := apiV1.Group("/analytic/sales")
	GroupSalesMetrices(server, groupSalesMetrics)

	groupTaxSettings := apiV1.Group("/tax")
	GroupTaxSettings(server, groupTaxSettings)

	groupShippingOption := apiV1.Group("/shipping-option")
	GroupShippingOption(server, groupShippingOption)
}

func GroupProductManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductManagement(server)
	e.POST("", handler.Create)
	e.POST("/attribute/:id", handler.CreateAttributes)
	e.POST("/shipping/:id", handler.CreateShippingData)
	e.GET("", handler.ReadAll)
	e.GET("/:id", handler.ReadByID)
	e.GET("/paging", handler.ReadPaging)
	e.PUT("/:id", handler.Update)
	e.PUT("/attribute/:id", handler.UpdateAttributes)
	e.PUT("/category/:id", handler.UpdateCategories)
	e.PUT("/channel/:id", handler.UpdateRelatedChannels)
	e.PUT("/content/:id", handler.UpdateRelatedContents)
	e.PUT("/min-stock-level/:id", handler.UpdateMinimumStockLevel)
	e.PUT("/shipping/:id", handler.UpdateShippingData)
	e.PUT("/tag/:id", handler.UpdateTags)
	e.PUT("/variation/:id", handler.UpdateVariations)
	e.DELETE("/:id", handler.Delete)
	e.DELETE("/attribute/:id", handler.UpdateAttributes)
	e.DELETE("/shipping/:id", handler.DeleteShippingData)
}

func GroupCategory(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersCategories(server)
	e.POST("", handler.CreateCategory)
	e.GET("", handler.ReadCategory)
	e.PUT("/:id", handler.UpdateCategory)
	e.DELETE("/:id", handler.DeleteCategory)
}

func GroupShoppingCart(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersShoppingCart(server)
	e.POST("", handler.Create)
	e.GET("", handler.Read)
	e.GET("/count", handler.ReadItemCount)
	e.PUT("/:id", handler.UpdateQuantity)
	e.DELETE("/:id", handler.DeleteByID)
	e.DELETE("", handler.DeleteAll)
}

func GroupProductReviews(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductReviews(server)
	e.POST("", handler.CreateReview)
	e.GET("/publish/:id", handler.ReadPublishedReviews)
	e.GET("/:id", handler.ReadAll)
	e.PUT("/moderate/:id", handler.ModerateReview)
}

func GroupOrderManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersOrderManagement(server)
	e.POST("", handler.Create)
	e.GET("/:id", handler.ReadByID)
	e.GET("/customer", handler.ReadByCustomerID)
	e.GET("/store", handler.ReadByStoreID)
	e.PUT("/status/:id", handler.UpdateStatus)
}

func GroupInventoryManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersInventoryManagement(server)
	e.PUT("/backorder/:id", handler.UpdateBackOrder)
	e.PUT("/level/:id", handler.UpdateShowStockLevelStatus)
	e.PUT("/out/:id", handler.UpdateShowOutOfStockStatus)
}

func GroupGeneralStoreOffering(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersGeneralStoreOffering(server)
	e.POST("", handler.Create)
}

func GroupSalesMetrices(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersSalesMetrics(server)
	e.GET("/revenue", handler.ReadRevenue)
	e.GET("/aov", handler.ReadAOV)
	e.GET("/product", handler.ReadSalesByProduct)
	e.GET("/category", handler.ReadSalesByCategory)
	e.GET("/clv", handler.ReadCLV)
}

func GroupTaxSettings(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersTaxSettings(server)
	e.GET("", handler.ReadTaxSetting)
}

func GroupShippingOption(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersTaxSettings(server)
	e.GET("", handler.ReadTaxSetting)
}
