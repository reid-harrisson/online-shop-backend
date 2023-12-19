package cart

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(model *models.CartItems, request *requests.RequestCartItem) error {
	if err := service.DB.Where("user_id = ? And store_product_id = ?", request.UserID, request.ProductID).First(model).Error; err != nil {
		model.UserID = request.UserID
		model.ProductID = request.ProductID
		model.Quantity = request.Quantity
		if err = service.DB.Create(model).Error; err != nil {
			return err
		}
	} else {
		model.Quantity += request.Quantity
	}
	modelProduct := models.Products{}
	if err := service.DB.First(&modelProduct, model.ProductID).Error; err != nil {
		return err
	} else if modelProduct.StockQuantity < model.Quantity {
		model.Quantity = modelProduct.StockQuantity
	}
	return service.DB.Save(model).Error
}
