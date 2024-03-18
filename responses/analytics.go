package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
)

type ResponseSalesRevenue struct {
	StoreID uint64  `json:"store_id"`
	Revenue float64 `json:"revenue"`
}

type ResponseSalesAOV struct {
	StoreID uint64  `json:"store_id"`
	AOV     float64 `json:"aov"`
}

type ResponseProductSales struct {
	ProductID uint64  `json:"product_id"`
	Revenue   float64 `json:"revenue"`
	Quantity  float64 `json:"quantity"`
}

type ResponseSalesByProduct struct {
	StoreID  uint64                 `json:"store_id"`
	Products []ResponseProductSales `json:"products"`
}

type ResponseCategorySales struct {
	Category string  `json:"category"`
	Revenue  float64 `json:"revenue"`
	Quantity float64 `json:"quantity"`
}

type ResponseSalesByCategory struct {
	StoreID    uint64                  `json:"store_id"`
	Categories []ResponseCategorySales `json:"categories"`
}

type ResponseCustomerSales struct {
	CustomerID uint64  `json:"customer"`
	Revenue    float64 `json:"revenue"`
	Quantity   float64 `json:"quantity"`
}

type ResponseSalesCLV struct {
	StoreID uint64                  `json:"store_id"`
	CLV     []ResponseCustomerSales `json:"clv"`
}

type ResponseSalesReport struct {
	VariationID uint64  `json:"variation_id"`
	ProductID   uint64  `json:"product_id"`
	StoreID     uint64  `json:"store_id"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type ResponseCustomerInsight struct {
	MaleCount   uint64  `json:"male_count"`
	FemaleCount uint64  `json:"female_count"`
	AverageAge  float64 `json:"average_age"`
	YoungestAge uint64  `json:"youngest_age"`
	OldestAge   uint64  `json:"oldest_age"`
}

type ResponseStockLevelAnalytic struct {
	ProductID    uint64  `json:"product_id"`
	StockLevel   float64 `json:"stock_level"`
	Availability string  `json:"availability"`
}

type ResponseVisitorAnalytic struct {
	Visitor       uint64  `json:"visitors"`
	UniqueVisitor uint64  `json:"unique_visitors"`
	PageView      uint64  `json:"page_views"`
	BounceRate    float64 `json:"bounce_rate"`
}

type ResponseConventionRate struct {
	Rate float64 `json:"rate"`
}

type ResponseShoppingCartAbandonment struct {
	Rate float64 `json:"rate"`
}

type ResponseCheckoutFunnelAnalytic struct {
	Page    string `json:"step"`
	Visitor uint64 `json:"visitors"`
}

type ResponseFullFunnelAnalytic struct {
	Page    string `json:"step"`
	Visitor uint64 `json:"visitors"`
}

type ResponseProductViewAnalytic struct {
	ProductID   uint64 `json:"product_id"`
	ProductName string `json:"product_name"`
	PageView    uint64 `json:"page_views"`
	Purchase    uint64 `json:"purchase"`
}

type ResponseRepeatCustomerRate struct {
	Rate float64 `json:"rate"`
}

type ResponseCustomerChurnRate struct {
	Rate float64 `json:"rate"`
}

type ResponseTopSellingProduct struct {
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Sales       float64 `json:"sales"`
}

type ResponseOrderTrendAnalytic struct {
	Date  time.Time `json:"date"`
	Count uint64    `json:"count"`
	Sales float64   `json:"sales"`
}

type ResponseCustomerDataByLocation struct {
	Location  string `json:"location"`
	Customers uint64 `json:"customers"`
}

type ResponseCustomerSatisfaction struct {
	ProductID     uint64  `json:"product_id"`
	AverageRating float64 `json:"average_rating"`
	NPS           float64 `json:"nps"`
	Rating1       float64 `json:"rating1"`
	Rating2       float64 `json:"rating2"`
	Rating3       float64 `json:"rating3"`
	Rating4       float64 `json:"rating4"`
	Rating5       float64 `json:"rating5"`
}

type ResponsePageLoadingTime struct {
	Page        string  `json:"page"`
	AverageTime float64 `json:"average_time"`
	MaximumTime float64 `json:"maximum_time"`
	MinimumTime float64 `json:"minimum_time"`
}

func NewResponseSalesRevenue(c echo.Context, statusCode int, modelSale models.StoreSales) error {
	return Response(c, statusCode, ResponseSalesRevenue{
		StoreID: modelSale.StoreID,
		Revenue: modelSale.Price,
	})
}

func NewResponseSalesAOV(c echo.Context, statusCode int, modelSale models.StoreSales) error {
	return Response(c, statusCode, ResponseSalesAOV{
		StoreID: modelSale.StoreID,
		AOV:     modelSale.Price,
	})
}

func NewResponseSalesByProduct(c echo.Context, statusCode int, modelSales []models.ProductSales, storeID uint64) error {
	responseSales := make([]ResponseProductSales, 0)
	for _, modelOrder := range modelSales {
		responseSales = append(responseSales, ResponseProductSales{
			ProductID: modelOrder.ProductID,
			Revenue:   modelOrder.Total,
			Quantity:  modelOrder.Quantity,
		})
	}
	return Response(c, statusCode, ResponseSalesByProduct{
		StoreID:  storeID,
		Products: responseSales,
	})
}

func NewResponseSalesByCategory(c echo.Context, statusCode int, modelSales []models.CategorySales, storeID uint64) error {
	responseSales := make([]ResponseCategorySales, 0)
	for _, modelOrder := range modelSales {
		responseSales = append(responseSales, ResponseCategorySales{
			Category: modelOrder.Category,
			Revenue:  modelOrder.Total,
			Quantity: modelOrder.Quantity,
		})
	}
	return Response(c, statusCode, ResponseSalesByCategory{
		StoreID:    storeID,
		Categories: responseSales,
	})
}

func NewResponseSalesCLV(c echo.Context, statusCode int, modelSales []models.CustomerSales, storeID uint64) error {
	responseSales := make([]ResponseCustomerSales, 0)
	for _, modelOrder := range modelSales {
		responseSales = append(responseSales, ResponseCustomerSales{
			CustomerID: modelOrder.CustomerID,
			Revenue:    modelOrder.Total,
			Quantity:   modelOrder.Quantity,
		})
	}
	return Response(c, statusCode, ResponseSalesCLV{
		StoreID: storeID,
		CLV:     responseSales,
	})
}

func NewResponseSalesReports(c echo.Context, statusCode int, modelReports []models.SalesReports) error {
	responseReports := make([]ResponseSalesReport, 0)
	for _, modelReport := range modelReports {
		responseReports = append(responseReports, ResponseSalesReport{
			VariationID: modelReport.VariationID,
			ProductID:   modelReport.ProductID,
			StoreID:     modelReport.StoreID,
			Price:       modelReport.Price,
			Quantity:    modelReport.Quantity,
			TotalPrice:  modelReport.TotalPrice,
		})
	}
	return Response(c, statusCode, responseReports)
}

func NewResponseCustomerInsights(c echo.Context, statusCode int, modelInsight models.CustomerInsights) error {
	return Response(c, statusCode, ResponseCustomerInsight{
		MaleCount:   modelInsight.MaleCount,
		FemaleCount: modelInsight.FemaleCount,
		AverageAge:  modelInsight.AverageAge,
		YoungestAge: modelInsight.YoungestAge,
		OldestAge:   modelInsight.OldestAge,
	})
}

func NewResponseStockLevelAnalytics(c echo.Context, statusCode int, modelLevels []models.StockLevelAnalytics) error {
	responseReports := make([]ResponseStockLevelAnalytic, 0)
	for _, modelLevel := range modelLevels {
		responseReports = append(responseReports, ResponseStockLevelAnalytic{
			ProductID:    modelLevel.ProductID,
			StockLevel:   modelLevel.StockLevel,
			Availability: modelLevel.Availability,
		})
	}
	return Response(c, statusCode, responseReports)
}

func NewResponseVisitorAnalytic(c echo.Context, statusCode int, modelVisitor models.VisitorAnalytics) error {
	return Response(c, statusCode, ResponseVisitorAnalytic{
		Visitor:       modelVisitor.Visitor,
		UniqueVisitor: modelVisitor.UniqueVisitor,
		PageView:      modelVisitor.PageView,
		BounceRate:    modelVisitor.BounceRate,
	})
}

func NewResponseConventionRate(c echo.Context, statusCode int, modelRate models.ConventionRate) error {
	return Response(c, statusCode, ResponseConventionRate{
		Rate: modelRate.Rate,
	})
}

func NewResponseShoppingCartAbandonment(c echo.Context, statusCode int, modelRate models.ShoppingCartAbandonment) error {
	return Response(c, statusCode, ResponseShoppingCartAbandonment{
		Rate: modelRate.Rate,
	})
}

func NewResponseCheckoutFunnelAnalytics(c echo.Context, statusCode int, modelSteps []models.CheckoutFunnelAnalytics) error {
	sort.Slice(modelSteps, func(i, j int) bool {
		return modelSteps[i].Page < modelSteps[j].Page
	})
	responseSteps := make([]ResponseCheckoutFunnelAnalytic, 0)
	for _, modelStep := range modelSteps {
		responseSteps = append(responseSteps, ResponseCheckoutFunnelAnalytic{
			Page:    utils.PageTypeToString(modelStep.Page),
			Visitor: modelStep.PageView,
		})
	}
	return Response(c, statusCode, responseSteps)
}

func NewResponseFullFunnelAnalytics(c echo.Context, statusCode int, modelSteps []models.FullFunnelAnalytics) error {
	sort.Slice(modelSteps, func(i, j int) bool {
		return modelSteps[i].Page < modelSteps[j].Page
	})
	responseSteps := make([]ResponseFullFunnelAnalytic, 0)
	for _, modelStep := range modelSteps {
		responseSteps = append(responseSteps, ResponseFullFunnelAnalytic{
			Page:    utils.PageTypeToString(modelStep.Page),
			Visitor: modelStep.PageView,
		})
	}
	return Response(c, statusCode, responseSteps)
}

func NewResponseProductViewAnalytics(c echo.Context, statusCode int, modelViews []models.ProductViewAnalytics) error {
	responseViews := make([]ResponseProductViewAnalytic, 0)
	for _, modelView := range modelViews {
		responseViews = append(responseViews, ResponseProductViewAnalytic{
			ProductID:   modelView.ProductID,
			ProductName: modelView.ProductName,
			PageView:    modelView.PageView,
			Purchase:    modelView.Purchase,
		})
	}
	return Response(c, statusCode, responseViews)
}

func NewResponseRepeatCustomerRate(c echo.Context, statusCode int, modelRates []models.RepeatCustomerRates) error {
	mapCount := make(map[uint64]int64)
	for _, modelRate := range modelRates {
		if mapCount[modelRate.ProductID] == int64(modelRate.CustomerID) {
			mapCount[modelRate.ProductID] = -1
		} else if mapCount[modelRate.ProductID] != -1 {
			mapCount[modelRate.ProductID] = int64(modelRate.CustomerID)
		}
	}
	count := 0
	repeat := 0
	for _, flag := range mapCount {
		if flag == -1 {
			repeat++
		}
		count++
	}
	return Response(c, statusCode, ResponseRepeatCustomerRate{
		Rate: float64(repeat) / float64(count),
	})
}

func NewResponseCustomerChurnRate(c echo.Context, statusCode int, modelRate models.CustomerChurnRates) error {
	return Response(c, statusCode, ResponseCustomerChurnRate{
		Rate: modelRate.Rate,
	})
}

func NewResponseTopSellingProduct(c echo.Context, statusCode int, modelProducts []models.TopSellingProducts) error {
	responseProducts := make([]ResponseTopSellingProduct, 0)
	for _, modelProduct := range modelProducts {
		responseProducts = append(responseProducts, ResponseTopSellingProduct{
			ProductID:   modelProduct.ProductID,
			ProductName: modelProduct.ProductName,
			Sales:       modelProduct.Sales,
		})
	}
	return Response(c, statusCode, responseProducts)
}

func NewResponseOrderTrendAnalytics(c echo.Context, statusCode int, modelTrends []models.OrderTrendAnalytics) error {
	responseTrends := make([]ResponseOrderTrendAnalytic, 0)
	for _, modelTrend := range modelTrends {
		responseTrends = append(responseTrends, ResponseOrderTrendAnalytic{
			Date:  modelTrend.Date,
			Count: modelTrend.Count,
			Sales: modelTrend.Sales,
		})
	}
	return Response(c, statusCode, responseTrends)
}

func NewResponseCustomerDataByLocation(c echo.Context, statusCode int, modelLocations []models.CustomerDataByLocation) error {
	responseLocations := make([]ResponseCustomerDataByLocation, 0)
	for _, modelLocation := range modelLocations {
		responseLocations = append(responseLocations, ResponseCustomerDataByLocation{
			Location:  modelLocation.Location,
			Customers: modelLocation.Customers,
		})
	}
	return Response(c, statusCode, responseLocations)
}

func NewResponseCustomerSatisfaction(c echo.Context, statusCode int, modelRates []models.CustomerSatisfaction) error {
	responseRates := make([]ResponseCustomerSatisfaction, 0)
	for _, modelRate := range modelRates {
		responseRates = append(responseRates, ResponseCustomerSatisfaction{
			ProductID:     modelRate.ProductID,
			AverageRating: utils.Round(modelRate.AverageRating),
			NPS:           utils.Round(modelRate.NPS * 100),
			Rating1:       utils.Round(modelRate.Rating1),
			Rating2:       utils.Round(modelRate.Rating2),
			Rating3:       utils.Round(modelRate.Rating3),
			Rating4:       utils.Round(modelRate.Rating4),
			Rating5:       utils.Round(modelRate.Rating5),
		})
	}
	return Response(c, statusCode, responseRates)
}

func NewResponsePageLoadingTime(c echo.Context, statusCode int, modelSteps []models.PageLoadingTime) error {
	sort.Slice(modelSteps, func(i, j int) bool {
		return modelSteps[i].Page < modelSteps[j].Page
	})
	responseSteps := make([]ResponsePageLoadingTime, 0)
	for _, modelStep := range modelSteps {
		responseSteps = append(responseSteps, ResponsePageLoadingTime{
			Page:        utils.PageTypeToString(modelStep.Page),
			AverageTime: utils.Round(modelStep.AverageTime),
			MaximumTime: utils.Round(modelStep.MaximumTime),
			MinimumTime: utils.Round(modelStep.MinimumTime),
		})
	}
	return Response(c, statusCode, responseSteps)
}
