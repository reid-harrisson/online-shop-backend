package prodattrvalsvc

import "OnlineStoreBackend/models"

func (service *Service) Create(attributeID uint64, value string) error {
	return service.DB.Create(&models.ProductAttributeValues{
		AttributeID:    attributeID,
		AttributeValue: value,
	}).Error
}
