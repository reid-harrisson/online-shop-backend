package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) UpdateQuantity(modelItem *models.CartItems, modelVar models.ProductVariations, quantity float64) {
	modelItem.Quantity = quantity
	if modelItem.Quantity > modelVar.StockLevel {
		modelItem.Quantity = modelVar.StockLevel
	}
	service.DB.Save(modelItem)
}
