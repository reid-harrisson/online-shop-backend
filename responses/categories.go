package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseCategory struct {
	ID           uint64 `json:"id"`
	ProductID    uint64 `json:"product_id"`
	CategoryID   uint64 `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func NewResponseProductCategories(c echo.Context, statusCode int, modelCategories []models.ProductCategoriesWithName) error {
	responseTags := make([]ResponseCategory, 0)
	for _, modelTag := range modelCategories {
		responseTags = append(responseTags, ResponseCategory{
			ID:           uint64(modelTag.ID),
			ProductID:    modelTag.ProductID,
			CategoryID:   modelTag.CategoryID,
			CategoryName: modelTag.CategoryName,
		})
	}
	return Response(c, statusCode, responseTags)
}
