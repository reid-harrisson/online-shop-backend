package cartsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelCart *models.CartItems, customerID uint64, modelProduct models.Products, quantity float64) error {
	productID := uint64(modelProduct.ID)

	if err := service.DB.Where("customer_id = ? And product_id = ?", customerID, productID).First(modelCart).Error; err != nil {
		modelCart.CustomerID = customerID
		modelCart.ProductID = productID
		modelCart.Quantity = quantity
		if modelCart.Quantity > modelProduct.StockQuantity {
			modelCart.Quantity = modelProduct.StockQuantity
		}
		return service.DB.Create(modelCart).Error
	}
	modelCart.Quantity += quantity
	if modelCart.Quantity > modelProduct.StockQuantity {
		modelCart.Quantity = modelProduct.StockQuantity
	}
	return service.DB.Save(modelCart).Error
}
