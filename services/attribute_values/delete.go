package prodattrvalsvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(value string, productID uint64) {
	service.DB.Where("value = ? And product_id = ?", value, productID).Delete(models.ProductAttributeValues{})
}

func (service *Service) DeleteByID(attributeValueID uint64) error {
	return service.DB.Where("id = ?", attributeValueID).Delete(&models.ProductAttributeValues{}).Error
}
