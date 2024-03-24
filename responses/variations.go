package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseProductVariation struct {
	ID               uint64   `json:"id"`
	Sku              string   `json:"sku"`
	ProductID        uint64   `json:"product_id"`
	Price            float64  `json:"price"`
	StockLevel       float64  `json:"stock_level"`
	Title            string   `json:"title"`
	ImageUrls        []string `json:"image_urls"`
	DiscountAmount   float64  `json:"discount_amount"`
	DiscountType     string   `json:"discount_type"`
	Description      string   `json:"description"`
	BackOrderAllowed string   `json:"back_order_allowed"`
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
	imageUrls := make([]string, 0)
	json.Unmarshal([]byte(modelVar.ImageUrls), &imageUrls)
	return Response(c, statusCode, ResponseProductVariation{
		ID:               uint64(modelVar.ID),
		Sku:              modelVar.Sku,
		ProductID:        modelVar.ProductID,
		Price:            price,
		StockLevel:       modelVar.StockLevel,
		Title:            modelVar.Title,
		Description:      modelVar.Description,
		ImageUrls:        imageUrls,
		DiscountAmount:   modelVar.DiscountAmount,
		DiscountType:     utils.DiscountTypeToString(modelVar.DiscountType),
		BackOrderAllowed: utils.SimpleStatusToString(modelVar.BackOrderStatus),
	})
}

func NewResponseProductVariationsInProduct(c echo.Context, statusCode int, modelVars []models.ProductVariationsWithAttributeValue) error {
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
		imageUrls := make([]string, 0)
		json.Unmarshal([]byte(modelVars[index].ImageUrls), &imageUrls)
		responseVars = append(responseVars, ResponseProductVariationsInProduct{
			ResponseProductVariation: ResponseProductVariation{
				ID:               uint64(modelVars[index].ID),
				Sku:              modelVars[index].Sku,
				ProductID:        modelVars[index].ProductID,
				Price:            price,
				StockLevel:       modelVars[index].StockLevel,
				Description:      modelVars[index].Description,
				ImageUrls:        imageUrls,
				Title:            modelVars[index].Title,
				DiscountAmount:   modelVars[index].DiscountAmount,
				DiscountType:     utils.DiscountTypeToString(modelVars[index].DiscountType),
				BackOrderAllowed: utils.SimpleStatusToString(modelVars[index].BackOrderStatus),
			},
			Attributes: responseAttrs,
		})
	}
	return Response(c, statusCode, responseVars)
}
