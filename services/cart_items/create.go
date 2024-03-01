package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelItem *models.CartItems, customerID uint64, modelVariation *models.ProductVariations, quantity float64) error {
	modelItem.CustomerID = customerID
	modelItem.VariationID = uint64(modelVariation.ID)
	modelItem.Quantity += quantity

	if modelItem.Quantity > modelVariation.StockLevel {
		modelItem.Quantity = modelVariation.StockLevel
	}

	return service.DB.Save(modelItem).Error
}
