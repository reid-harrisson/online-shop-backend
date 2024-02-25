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
	GroupStoreManagement(server, groupGeneralStoreOffering)

	groupSalesMetrics := apiV1.Group("/analytic/sales")
	GroupSalesMetrices(server, groupSalesMetrics)

	groupTaxSettings := apiV1.Group("/tax")
	GroupTaxSettings(server, groupTaxSettings)

	groupShippingOption := apiV1.Group("/shipping")
	GroupShippingOption(server, groupShippingOption)

	groupCustomers := apiV1.Group("/customer")
	GroupCustomers(server, groupCustomers)

	groupVariations := apiV1.Group("/variation")
	GroupVariations(server, groupVariations)
}

func GroupProductManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductManagement(server)
	e.POST("", handler.Create)
	e.POST("/attribute/:id", handler.CreateAttributes)
	e.POST("/shipping/:id", handler.CreateShippingData)
	e.POST("/attribute-value/:id", handler.CreateAttributeValueByID)
	e.POST("/linked", handler.CreateLinkedProduct)
	e.GET("", handler.ReadAll)
	e.GET("/:id", handler.ReadByID)
	e.GET("/paging", handler.ReadPaging)
	e.GET("/linked", handler.ReadLinkedProduct)
	e.PUT("/:id", handler.Update)
	e.PUT("/approve/:id", handler.Approve)
	e.PUT("/attribute-value/:id", handler.UpdateAttributeValueByID)
	e.PUT("/attribute/:id", handler.UpdateAttributes)
	e.PUT("/category/:id", handler.UpdateCategories)
	e.PUT("/channel/:id", handler.UpdateRelatedChannels)
	e.PUT("/content/:id", handler.UpdateRelatedContents)
	e.PUT("/min-stock-level/:id", handler.UpdateMinimumStockLevel)
	e.PUT("/publish/:id", handler.Publish)
	e.PUT("/reject/:id", handler.Reject)
	e.PUT("/shipping/:id", handler.UpdateShippingData)
	e.PUT("/tag/:id", handler.UpdateTags)
	e.DELETE("/:id", handler.Delete)
	e.DELETE("/attribute-value/:id", handler.DeleteAttributeValueByID)
	e.DELETE("/attribute/:id", handler.DeleteAttributes)
	e.DELETE("/shipping/:id", handler.DeleteShippingData)
	e.DELETE("/linked/:id", handler.DeleteLinkedProduct)
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
	e.POST("/email-template", handler.CreateEmailTemplate)
	e.GET("/:id", handler.ReadByID)
	e.GET("/customer", handler.ReadByCustomerID)
	e.GET("/store", handler.ReadByStoreID)
	e.GET("/email-template/:id", handler.ReadEmailTemplateByStoreID)
	e.PUT("/status/:id", handler.UpdateStatus)
	e.PUT("/billing-address/:id", handler.UpdateBillingAddress)
	e.PUT("/shipping-address/:id", handler.UpdateShippingAddress)
	e.PUT("/email-template/:id", handler.UpdateEmailTemplate)
	e.DELETE("/email-template/:id", handler.DeleteEmailTemplate)
}

func GroupInventoryManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersInventoryManagement(server)
	e.PUT("/backorder/:id", handler.UpdateBackOrderStatus)
	e.PUT("/stock-level/:id", handler.UpdateShowStockLevelStatus)
	e.PUT("/out-of-stock/:id", handler.UpdateShowOutOfStockStatus)
}

func GroupStoreManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersStoreManagement(server)
	e.POST("", handler.Create)
	e.GET("", handler.ReadAll)
	e.PUT("/:id", handler.Update)
	e.DELETE("/:id", handler.Delete)
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
	handler := handlers.NewHandlersShippingOptions(server)
	e.POST("/store", handler.CreateShippingOption)
	e.GET("/store", handler.ReadShippingOption)
	e.PUT("/order", handler.UpdateShippingMethod)
}

func GroupCustomers(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersCustomers(server)
	e.POST("/address", handler.CreateCustomerAddress)
	e.GET("/address", handler.ReadCustomerAddress)
	e.PUT("/address/:id", handler.UpdateCustomerAddress)
}

func GroupVariations(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductVariations(server)
	e.POST("", handler.Create)
	e.GET("/store", handler.ReadVariationsInStore)
	e.GET("/product", handler.ReadVariationsInProduct)
	e.PUT("/stock-level/:id", handler.UpdateStockLevel)
	e.DELETE("/:id", handler.Delete)
}
