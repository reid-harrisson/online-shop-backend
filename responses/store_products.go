package responses

import (
	"OnlineStoreBackend/models"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseProduct struct {
	ID               uint64   `json:"id"`
	StoreID          uint64   `json:"store_id"`
	Name             string   `json:"name"`
	Brief            string   `json:"brief"`
	Description      string   `json:"description"`
	ImageUrls        []string `json:"image_urls"`
	SKU              string   `json:"sku"`
	UnitPriceRegular float64  `json:"unit_price_regular"`
	UnitPriceSale    float64  `json:"unit_price_sale"`
	StockQuantity    float64  `json:"stock_quantity"`
	LinkedProductIDs []int    `json:"linked_product_ids"`
	Active           int8     `json:"active"`
}

type ResponseProductsPaging struct {
	Data       []ResponseProduct `json:"data"`
	TotalCount uint64            `json:"total_count"`
}

type ResponseProductDetail struct {
	ResponseProduct
	Attributes      map[string]string       `json:"attributes"`
	Tags            []string                `json:"tags"`
	RelatedChannels []string                `json:"related_channels"`
	RelatedContents []string                `json:"related_contents"`
	Reviews         []ResponseProductReview `json:"reviews"`
	ShippingData    ResponseShippingData    `json:"shipping_data"`
}

func NewResponseProduct(c echo.Context, statusCode int, modelProduct models.Products) error {
	imageUrls := make([]string, 0)
	linkedProductIDs := make([]int, 0)
	json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
	json.Unmarshal([]byte(modelProduct.LinkedProductIDs), &linkedProductIDs)
	responseProduct := ResponseProduct{
		ID:               uint64(modelProduct.ID),
		StoreID:          modelProduct.StoreID,
		Name:             modelProduct.Name,
		Brief:            modelProduct.Brief,
		Description:      modelProduct.Description,
		SKU:              modelProduct.SKU,
		ImageUrls:        imageUrls,
		UnitPriceRegular: modelProduct.UnitPriceRegular,
		UnitPriceSale:    modelProduct.UnitPriceSale,
		StockQuantity:    modelProduct.StockQuantity,
		LinkedProductIDs: linkedProductIDs,
		Active:           modelProduct.Active,
	}
	return Response(c, statusCode, responseProduct)
}

func NewResponseProductDetail(c echo.Context, statusCode int, modelProduct models.ProductDetails) error {
	imageUrls := make([]string, 0)
	linkedProductIDs := make([]int, 0)
	json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
	json.Unmarshal([]byte(modelProduct.LinkedProductIDs), &linkedProductIDs)
	responseProdRevs := make([]ResponseProductReview, 0)
	for _, modelReview := range modelProduct.Reviews {
		responseProdRevs = append(responseProdRevs, ResponseProductReview{
			ID:         uint64(modelReview.ID),
			Comment:    modelReview.Comment,
			CustomerID: modelReview.CustomerID,
			ProductID:  modelReview.ProductID,
			Status:     modelReview.Status,
		})
	}
	responseShipData := ResponseShippingData{
		ID:             uint64(modelProduct.ShipData.ID),
		Weight:         modelProduct.ShipData.Weight,
		ProductID:      modelProduct.ShipData.ProductID,
		Dimension:      modelProduct.ShipData.Dimension,
		Classification: modelProduct.ShipData.Classification,
	}
	responseProduct := ResponseProduct{
		ID:               uint64(modelProduct.ID),
		StoreID:          modelProduct.StoreID,
		Name:             modelProduct.Name,
		Brief:            modelProduct.Brief,
		Description:      modelProduct.Description,
		SKU:              modelProduct.SKU,
		ImageUrls:        imageUrls,
		UnitPriceRegular: modelProduct.UnitPriceRegular,
		UnitPriceSale:    modelProduct.UnitPriceSale,
		StockQuantity:    modelProduct.StockQuantity,
		LinkedProductIDs: linkedProductIDs,
		Active:           modelProduct.Active,
	}
	return Response(c, statusCode, ResponseProductDetail{
		ResponseProduct: responseProduct,
		Attributes:      modelProduct.Attributes,
		RelatedChannels: modelProduct.RelatedChannels,
		RelatedContents: modelProduct.RelatedContents,
		Tags:            modelProduct.Tags,
		ShippingData:    responseShipData,
		Reviews:         responseProdRevs,
	})
}

func NewResponseProducts(c echo.Context, statusCode int, modelProducts []models.Products) error {
	responseProducts := make([]ResponseProduct, 0)
	for _, modelProduct := range modelProducts {
		imageUrls := make([]string, 0)
		linkedProductIDs := make([]int, 0)
		json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
		json.Unmarshal([]byte(modelProduct.LinkedProductIDs), &linkedProductIDs)
		responseProducts = append(responseProducts, ResponseProduct{
			ID:               uint64(modelProduct.ID),
			StoreID:          modelProduct.StoreID,
			Name:             modelProduct.Name,
			Brief:            modelProduct.Brief,
			Description:      modelProduct.Description,
			SKU:              modelProduct.SKU,
			ImageUrls:        imageUrls,
			UnitPriceRegular: modelProduct.UnitPriceRegular,
			UnitPriceSale:    modelProduct.UnitPriceSale,
			StockQuantity:    modelProduct.StockQuantity,
			LinkedProductIDs: linkedProductIDs,
			Active:           modelProduct.Active,
		})
	}
	return Response(c, statusCode, responseProducts)
}

func NewResponseProductsPaging(c echo.Context, statusCode int, modelProducts []models.Products, totalCount uint64) error {
	responseProducts := make([]ResponseProduct, 0)
	for _, modelProduct := range modelProducts {
		imageUrls := make([]string, 0)
		linkedProductIDs := make([]int, 0)
		json.Unmarshal([]byte(modelProduct.ImageUrls), &imageUrls)
		json.Unmarshal([]byte(modelProduct.LinkedProductIDs), &linkedProductIDs)
		responseProducts = append(responseProducts, ResponseProduct{
			ID:               uint64(modelProduct.ID),
			StoreID:          modelProduct.StoreID,
			Name:             modelProduct.Name,
			Brief:            modelProduct.Brief,
			Description:      modelProduct.Description,
			SKU:              modelProduct.SKU,
			ImageUrls:        imageUrls,
			UnitPriceRegular: modelProduct.UnitPriceRegular,
			UnitPriceSale:    modelProduct.UnitPriceSale,
			StockQuantity:    modelProduct.StockQuantity,
			LinkedProductIDs: linkedProductIDs,
			Active:           modelProduct.Active,
		})
	}
	return Response(c, statusCode, ResponseProductsPaging{
		Data:       responseProducts,
		TotalCount: totalCount,
	})
}
