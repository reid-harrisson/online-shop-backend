package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) UpdateQuantity(cartID uint64, modelItem *models.CartItems, modelProduct models.Products, quantity float64) {
	modelItem.Quantity = quantity
	if modelItem.Quantity > modelProduct.StockQuantity {
		modelItem.Quantity = modelProduct.StockQuantity
	}
	service.DB.Save(modelItem)
}
