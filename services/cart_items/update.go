package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) UpdateQuantity(cartID uint64, modelItem *models.CartItems, quantity float64) error {
	var modelVariation = models.Variations{}

	if err := service.DB.Model(models.CartItems{}).First(modelItem, cartID).Error; err != nil {
		return err
	}

	if err := service.DB.Model(models.Variations{}).First(modelVariation, modelItem.VariationID).Error; err != nil {
		return err
	}

	modelItem.Quantity = quantity
	if modelItem.Quantity > modelVariation.StockLevel {
		modelItem.Quantity = modelVariation.StockLevel
	}

	return service.DB.
		Where("id = ?", modelItem.ID).
		Update("quantity", modelItem.Quantity).
		Error
}
