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

	groupVariations := apiV1.Group("/variation")
	GroupVariations(server, groupVariations)

	groupAnalytics := apiV1.Group("/analytic")
	GroupAnalytics(server, groupAnalytics)

	groupVisitors := apiV1.Group("/visit")
	GroupVisitors(server, groupVisitors)

	groupUpload := apiV1.Group("/upload")
	GroupUpload(server, groupUpload)

	groupCoupon := apiV1.Group("/coupon")
	GroupCoupon(server, groupCoupon)

	groupCheckout := apiV1.Group("/checkout")
	GroupCheckout(server, groupCheckout)

	groupCombo := apiV1.Group("/combo")
	GroupCombo(server, groupCombo)
}

func GroupCombo(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersCombos(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.GET("", handler.ReadAll, AuthMiddleware(server))
	e.GET("/approved", handler.ReadApproved, AuthMiddleware(server))
	e.GET("/publish", handler.ReadPublished, AuthMiddleware(server))
	e.PUT("/:id", handler.Update, AuthMiddleware(server))
	e.PUT("/approve/:id", handler.UpdateApprove, AuthMiddleware(server))
	e.PUT("/reject/:id", handler.UpdateReject, AuthMiddleware(server))
	e.PUT("/publish/:id", handler.UpdatePublish, AuthMiddleware(server))
	e.DELETE("/:id", handler.Delete, AuthMiddleware(server))
}

func GroupCheckout(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersCheckoutProcess(server)
	e.POST("/address", handler.CreateAddress, AuthMiddleware(server))
	e.POST("", handler.Read, AuthMiddleware(server))
	e.POST("/combo", handler.ReadCombo, AuthMiddleware(server))
	e.GET("/address", handler.ReadAddresses, AuthMiddleware(server))
	e.GET("/coupon", handler.ReadCoupon, AuthMiddleware(server))
	e.PUT("/address/:id", handler.UpdateAddress, AuthMiddleware(server))
}

func GroupCoupon(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersCoupons(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.GET("", handler.Read, AuthMiddleware(server))
	e.PUT("/:id", handler.Update, AuthMiddleware(server))
	e.DELETE("/:id", handler.Delete, AuthMiddleware(server))
}

func GroupVisitors(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersVisitors(server)
	e.POST("", handler.Create, AuthMiddleware(server))
}

func GroupAnalytics(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersAnalytics(server)
	e.GET("/sales-report", handler.ReadSalesReports, AuthMiddleware(server))
	e.GET("/customer-insight", handler.ReadCustomerInsight, AuthMiddleware(server))
	e.GET("/stock/daily", handler.ReadStockAnalyticDaily, AuthMiddleware(server))
	e.GET("/stock/weekly", handler.ReadStockAnalyticWeekly, AuthMiddleware(server))
	e.GET("/stock/monthly", handler.ReadStockAnalyticMonthly, AuthMiddleware(server))
	e.GET("/stock/weekday", handler.ReadStockAnalyticWeekDay, AuthMiddleware(server))
	e.GET("/stock/hour", handler.ReadStockAnalyticHour, AuthMiddleware(server))
	e.GET("/stock/month", handler.ReadStockAnalyticMonth, AuthMiddleware(server))
	e.GET("/visitor", handler.ReadVisitor, AuthMiddleware(server))
	e.GET("/convention-rate", handler.ReadConventionRate, AuthMiddleware(server))
	e.GET("/abandonment", handler.ReadShoppingCartAbandonment, AuthMiddleware(server))
	e.GET("/checkout-funnel", handler.ReadCheckoutFunnelAnalytics, AuthMiddleware(server))
	e.GET("/full-funnel", handler.ReadFullFunnelAnalytics, AuthMiddleware(server))
	e.GET("/product-view", handler.ReadProductViewAnalytics, AuthMiddleware(server))
	e.GET("/repeat-rate", handler.ReadRepeatCustomerRate, AuthMiddleware(server))
	e.GET("/churn-rate", handler.ReadCustomerChurnRate, AuthMiddleware(server))
	e.GET("/top-selling", handler.ReadTopSellingProducts, AuthMiddleware(server))
	e.GET("/order-trend", handler.ReadOrderTrendAnalytics, AuthMiddleware(server))
	e.GET("/customer-location", handler.ReadCustomerDataByLocation, AuthMiddleware(server))
	e.GET("/satisfaction", handler.ReadCustomerSatisfaction, AuthMiddleware(server))
	e.GET("/loading-time", handler.ReadPageLoadingTime, AuthMiddleware(server))
	e.GET("/sales/revenue", handler.ReadRevenue, AuthMiddleware(server))
	e.GET("/sales/aov", handler.ReadAOV, AuthMiddleware(server))
	e.GET("/sales/product", handler.ReadSalesByProduct, AuthMiddleware(server))
	e.GET("/sales/category", handler.ReadSalesByCategory, AuthMiddleware(server))
	e.GET("/sales/clv", handler.ReadCLV, AuthMiddleware(server))
}

func GroupProductManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductManagement(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.POST("/attribute/:id", handler.CreateAttributes, AuthMiddleware(server))
	e.POST("/shipping/:id", handler.CreateShippingData, AuthMiddleware(server))
	e.POST("/attribute-value/:id", handler.CreateAttributeValueByID, AuthMiddleware(server))
	e.POST("/linked", handler.CreateLinkedProduct, AuthMiddleware(server))
	e.GET("", handler.ReadAll, AuthMiddleware(server))
	e.GET("/approved", handler.ReadApproved)
	e.GET("/:id", handler.ReadByID, AuthMiddleware(server))
	e.GET("/paging", handler.ReadPaging, AuthMiddleware(server))
	e.GET("/linked", handler.ReadLinkedProduct, AuthMiddleware(server))
	e.GET("/category", handler.ReadByCategory, AuthMiddleware(server))
	e.GET("/search", handler.ReadSearch, AuthMiddleware(server))
	e.PUT("/:id", handler.Update, AuthMiddleware(server))
	e.PUT("/approve/:id", handler.Approve, AuthMiddleware(server))
	e.PUT("/attribute-value/:id", handler.UpdateAttributeValueByID, AuthMiddleware(server))
	e.PUT("/attribute/:id", handler.UpdateAttributes, AuthMiddleware(server))
	e.PUT("/category/:id", handler.UpdateCategories, AuthMiddleware(server))
	e.PUT("/channel/:id", handler.UpdateRelatedChannels, AuthMiddleware(server))
	e.PUT("/content/:id", handler.UpdateRelatedContents, AuthMiddleware(server))
	e.PUT("/publish/:id", handler.Publish, AuthMiddleware(server))
	e.PUT("/reject/:id", handler.Reject, AuthMiddleware(server))
	e.PUT("/shipping/:id", handler.UpdateShippingData, AuthMiddleware(server))
	e.PUT("/tag/:id", handler.UpdateTags, AuthMiddleware(server))
	e.DELETE("/:id", handler.Delete, AuthMiddleware(server))
	e.DELETE("/attribute-value/:id", handler.DeleteAttributeValueByID, AuthMiddleware(server))
	e.DELETE("/attribute/:id", handler.DeleteAttributes, AuthMiddleware(server))
	e.DELETE("/shipping/:id", handler.DeleteShippingData, AuthMiddleware(server))
	e.DELETE("/linked/:id", handler.DeleteLinkedProduct, AuthMiddleware(server))
}

func GroupShoppingCart(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersShoppingCart(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.GET("", handler.Read, AuthMiddleware(server))
	e.PUT("/:id", handler.UpdateQuantity, AuthMiddleware(server))
	e.DELETE("/:id", handler.Delete, AuthMiddleware(server))
	e.DELETE("", handler.DeleteAll, AuthMiddleware(server))
}

func GroupProductReviews(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductReviews(server)
	e.POST("", handler.CreateReview, AuthMiddleware(server))
	e.GET("/publish/:id", handler.ReadPublishedReviews, AuthMiddleware(server))
	e.GET("/:id", handler.ReadAll, AuthMiddleware(server))
	e.PUT("/moderate/:id", handler.ModerateReview, AuthMiddleware(server))
}

func GroupOrderManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersOrderManagement(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.POST("/combo", handler.CreateCombo, AuthMiddleware(server))
	e.GET("/:id", handler.ReadByID, AuthMiddleware(server))
	e.GET("/customer", handler.ReadByCustomerID, AuthMiddleware(server))
	e.GET("/store", handler.ReadByStoreID, AuthMiddleware(server))
	e.PUT("/status/:id", handler.UpdateStatus, AuthMiddleware(server))
	e.PUT("/status", handler.UpdateOrderItemStatus, AuthMiddleware(server))
}

func GroupInventoryManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersInventoryManagement(server)
	e.GET("/:id", handler.ReadInventory, AuthMiddleware(server))
	e.PUT("/min-stock-level/:id", handler.UpdateMinimumStockLevel, AuthMiddleware(server))
	e.PUT("/stock-level/:id", handler.UpdateStockLevel, AuthMiddleware(server))
}

func GroupStoreManagement(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersStoreManagement(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.POST("/:id/category", handler.CreateCategory, AuthMiddleware(server))
	e.POST("/:id/tag", handler.CreateTag, AuthMiddleware(server))
	e.GET("/all", handler.ReadAll)
	e.GET("/user", handler.ReadByUser, AuthMiddleware(server))
	e.GET("/:id/category", handler.ReadCategory)
	e.GET("/:id/tag", handler.ReadTag)
	e.PUT("/:id", handler.Update, AuthMiddleware(server))
	e.PUT("/:id/category/:category_id", handler.UpdateCategory, AuthMiddleware(server))
	e.PUT("/:id/tag/:tag_id", handler.UpdateTag, AuthMiddleware(server))
	e.DELETE("/:id", handler.Delete, AuthMiddleware(server))
	e.DELETE("/:id/category/:category_id", handler.DeleteCategory, AuthMiddleware(server))
	e.DELETE("/:id/tag/:tag_id", handler.DeleteTag, AuthMiddleware(server))
	e.POST("/:id/template", handler.CreateTemplate, AuthMiddleware(server))
	e.GET("/:id/template", handler.ReadTemplate, AuthMiddleware(server))
	e.PUT("/:id/template/:template_id", handler.UpdateTemplate, AuthMiddleware(server))
	e.DELETE("/:id/template/:template_id", handler.DeleteTemplate, AuthMiddleware(server))
}

func GroupTaxSettings(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersTaxSettings(server)
	e.GET("", handler.ReadTaxSetting, AuthMiddleware(server))
}

func GroupShippingOptions(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersShippingOptions(server)
	e.POST("/rate", handler.CreateShippingRate, AuthMiddleware(server))
	e.GET("/rate", handler.ReadShippingRate, AuthMiddleware(server))
	e.PUT("/rate/:id", handler.UpdateShippingRate, AuthMiddleware(server))
	e.DELETE("/rate/:id", handler.DeleteShippingRate, AuthMiddleware(server))
}

func GroupVariations(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersProductVariations(server)
	e.POST("", handler.Create, AuthMiddleware(server))
	e.GET("", handler.ReadByAttributeValues)
	e.GET("/product", handler.ReadByProduct, AuthMiddleware(server))
	e.PUT("/:id", handler.Update, AuthMiddleware(server))
	e.PUT("/back-order/:id", handler.UpdateBackOrder, AuthMiddleware(server))
	e.DELETE("/:id", handler.Delete, AuthMiddleware(server))
}

func GroupUpload(server *s.Server, e *echo.Group) {
	handler := handlers.NewHandlersUpload(server)
	e.POST("/csv", handler.UploadCSV, AuthMiddleware(server))
}
