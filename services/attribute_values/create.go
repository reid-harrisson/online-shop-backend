package prodattrvalsvc

import "OnlineStoreBackend/models"

func (service *Service) Create(attributeID uint64, value string) error {
	return service.DB.Create(&models.ProductAttributeValues{
		AttributeID:    attributeID,
		AttributeValue: value,
	}).Error
}

func (service *Service) CreateWithModel(modelVal *models.ProductAttributeValues, attributeID uint64, value string) error {
	modelVal.AttributeID = attributeID
	modelVal.AttributeValue = value
	return service.DB.Create(&modelVal).Error
}
