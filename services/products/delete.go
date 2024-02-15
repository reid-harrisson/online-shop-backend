package prodsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(productID uint64) error {
	query := service.DB.Where("product_id = ?", productID)
	query.Delete(models.ProductChannels{})
	query.Delete(models.ProductContents{})
	query.Delete(models.BaseAttributes{})
	query.Delete(models.BaseTags{})
	query.Delete(models.ProductReviews{})
	return service.DB.Delete(models.Products{}, productID).Error
}
