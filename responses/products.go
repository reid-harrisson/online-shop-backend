package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseProduct struct {
	ID               uint64   `json:"id"`
	StoreID          uint64   `json:"store_id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"short_description"`
	LongDescription  string   `json:"long_description"`
	ImageUrls        []string `json:"image_urls"`
	Status           string   `json:"status"`
}

type ResponseProductsPaging struct {
	Data       []ResponseProduct `json:"data"`
	TotalCount int64             `json:"total_count"`
}

type ResponseProductWithDetail struct {
	ResponseProduct
	RelatedChannels []uint64                 `json:"related_channels"`
	RelatedContents []uint64                 `json:"related_contents"`
	Tags            []string                 `json:"tags"`
	Categories      []string                 `json:"categories"`
	AttributeValues []ResponseAttributeValue `json:"attributes"`
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
		ImageUrls:        imageUrls,
		Status:           utils.ProductStatusToString(modelProduct.Status),
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

	categories := make([]string, 0)
	for _, modelCategory := range modelDetail.Categories {
		categories = append(categories, modelCategory.CategoryName)
	}

	tags := make([]string, 0)
	for _, modelTag := range modelDetail.Tags {
		tags = append(tags, modelTag.TagName)
	}

	responseValues := make([]ResponseAttributeValue, 0)
	mapValues := make(map[string][]ResponseAttributeValueItem)
	mapIndexes := make(map[string]int)
	for index, modelValue := range modelDetail.AttributeValues {
		mapValues[modelValue.AttributeName] = append(mapValues[modelValue.AttributeName], ResponseAttributeValueItem{
			AttributeValueID: uint64(modelValue.ID),
			Value:            modelValue.AttributeValue,
		})
		mapIndexes[modelValue.AttributeName] = index
	}
	for _, modelValue := range modelDetail.AttributeValues {
		if mapIndexes[modelValue.AttributeName] != -1 {
			responseValues = append(responseValues, ResponseAttributeValue{
				AttributeID:   modelValue.AttributeID,
				AttributeName: modelValue.AttributeName,
				Values:        mapValues[modelValue.AttributeName],
			})
			mapIndexes[modelValue.AttributeName] = -1
		}
	}

	return Response(c, statusCode, ResponseProductWithDetail{
		ResponseProduct: ResponseProduct{
			ID:               uint64(modelDetail.ID),
			StoreID:          modelDetail.StoreID,
			Title:            modelDetail.Title,
			ShortDescription: modelDetail.ShortDescription,
			LongDescription:  modelDetail.LongDescription,
			ImageUrls:        imageUrls,
			Status:           utils.ProductStatusToString(modelDetail.Status),
		},
		RelatedChannels: relatedChannels,
		RelatedContents: relatedContents,
		Categories:      categories,
		Tags:            tags,
		AttributeValues: responseValues,
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
			ImageUrls:        imageUrls,
			Status:           utils.ProductStatusToString(modelProduct.Status),
		})
	}
	return Response(c, statusCode, responseProducts)
}

func NewResponseProductsPaging(c echo.Context, statusCode int, modelProducts []models.Products, totalCount int64) error {
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
			ImageUrls:        imageUrls,
			Status:           utils.ProductStatusToString(modelProduct.Status),
		})
	}
	return Response(c, statusCode, ResponseProductsPaging{
		Data:       responseProducts,
		TotalCount: totalCount,
	})
}
