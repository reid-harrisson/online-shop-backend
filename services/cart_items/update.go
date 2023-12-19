package cart

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(model *models.CartItems, id uint64, request *requests.RequestCartItem) error {
	if err := service.DB.First(model, id).Error; err != nil {
		return err
	}
	model.Quantity = request.Quantity
	modelProduct := models.Products{}
	if err := service.DB.First(&modelProduct, model.ProductID).Error; err != nil {
		return err
	} else if modelProduct.StockQuantity < model.Quantity {
		model.Quantity = modelProduct.StockQuantity
	}
	return service.DB.Save(model).Error
}
