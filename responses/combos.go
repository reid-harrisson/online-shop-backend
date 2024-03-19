package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseComboItem struct {
	ID          uint64  `json:"id"`
	ComboID     uint64  `json:"combo_id"`
	VariationID uint64  `json:"variation_id"`
	Quantity    float64 `json:"quantity"`
}

type ResponseCombo struct {
	ID             uint64              `json:"id"`
	StoreID        uint64              `json:"store_id"`
	DiscountAmount float64             `json:"discount_amount"`
	DiscountType   string              `json:"discount_type"`
	ImageUrls      []string            `json:"image_urls"`
	Description    string              `json:"description"`
	Title          string              `json:"title"`
	Items          []ResponseComboItem `json:"items"`
	Status         string              `json:"status"`
}

func NewResponseCombo(c echo.Context, statusCode int, modelCombo models.Combos, modelItems []models.ComboItems) error {
	responseItems := []ResponseComboItem{}
	for _, modelItem := range modelItems {
		responseItems = append(responseItems, ResponseComboItem{
			ID:          uint64(modelItem.ID),
			ComboID:     modelItem.ComboID,
			VariationID: modelItem.VariationID,
			Quantity:    modelItem.Quantity,
		})
	}
	imageUrls := []string{}
	json.Unmarshal([]byte(modelCombo.ImageUrls), &imageUrls)
	return Response(c, statusCode, ResponseCombo{
		ID:             uint64(modelCombo.ID),
		StoreID:        modelCombo.StoreID,
		DiscountAmount: modelCombo.DiscountAmount,
		DiscountType:   utils.DiscountTypeToString(modelCombo.DiscountType),
		ImageUrls:      imageUrls,
		Description:    modelCombo.Description,
		Title:          modelCombo.Title,
		Items:          responseItems,
		Status:         utils.ProductStatusToString(modelCombo.Status),
	})
}

func NewResponseCombos(c echo.Context, statusCode int, modelCombos []models.Combos, modelItems []models.ComboItems) error {
	indices := map[uint64][]int{}
	for index, modelItem := range modelItems {
		indices[modelItem.ComboID] = append(indices[modelItem.ComboID], index)
	}

	responseCombos := []ResponseCombo{}
	for _, modelCombo := range modelCombos {
		responseItems := []ResponseComboItem{}
		for _, index := range indices[uint64(modelCombo.ID)] {
			responseItems = append(responseItems, ResponseComboItem{
				ID:          uint64(modelItems[index].ID),
				ComboID:     modelItems[index].ComboID,
				VariationID: modelItems[index].VariationID,
				Quantity:    modelItems[index].Quantity,
			})
		}
		imageUrls := []string{}
		json.Unmarshal([]byte(modelCombo.ImageUrls), &imageUrls)
		responseCombos = append(responseCombos, ResponseCombo{
			ID:             uint64(modelCombo.ID),
			StoreID:        modelCombo.StoreID,
			DiscountAmount: modelCombo.DiscountAmount,
			DiscountType:   utils.DiscountTypeToString(modelCombo.DiscountType),
			ImageUrls:      imageUrls,
			Description:    modelCombo.Description,
			Title:          modelCombo.Title,
			Items:          responseItems,
			Status:         utils.ProductStatusToString(modelCombo.Status),
		})
	}
	return Response(c, statusCode, responseCombos)
}
