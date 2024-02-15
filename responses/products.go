package responses

import (
	"OnlineStoreBackend/models"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseProduct struct {
	ID                uint64   `json:"id"`
	StoreID           uint64   `json:"store_id"`
	Title             string   `json:"title"`
	ShortDescription  string   `json:"short_description"`
	LongDescription   string   `json:"long_description"`
	ImageUrls         []string `json:"image_urls"`
	SKU               string   `json:"sku"`
	UnitPriceRegular  float64  `json:"unit_price_regular"`
	UnitPriceSale     float64  `json:"unit_price_sale"`
	MinimumStockLevel float64  `json:"minimum_stock_level"`
	MaximumStockLevel float64  `json:"maximum_stock_level"`
	StockQuantity     float64  `json:"stock_quantity"`
	Active            int8     `json:"active"`
}

type ResponseProductsPaging struct {
	Data       []ResponseProduct `json:"data"`
	TotalCount uint64            `json:"total_count"`
}

type ResponseProductWithDetail struct {
	ResponseProduct
	RelatedChannels []uint64             `json:"related_channels"`
	RelatedContents []uint64             `json:"related_contents"`
	Tags            []string             `json:"tags"`
	Categories      []string             `json:"categories"`
	Attributes      []string             `json:"attributes"`
	Variations      []map[string]string  `json:"variations"`
	ShippingData    ResponseShippingData `json:"shipping_data"`
}

func NewResponseProduct(c echo.Context, statusCode int, modelProduct models.Products) error {
	imageUrls := make([]string, 0)
	json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
	responseProduct := ResponseProduct{
		ID:               uint64(modelProduct.ID),
		StoreID:          modelProduct.StoreID,
		Title:            modelProduct.Title,
		ShortDescription: modelProduct.ShortDescription,
		LongDescription:  modelProduct.LongDescription,
		SKU:              modelProduct.SKU,
		ImageUrls:        imageUrls,
		UnitPriceRegular: modelProduct.UnitPriceRegular,
		UnitPriceSale:    modelProduct.UnitPriceSale,
		StockQuantity:    modelProduct.StockQuantity,
		Active:           modelProduct.Active,
	}
	return Response(c, statusCode, responseProduct)
}

func NewResponseProductWithDetail(c echo.Context, statusCode int, modelDetail models.ProductsWithDetail) error {
	imageUrls := make([]string, 0)
	json.Unmarshal([]byte(modelDetail.ImageUrls), &imageUrls)

	relatedChannels := make([]uint64, 0)
	for _, modelChannel := range modelDetail.RelatedChannels {
		relatedChannels = append(relatedChannels, modelChannel.ChannelID)
	}

	relatedContents := make([]uint64, 0)
	for _, modelContent := range modelDetail.RelatedContents {
		relatedContents = append(relatedContents, modelContent.ContentID)
	}

	tags := make([]string, 0)
	for _, modelTag := range modelDetail.Tags {
		tags = append(tags, modelTag.TagName)
	}

	categories := make([]string, 0)
	for _, modelCategory := range modelDetail.Categories {
		categories = append(categories, modelCategory.CategoryName)
	}

	attributes := make([]string, 0)
	for _, modelAttr := range modelDetail.Attributes {
		attributes = append(attributes, modelAttr.AttributeName)
	}

	variations := make([]map[string]string, 0)
	variationID := uint64(0)
	if len(modelDetail.Variations) > 0 {
		variationID = uint64(modelDetail.Variations[0].ID)
	}
	variation := make(map[string]string)
	for _, modelVar := range modelDetail.Variations {
		variation[modelVar.AttributeName] = modelVar.Value
		if uint64(modelVar.ID) != variationID {
			variations = append(variations, variation)
			variation = make(map[string]string)
			variationID = uint64(modelVar.ID)
		}
	}
	if len(variation) > 0 {
		variations = append(variations, variation)
	}

	return Response(c, statusCode, ResponseProductWithDetail{
		ResponseProduct: ResponseProduct{
			ID:                uint64(modelDetail.ID),
			StoreID:           modelDetail.StoreID,
			Title:             modelDetail.Title,
			ShortDescription:  modelDetail.ShortDescription,
			LongDescription:   modelDetail.LongDescription,
			ImageUrls:         imageUrls,
			SKU:               modelDetail.SKU,
			UnitPriceRegular:  modelDetail.UnitPriceRegular,
			UnitPriceSale:     modelDetail.UnitPriceSale,
			MinimumStockLevel: modelDetail.MinimumStockLevel,
			StockQuantity:     modelDetail.StockQuantity,
			Active:            modelDetail.Active,
		},
		RelatedChannels: relatedChannels,
		RelatedContents: relatedContents,
		Tags:            tags,
		Categories:      categories,
		Attributes:      attributes,
		Variations:      variations,
		ShippingData: ResponseShippingData{
			ID:             uint64(modelDetail.ShippingData.ID),
			Weight:         modelDetail.ShippingData.Weight,
			Width:          modelDetail.ShippingData.Width,
			Height:         modelDetail.ShippingData.Height,
			Depth:          modelDetail.ShippingData.Depth,
			Classification: modelDetail.ShippingData.Classification,
		},
	})
}

func NewResponseProducts(c echo.Context, statusCode int, modelProducts []models.Products) error {
	responseProducts := make([]ResponseProduct, 0)
	for _, modelProduct := range modelProducts {
		imageUrls := make([]string, 0)
		json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
		responseProducts = append(responseProducts, ResponseProduct{
			ID:               uint64(modelProduct.ID),
			StoreID:          modelProduct.StoreID,
			Title:            modelProduct.Title,
			ShortDescription: modelProduct.ShortDescription,
			LongDescription:  modelProduct.LongDescription,
			SKU:              modelProduct.SKU,
			ImageUrls:        imageUrls,
			UnitPriceRegular: modelProduct.UnitPriceRegular,
			UnitPriceSale:    modelProduct.UnitPriceSale,
			StockQuantity:    modelProduct.StockQuantity,
			Active:           modelProduct.Active,
		})
	}
	return Response(c, statusCode, responseProducts)
}

func NewResponseProductsPaging(c echo.Context, statusCode int, modelProducts []models.Products, totalCount uint64) error {
	responseProducts := make([]ResponseProduct, 0)
	for _, modelProduct := range modelProducts {
		imageUrls := make([]string, 0)
		json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
		responseProducts = append(responseProducts, ResponseProduct{
			ID:               uint64(modelProduct.ID),
			StoreID:          modelProduct.StoreID,
			Title:            modelProduct.Title,
			ShortDescription: modelProduct.ShortDescription,
			LongDescription:  modelProduct.LongDescription,
			SKU:              modelProduct.SKU,
			ImageUrls:        imageUrls,
			UnitPriceRegular: modelProduct.UnitPriceRegular,
			UnitPriceSale:    modelProduct.UnitPriceSale,
			StockQuantity:    modelProduct.StockQuantity,
			Active:           modelProduct.Active,
		})
	}
	return Response(c, statusCode, ResponseProductsPaging{
		Data:       responseProducts,
		TotalCount: totalCount,
	})
}
