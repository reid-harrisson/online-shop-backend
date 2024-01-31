package prod

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(productID uint64) error {
	query := service.DB.Where("store_product_id = ?", productID)
	query.Delete(models.ProductChannels{})
	query.Delete(models.ProductContents{})
	query.Delete(models.Attributes{})
	query.Delete(models.Tags{})
	query.Delete(models.ProductReviews{})
	return service.DB.Delete(models.Products{}, productID).Error
}
