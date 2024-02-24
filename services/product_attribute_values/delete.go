package prodattrvalsvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(value string, productID uint64) {
	service.DB.Where("value = ? And product_id = ?", value, productID).Delete(models.ProductAttributeValues{})
}
