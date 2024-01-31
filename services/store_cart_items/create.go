package cart

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelCart *models.CartItems, req *requests.RequestCartItem) error {
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	if err := prodRepo.Read(&modelProduct, req.ProductID); err != nil {
		return err
	}
	if err := service.DB.Where("customer_id = ? And store_product_id = ?", req.CustomerID, req.ProductID).First(modelCart).Error; err != nil {
		modelCart.CustomerID = req.CustomerID
		modelCart.ProductID = req.ProductID
		modelCart.Quantity = req.Quantity
		modelCart.StoreID = modelProduct.StoreID
		if modelCart.Quantity > modelProduct.StockQuantity {
			modelCart.Quantity = modelProduct.StockQuantity
		}
		return service.DB.Create(modelCart).Error
	}
	modelCart.Quantity += req.Quantity
	if modelCart.Quantity > modelProduct.StockQuantity {
		modelCart.Quantity = modelProduct.StockQuantity
	}
	return service.DB.Save(modelCart).Error
}
