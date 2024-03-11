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

	healthHandler := handlers.NewHealthHandler(server)
	storeServer.GET("/health", healthHandler.HealthCheck)

	apiV1 := storeServer.Group("/api/v1")

	apiV1.Use(middleware.Logger())
	apiV1.Use(middleware.Recover())

	groupStoreManagement := apiV1.Group("/store")
	GroupStoreManagement(server, groupStoreManagement)

	groupProductManagement := apiV1.Group("/product")
	GroupProductManagement(server, groupProductManagement)

	groupShoppingCart := apiV1.Group("/cart")
	GroupShoppingCart(server, groupShoppingCart)

	groupProductReviews := apiV1.Group("/review")
	GroupProductReviews(server, groupProductReviews)

	groupOrderManagement := apiV1.Group("/order")
	GroupOrderManagement(server, groupOrderManagement)

	groupInventoryManagement := apiV1.Group("/inventory")
	GroupInventoryManagement(server, groupInventoryManagement)

	groupTaxSettings := apiV1.Group("/tax")
	GroupTaxSettings(server, groupTaxSettings)

	groupShippingOptions := apiV1.Group("/shipping")
	GroupShippingOptions(server, groupShippingOptions)

	groupCustomers := apiV1.Group("/customer")
	GroupCustomers(server, groupCustomers)

	groupVariations := apiV1.Group("/variation")
	GroupVariations(server, groupVariations)

	groupAnalytics := apiV1.Group("/analytic")
	GroupAnalytics(server, groupAnalytics)

	groupVisitors := apiV1.Group("/visit")
	GroupVisitors(server, groupVisitors)

	groupUpload := apiV1.Group("/upload")
	GroupUpload(server, groupUpload)
}

func GroupVisitors(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersVisitors(server)
	e.POST("", handler.Create)
}

func GroupAnalytics(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersAnalytics(server)
	e.GET("/sales-report", handler.ReadSalesReports)
	e.GET("/customer-insight", handler.ReadCustomerInsight)
	e.GET("/stock-level", handler.ReadStockLevels)
	e.GET("/visitor", handler.ReadVisitor)
	e.GET("/convention-rate", handler.ReadConventionRate)
	e.GET("/abandonment", handler.ReadShoppingCartAbandonment)
	e.GET("/checkout-funnel", handler.ReadCheckoutFunnelAnalytics)
	e.GET("/full-funnel", handler.ReadFullFunnelAnalytics)
	e.GET("/product-view", handler.ReadProductViewAnalytics)
	e.GET("/repeat-rate", handler.ReadRepeatCustomerRate)
	e.GET("/churn-rate", handler.ReadCustomerChurnRate)
	e.GET("/top-selling", handler.ReadTopSellingProducts)
	e.GET("/order-trend", handler.ReadOrderTrendAnalytics)
	e.GET("/customer-location", handler.ReadCustomerDataByLocation)
	e.GET("/satisfaction", handler.ReadCustomerSatisfaction)
	e.GET("/loading-time", handler.ReadPageLoadingTime)
	e.GET("/sales/revenue", handler.ReadRevenue)
	e.GET("/sales/aov", handler.ReadAOV)
	e.GET("/sales/product", handler.ReadSalesByProduct)
	e.GET("/sales/category", handler.ReadSalesByCategory)
	e.GET("/sales/clv", handler.ReadCLV)
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
	e.PUT("/min-stock-level/:id", handler.UpdateMinimumStockLevel)
	e.PUT("/stock-level/:id", handler.UpdateStockLevel)
}

func GroupStoreManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersStoreManagement(server)
	e.POST("", handler.Create)
	e.POST("/:id/category", handler.CreateCategory)
	e.POST("/:id/tag", handler.CreateTag)
	e.GET("", handler.Read)
	e.GET("/:id/category", handler.ReadCategory)
	e.GET("/:id/tag", handler.ReadTag)
	e.PUT("/:id", handler.Update)
	e.PUT("/:id/category/:category_id", handler.UpdateCategory)
	e.PUT("/:id/tag/:tag_id", handler.UpdateTag)
	e.DELETE("/:id", handler.Delete)
	e.DELETE("/:id/category/:category_id", handler.DeleteCategory)
	e.DELETE("/:id/tag/:tag_id", handler.DeleteTag)
}

func GroupTaxSettings(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersTaxSettings(server)
	e.GET("", handler.ReadTaxSetting)
}

func GroupShippingOptions(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersShippingOptions(server)
	// e.POST("/zone", handler.CreateShippingZone)
	// e.POST("/class", handler.CreateShippingClass)
	// e.POST("/local-pickup", handler.CreateShippingLocalPickup)
	// e.POST("/free", handler.CreateShippingFree)
	// e.POST("/flat-rate", handler.CreateShippingFlatRate)
	// e.POST("/table-rate", handler.CreateShippingTableRate)
	e.POST("/rate", handler.CreateShippingRate)
	// e.GET("", handler.ReadAllShippingMethod)
	e.GET("/rate", handler.ReadShippingRate)
	// e.GET("/free/:id", handler.ReadShippingFree)
	// e.GET("/local-pickup/:id", handler.ReadShippingLocalPickup)
	// e.GET("/flat-rate/:id", handler.ReadShippingFlatRate)
	// e.GET("/table-rate/:id", handler.ReadShippingTableRate)
	// e.PUT("/order", handler.UpdateShippingMethod)
	// e.PUT("/class/:id", handler.UpdateShippingClass)
	// e.PUT("/zone/:id", handler.UpdateShippingZone)
	// e.PUT("/free/:id", handler.UpdateShippingFree)
	// e.PUT("/local-pickup/:id", handler.UpdateShippingLocalPickup)
	// e.PUT("/flat-rate/:id", handler.UpdateShippingFlatRate)
	// e.PUT("/table-rate/:id", handler.UpdateShippingTableRate)
	e.PUT("/rate/:id", handler.UpdateShippingRate)
	e.DELETE("/rate/:id", handler.DeleteShippingRate)
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
	e.PUT("/:id", handler.Update)
	e.PUT("/back-order/:id", handler.UpdateBackOrder)
	e.DELETE("/:id", handler.Delete)
}

func GroupUpload(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersUpload(server)
	e.POST("/csv", handler.UploadCSV)
}
