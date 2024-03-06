package cartsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
)

func (service *Service) Create(modelItem *models.CartItems, customerID uint64, modelVar *models.ProductVariations, quantity float64) error {
	modelItem.CustomerID = customerID
	modelItem.VariationID = uint64(modelVar.ID)
	modelItem.Quantity += quantity

	if modelItem.Quantity > modelVar.StockLevel {
		if modelVar.BackOrderStatus == utils.Disabled {
			modelItem.Quantity = modelVar.StockLevel
		}
	}

	return service.DB.Save(modelItem).Error
}
