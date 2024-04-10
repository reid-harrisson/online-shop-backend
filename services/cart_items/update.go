package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) UpdateQuantity(modelItem *models.CartItems, modelVar models.Variations, quantity float64) error {
	modelItem.Quantity = quantity
	if modelItem.Quantity > modelVar.StockLevel {
		modelItem.Quantity = modelVar.StockLevel
	}
	return service.DB.Where("id = ?", modelItem.ID).Update("quantity", modelItem.Quantity).Error
}
