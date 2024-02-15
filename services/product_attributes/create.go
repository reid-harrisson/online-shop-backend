package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	attrsvc "OnlineStoreBackend/services/attributes"
)

func (service *Service) Create(productID uint64, req *requests.RequestAttribute, modelPAttrs *[]models.ProductAttributes) {
	attrService := attrsvc.NewServiceAttribute(service.DB)
	for _, attribute := range req.Attributes {
		modelAttr := models.BaseAttributes{}
		attrService.Create(attribute, &modelAttr)
		modelPAttr := models.ProductAttributes{
			AttributeID: uint64(modelAttr.ID),
			ProductID:   productID,
		}
		service.DB.Where("product_id = ? And attribute_id = ?", productID, modelAttr.ID).First(&modelPAttr)
		service.DB.Save(&modelPAttr)
		*modelPAttrs = append(*modelPAttrs, modelPAttr)
	}
}
