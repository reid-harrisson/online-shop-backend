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

func (service *Service) CreateWithCSV(modelNewAttrs *[]models.Attributes, attrMatches []string, attrIndices map[string]int) error {
	modelCurAttrs := []models.Attributes{}
	if err := service.DB.Where("Concat(product_id, ':', attribute_name) In (?)", attrMatches).Find(&modelCurAttrs).Error; err != nil {
		return err
	}
	for _, modelAttr := range modelCurAttrs {
		match := fmt.Sprintf("%d:%s", modelAttr.ProductID, modelAttr.AttributeName)
		index := attrIndices[match]
		(*modelNewAttrs)[index].ID = modelAttr.ID
	}
	return service.DB.Save(modelNewAttrs).Error
}
