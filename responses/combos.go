package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseComboItem struct {
	ComboID     uint64  `gorm:"column:combo_id; type:bigint(20)"`
	VariationID uint64  `gorm:"column:variation_id; type:bigint(20)"`
	Quantity    float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

type ResponseCombo struct {
	ID             uint64              `json:"id"`
	StoreID        uint64              `json:"store_id"`
	DiscountAmount float64             `json:"discount_amount"`
	DiscountType   string              `json:"discount_type"`
	ImageUrls      string              `json:"image_urls"`
	Description    string              `json:"description"`
	Title          string              `json:"title"`
	Items          []ResponseComboItem `json:"items"`
}

func NewResponseCombo(c echo.Context, statusCode int, modelCombo models.Combos, modelItems []models.ComboItems) error {
	responseItems := []ResponseComboItem{}
	for _, modelItem := range modelItems {
		responseItems = append(responseItems, ResponseComboItem{
			VariationID: modelItem.VariationID,
			Quantity:    modelItem.Quantity,
		})
	}
	return Response(c, statusCode, ResponseCombo{
		ID:             uint64(modelCombo.ID),
		StoreID:        modelCombo.StoreID,
		DiscountAmount: modelCombo.DiscountAmount,
		DiscountType:   utils.DiscountTypeToString(modelCombo.DiscountType),
		ImageUrls:      modelCombo.ImageUrls,
		Description:    modelCombo.Description,
		Title:          modelCombo.Title,
		Items:          responseItems,
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
				VariationID: modelItems[index].VariationID,
				Quantity:    modelItems[index].Quantity,
			})
		}
		responseCombos = append(responseCombos, ResponseCombo{
			ID:             uint64(modelCombo.ID),
			StoreID:        modelCombo.StoreID,
			DiscountAmount: modelCombo.DiscountAmount,
			DiscountType:   utils.DiscountTypeToString(modelCombo.DiscountType),
			ImageUrls:      modelCombo.ImageUrls,
			Description:    modelCombo.Description,
			Title:          modelCombo.Title,
			Items:          responseItems,
		})
	}
	return Response(c, statusCode, responseCombos)
}
