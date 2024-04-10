package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Create(productID uint64, req *requests.RequestAttribute, modelAttr *models.Attributes) {
	modelAttr.AttributeName = req.Name
	modelAttr.ProductID = productID
	service.DB.Create(modelAttr)
}

func (service *Service) CreateWithCSV(modelNewAttrs *[]models.Attributes, attrMatches []string, attrIndices map[string]int) {
	modelCurAttrs := []models.Attributes{}
	service.DB.Where("Concat(product_id, ':', attribute_name) In (?)", attrMatches).Find(&modelCurAttrs)
	for _, modelAttr := range modelCurAttrs {
		match := fmt.Sprintf("%d:%s", modelAttr.ProductID, modelAttr.AttributeName)
		index := attrIndices[match]
		(*modelNewAttrs)[index].ID = modelAttr.ID
	}
	service.DB.Save(modelNewAttrs)
}
