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

type ResponseStoreCategory struct {
	ID          uint64   `json:"id"`
	StoreID     uint64   `json:"store_id"`
	Name        string   `json:"name"`
	ParentID    uint64   `json:"parent_id"`
	ChildrenIDs []uint64 `json:"children_ids"`
}

func NewResponseStoreCategories(c echo.Context, statusCode int, modelCategories []models.StoreCategoriesWithChildren) error {
	responseTags := make([]ResponseStoreCategory, 0)
	for _, modelCategory := range modelCategories {
		responseTags = append(responseTags, ResponseStoreCategory{
			ID:          uint64(modelCategory.ID),
			StoreID:     modelCategory.StoreID,
			Name:        modelCategory.Name,
			ParentID:    modelCategory.ParentID,
			ChildrenIDs: modelCategory.ChildrenIDs,
		})
	}
	return Response(c, statusCode, responseTags)
}
