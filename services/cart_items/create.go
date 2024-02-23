package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelCart *models.CartItems, customerID uint64, modelVariation models.ProductVariations, quantity float64) error {
	productID := uint64(modelVariation.ID)

	if err := service.DB.Where("customer_id = ? And product_id = ?", customerID, productID).First(modelCart).Error; err != nil {
		modelCart.CustomerID = customerID
		modelCart.VariationID = productID
		modelCart.Quantity = quantity

		if modelCart.Quantity > modelVariation.StockLevel {
			modelCart.Quantity = modelVariation.StockLevel
		}

		return service.DB.Create(modelCart).Error
	}

	modelCart.Quantity += quantity

	if modelCart.Quantity > modelVariation.StockLevel {
		modelCart.Quantity = modelVariation.StockLevel
	}

	return service.DB.Save(modelCart).Error
}
