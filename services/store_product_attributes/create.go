package prodAttr

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestProductAttribute, modelAttrs *[]models.Attributes) error {
	service.DB.Where("store_product_id = ?", productID).Delete(models.Attributes{})
	for attribute, value := range req.Attributes {
		modelAttr := models.Attributes{}
		modelAttr.ProductID = productID
		modelAttr.Attribute = attribute
		modelAttr.Value = value
		service.DB.Create(&modelAttr)
		*modelAttrs = append(*modelAttrs, modelAttr)
	}
	return nil
}
