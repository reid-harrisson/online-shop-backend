package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProduct struct {
	ID               uint64            `json:"id"`
	CompanyID        uint64            `json:"company_id"`
	UserID           uint64            `json:"user_id"`
	Name             string            `json:"name"`
	Brief            string            `json:"brief"`
	Description      string            `json:"description"`
	SKU              string            `json:"sku"`
	ImageUrls        []string          `json:"image_urls"`
	Tags             string            `json:"tags"`
	ContentIDs       string            `json:"content_ids"`
	ChannelIDs       string            `json:"channel_ids"`
	UnitPriceRegular float64           `json:"price_regular"`
	UnitPriceSale    float64           `json:"price_sale"`
	StockQuantity    float64           `json:"stock_quantity"`
	Attribtues       map[string]string `json:"attributes"`
	ShipingInfo      ResponseShipping  `json:"shping_data"`
	LinkedProductIDs string            `json:"linked_product_ids"`
	Active           int8              `json:"active"`
	Reviews          string            `json:"reviews"`
}

type ResponseProductsPaging struct {
	Data       []ResponseProduct `json:"data"`
	TotalCount uint64            `json:"total_count"`
}

func NewResponseProduct(c echo.Context, statusCode int, modelProductDetail models.ProductDetails) error {
	responseProduct := ResponseProduct{
		ID:               uint64(modelProductDetail.ID),
		UserID:           modelProductDetail.UserID,
		CompanyID:        modelProductDetail.CompanyID,
		Name:             modelProductDetail.Name,
		Brief:            modelProductDetail.Brief,
		Description:      modelProductDetail.Description,
		SKU:              modelProductDetail.SKU,
		ImageUrls:        modelProductDetail.ImgUrls,
		Tags:             modelProductDetail.Tags,
		ContentIDs:       modelProductDetail.ContentIDs,
		ChannelIDs:       modelProductDetail.ChannelIDs,
		UnitPriceRegular: modelProductDetail.UnitPriceRegular,
		UnitPriceSale:    modelProductDetail.UnitPriceSale,
		StockQuantity:    modelProductDetail.StockQuantity,
		Attribtues:       modelProductDetail.Attribs,
		ShipingInfo: ResponseShipping{
			Weight:         modelProductDetail.ShippingInfo.Weight,
			Dimension:      modelProductDetail.ShippingInfo.Dimension,
			Classification: modelProductDetail.ShippingInfo.Classification,
		},
		LinkedProductIDs: modelProductDetail.LinkedProductIDs,
		Active:           modelProductDetail.Active,
		Reviews:          modelProductDetail.Reviews,
	}
	return Response(c, statusCode, responseProduct)
}

func NewResponseProducts(c echo.Context, statusCode int, modelProductDetails []models.ProductDetails) error {
	responseProducts := make([]ResponseProduct, 0)
	for _, modelProductDetail := range modelProductDetails {
		responseProducts = append(responseProducts, ResponseProduct{
			ID:               uint64(modelProductDetail.ID),
			UserID:           modelProductDetail.UserID,
			CompanyID:        modelProductDetail.CompanyID,
			Name:             modelProductDetail.Name,
			Brief:            modelProductDetail.Brief,
			Description:      modelProductDetail.Description,
			SKU:              modelProductDetail.SKU,
			ImageUrls:        modelProductDetail.ImgUrls,
			Tags:             modelProductDetail.Tags,
			ContentIDs:       modelProductDetail.ContentIDs,
			ChannelIDs:       modelProductDetail.ChannelIDs,
			UnitPriceRegular: modelProductDetail.UnitPriceRegular,
			UnitPriceSale:    modelProductDetail.UnitPriceSale,
			StockQuantity:    modelProductDetail.StockQuantity,
			Attribtues:       modelProductDetail.Attribs,
			ShipingInfo: ResponseShipping{
				Weight:         modelProductDetail.ShippingInfo.Weight,
				Dimension:      modelProductDetail.ShippingInfo.Dimension,
				Classification: modelProductDetail.ShippingInfo.Classification,
			},
			LinkedProductIDs: modelProductDetail.LinkedProductIDs,
			Active:           modelProductDetail.Active,
			Reviews:          modelProductDetail.Reviews,
		})
	}
	return Response(c, statusCode, responseProducts)
}

func NewResponseProductsPaging(c echo.Context, statusCode int, modelProductDetails []models.ProductDetails, totalCount uint64) error {
	responseProducts := make([]ResponseProduct, 0)
	for _, modelProductDetail := range modelProductDetails {
		responseProducts = append(responseProducts, ResponseProduct{
			ID:               uint64(modelProductDetail.ID),
			UserID:           modelProductDetail.UserID,
			CompanyID:        modelProductDetail.CompanyID,
			Name:             modelProductDetail.Name,
			Brief:            modelProductDetail.Brief,
			Description:      modelProductDetail.Description,
			SKU:              modelProductDetail.SKU,
			ImageUrls:        modelProductDetail.ImgUrls,
			Tags:             modelProductDetail.Tags,
			ContentIDs:       modelProductDetail.ContentIDs,
			ChannelIDs:       modelProductDetail.ChannelIDs,
			UnitPriceRegular: modelProductDetail.UnitPriceRegular,
			UnitPriceSale:    modelProductDetail.UnitPriceSale,
			StockQuantity:    modelProductDetail.StockQuantity,
			Attribtues:       modelProductDetail.Attribs,
			ShipingInfo: ResponseShipping{
				Weight:         modelProductDetail.ShippingInfo.Weight,
				Dimension:      modelProductDetail.ShippingInfo.Dimension,
				Classification: modelProductDetail.ShippingInfo.Classification,
			},
			LinkedProductIDs: modelProductDetail.LinkedProductIDs,
			Active:           modelProductDetail.Active,
			Reviews:          modelProductDetail.Reviews,
		})
	}
	return Response(c, statusCode, ResponseProductsPaging{
		Data:       responseProducts,
		TotalCount: totalCount,
	})
}
