package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductCategory struct {
	ID           uint64 `json:"id"`
	ProductID    uint64 `json:"product_id"`
	CategoryID   uint64 `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type ResponseCategoryWithChildren struct {
	ID          uint64   `json:"id"`
	StoreID     uint64   `json:"store_id"`
	Name        string   `json:"name"`
	ParentID    uint64   `json:"parent_id"`
	ChildrenIDs []uint64 `json:"children_ids"`
}

type ResponseCategory struct {
	ID       uint64 `json:"id"`
	StoreID  uint64 `json:"store_id"`
	Name     string `json:"name"`
	ParentID uint64 `json:"parent_id"`
}

func NewResponseProductCategories(c echo.Context, statusCode int, modelCategories []models.ProductCategoriesWithName) error {
	responseCategories := make([]ResponseProductCategory, 0)
	for _, modelCategory := range modelCategories {
		responseCategories = append(responseCategories, ResponseProductCategory{
			ID:           uint64(modelCategory.ID),
			ProductID:    modelCategory.ProductID,
			CategoryID:   modelCategory.CategoryID,
			CategoryName: modelCategory.CategoryName,
		})
	}
	return Response(c, statusCode, responseCategories)
}

func NewResponseCategories(c echo.Context, statusCode int, modelCategories []models.CategoriesWithChildren) error {
	responseTags := make([]ResponseCategoryWithChildren, 0)
	for _, modelCategory := range modelCategories {
		responseTags = append(responseTags, ResponseCategoryWithChildren{
			ID:          uint64(modelCategory.ID),
			StoreID:     modelCategory.StoreID,
			Name:        modelCategory.Name,
			ParentID:    modelCategory.ParentID,
			ChildrenIDs: modelCategory.ChildrenIDs,
		})
	}
	return Response(c, statusCode, responseTags)
}

func NewResponseCategory(c echo.Context, statusCode int, modelCategory models.Categories) error {
	return Response(c, statusCode, ResponseCategoryWithChildren{
		ID:       uint64(modelCategory.ID),
		StoreID:  modelCategory.StoreID,
		Name:     modelCategory.Name,
		ParentID: modelCategory.ParentID,
	})
}
