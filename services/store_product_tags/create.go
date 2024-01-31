package prodTg

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestProductTag, modelTags *[]models.Tags) error {
	service.DB.Where("store_product_id = ?", productID).Delete(models.Tags{})
	for _, tag := range req.Tags {
		modelTag := models.Tags{}
		modelTag.ProductID = productID
		modelTag.Tag = tag
		service.DB.Create(&modelTag)
		*modelTags = append(*modelTags, modelTag)
	}
	return nil
}
