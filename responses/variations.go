package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseProductVariation struct {
	ID             uint64  `json:"id"`
	Sku            string  `json:"sku"`
	ProductID      uint64  `json:"product_id"`
	Price          float64 `json:"price"`
	StockLevel     float64 `json:"stock_level"`
	DiscountAmount float64 `json:"discount_amount"`
	DiscountType   string  `json:"discount_type"`
}

type ResponseProductVariationsInStore struct {
	ProductID         uint64                     `json:"product_id"`
	Title             string                     `json:"title"`
	StockLevel        float64                    `json:"stock_level"`
	MinimumStockLevel float64                    `json:"minimum_stock_level"`
	Variations        []ResponseProductVariation `json:"variations"`
}

type ResponseProductVariationWithAttribute struct {
	AttributeValueID uint64 `json:"attribute_value_id"`
	AttributeName    string `json:"attribute_name"`
	AttributeValue   string `json:"attribute_value"`
}
type ResponseProductVariationsInProduct struct {
	ResponseProductVariation
	Attributes []ResponseProductVariationWithAttribute `json:"attributes"`
}

func NewResponseProductVariation(c echo.Context, statusCode int, modelVar models.ProductVariations) error {
	price := modelVar.Price
	switch modelVar.DiscountType {
	case utils.PercentageOff:
		price = price - price*modelVar.DiscountAmount/100
	case utils.FixedAmountOff:
		price = price - modelVar.DiscountAmount
	}
	return Response(c, statusCode, ResponseProductVariation{
		ID:             uint64(modelVar.ID),
		Sku:            modelVar.Sku,
		ProductID:      modelVar.ProductID,
		Price:          price,
		StockLevel:     modelVar.StockLevel,
		DiscountAmount: modelVar.DiscountAmount,
		DiscountType:   utils.DiscountTypeToString(modelVar.DiscountType),
	})
}

func NewResponseProductVariationsInStore(c echo.Context, statusCode int, modelVars []models.ProductVariationsInStore) error {
	mapVars := make(map[uint64][]int)
	responseProducts := make([]ResponseProductVariationsInStore, 0)
	for index, modelVar := range modelVars {
		mapVars[modelVar.ProductID] = append(mapVars[modelVar.ProductID], index)
	}
	for productID, indexes := range mapVars {
		responseVars := make([]ResponseProductVariation, 0)
		minimumStockLevel := float64(0)
		title := ""
		stockLevel := float64(0)
		for _, index := range indexes {
			if title == "" {
				minimumStockLevel = modelVars[index].MinimumStockLevel
				title = modelVars[index].Title
			}
			stockLevel += modelVars[index].StockLevel
			price := modelVars[index].Price
			switch modelVars[index].DiscountType {
			case utils.PercentageOff:
				price = price - price*modelVars[index].DiscountAmount/100
			case utils.FixedAmountOff:
				price = price - modelVars[index].DiscountAmount
			}
			responseVars = append(responseVars, ResponseProductVariation{
				ID:             uint64(modelVars[index].ID),
				Sku:            modelVars[index].Sku,
				ProductID:      modelVars[index].ProductID,
				Price:          price,
				StockLevel:     modelVars[index].StockLevel,
				DiscountAmount: modelVars[index].DiscountAmount,
				DiscountType:   utils.DiscountTypeToString(modelVars[index].DiscountType),
			})
		}
		responseProducts = append(responseProducts, ResponseProductVariationsInStore{
			MinimumStockLevel: minimumStockLevel,
			Title:             title,
			StockLevel:        stockLevel,
			ProductID:         productID,
			Variations:        responseVars,
		})
	}
	return Response(c, statusCode, responseProducts)
}

func NewResponseProductVariationsInProduct(c echo.Context, statusCode int, modelVars []models.ProductVariationsInProduct) error {
	mapVars := make(map[uint64][]int)
	for index, modelVar := range modelVars {
		mapVars[uint64(modelVar.ID)] = append(mapVars[uint64(modelVar.ID)], index)
	}
	responseVars := make([]ResponseProductVariationsInProduct, 0)
	for _, indexes := range mapVars {
		responseAttrs := make([]ResponseProductVariationWithAttribute, 0)
		for _, index := range indexes {
			responseAttrs = append(responseAttrs, ResponseProductVariationWithAttribute{
				AttributeValueID: modelVars[index].AttributeValueID,
				AttributeName:    modelVars[index].AttributeName,
				AttributeValue:   modelVars[index].AttributeValue + modelVars[index].Unit,
			})
		}
		index := 0
		if len(indexes) > 0 {
			index = indexes[0]
		}
		price := modelVars[index].Price
		switch modelVars[index].DiscountType {
		case utils.PercentageOff:
			price = price - price*modelVars[index].DiscountAmount/100
		case utils.FixedAmountOff:
			price = price - modelVars[index].DiscountAmount
		}
		responseVars = append(responseVars, ResponseProductVariationsInProduct{
			ResponseProductVariation: ResponseProductVariation{
				ID:             uint64(modelVars[index].ID),
				Sku:            modelVars[index].Sku,
				ProductID:      modelVars[index].ProductID,
				Price:          price,
				StockLevel:     modelVars[index].StockLevel,
				DiscountAmount: modelVars[index].DiscountAmount,
				DiscountType:   utils.DiscountTypeToString(modelVars[index].DiscountType),
			},
			Attributes: responseAttrs,
		})
	}
	return Response(c, statusCode, responseVars)
}
