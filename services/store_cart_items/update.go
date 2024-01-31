package cart

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) UpdateQuantity(cartID uint64, modelCart *models.CartItems, req *requests.RequestProductQuantity) error {
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	if err := service.DB.First(modelCart, cartID).Error; err != nil {
		return err
	}
	if err := prodRepo.Read(&modelProduct, modelCart.ProductID); err != nil {
		return err
	}
	modelCart.Quantity = req.Quantity
	if modelCart.Quantity > modelProduct.StockQuantity {
		modelCart.Quantity = modelProduct.StockQuantity
	}
	return service.DB.Save(modelCart).Error
}
