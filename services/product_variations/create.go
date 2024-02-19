package prodvarsvc

import "OnlineStoreBackend/models"

func (service *Service) Create(attributeID uint64, variant string, productID uint64) {
	service.DB.Create(&models.ProductVariations{
		AttributeID: attributeID,
		Variant:     variant,
	})
}
