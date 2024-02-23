package prodattrvalsvc

import "OnlineStoreBackend/models"

func (service *Service) Create(attributeID uint64, value string, productID uint64) {
	service.DB.Create(&models.ProductAttributeValues{
		AttributeID: attributeID,
		Value:       value,
	})
}
