package cart

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(cartID uint64) error {
	return service.DB.Where("id = ?", cartID).Delete(models.CartItems{}).Error
}

func (Service *Service) DeleteAll(userID uint64) error {
	return Service.DB.Where("user_id = ?", userID).Delete(models.CartItems{}).Error
}

func (Service *Service) DeleteBuy(userID uint64) error {
	modelItems := make([]models.CartItems, 0)
	if err := Service.DB.Where("user_id = ?", userID).Find(&modelItems).Error; err != nil {
		return err
	}
	for _, model := range modelItems {
		modelProduct := models.Products{}
		if err := Service.DB.First(&modelProduct, model.ProductID).Error; err != nil {
			return err
		}
		modelProduct.StockQuantity -= model.Quantity
		if err := Service.DB.Save(&modelProduct).Error; err != nil {
			return err
		}
	}
	return Service.DB.Delete(&modelItems).Error
}
