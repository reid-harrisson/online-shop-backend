package prodattrvalsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(attributeID uint64, productID uint64, modelValues *[]models.ProductAttributeValuesWithDetail, req *requests.RequestProductAttributeValue) {
	filterKeys := make(map[string]int)
	for _, modelVar := range *modelValues {
		filterKeys[modelVar.AttributeValue] = 1
	}
	for _, value := range req.Values {
		if filterKeys[value] == 1 {
			filterKeys[value] = 3
		} else {
			filterKeys[value] = 2
		}
	}

	for value, key := range filterKeys {
		if key == 1 {
			service.Delete(value, productID)
		} else if key == 2 {
			service.Create(attributeID, value)
		}
	}

	varRepo := repositories.NewRepositoryProductAttributeValue(service.DB)
	varRepo.ReadByID(modelValues, attributeID)
}

func (service *Service) UpdateByID(attributeValueID uint64, value string) error {
	return service.DB.Model(&models.ProductAttributeValues{}).
		Where("id = ?", attributeValueID).
		Update("attribute_value", value).Error
}
