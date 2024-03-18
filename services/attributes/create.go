package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	"fmt"
	"strings"
)

func (service *Service) Create(productID uint64, req *requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	modelAttr.AttributeName = req.Name
	modelAttr.ProductID = productID
	service.DB.Create(modelAttr)
}

func (service *Service) CreateWithCSV(modelNewAttrs *[]models.ProductAttributes, attrMatches []string, attrIndices map[string]int) {
	modelCurAttrs := []models.ProductAttributes{}
	service.DB.Where("Concat(product_id, ':', attribute_name) In (?)", attrMatches).Find(&modelCurAttrs)
	for _, modelAttr := range modelCurAttrs {
		match := fmt.Sprintf("%d:%s", modelAttr.ProductID, modelAttr.AttributeName)
		index := attrIndices[match]
		(*modelNewAttrs)[index].ID = modelAttr.ID
	}
	service.DB.Save(modelNewAttrs)
}

func (service *Service) CreateWithCSV1(modelCsv *models.CSVs, productID uint64) {
	valService := prodattrvalsvc.NewServiceProductAttributeValue(service.DB)
	if modelCsv.AttributeName != "" {
		modelAttr := models.ProductAttributes{
			AttributeName: modelCsv.AttributeName,
			ProductID:     productID,
		}
		service.DB.Create(&modelAttr)
		values := strings.Split(modelCsv.Attribute1Values, ",")
		for _, value := range values {
			if value != "" {
				valService.Create(uint64(modelAttr.ID), strings.TrimSpace(value))
			}
		}
	}
	if modelCsv.Attribute2Name != "" {
		modelAttr := models.ProductAttributes{
			AttributeName: modelCsv.Attribute2Name,
			ProductID:     productID,
		}
		service.DB.Create(&modelAttr)
		values := strings.Split(modelCsv.Attribute1Values, ",")
		for _, value := range values {
			if value != "" {
				valService.Create(uint64(modelAttr.ID), strings.TrimSpace(value))
			}
		}
	}
}

func (service *Service) UpdateWithCSV(modelVals *[]models.ProductAttributeValues, modelCsv *models.CSVs, productID uint64) {
	valService := prodattrvalsvc.NewServiceProductAttributeValue(service.DB)
	if modelCsv.AttributeName != "" {
		modelAttr := models.ProductAttributes{}
		service.DB.Where("attribute_name = ? And product_id = ?", modelCsv.AttributeName, productID).First(&modelAttr)
		if modelAttr.ID == 0 {
			modelAttr.AttributeName = modelCsv.AttributeName
			modelAttr.ProductID = productID
			service.DB.Create(&modelAttr)
		}
		value := strings.TrimSpace(modelCsv.Attribute1Values)
		if value != "" {
			modelVal := models.ProductAttributeValues{}
			valService.DB.Where("attribute_id = ? And attribute_value = ?", modelAttr.ID, value).First(&modelVal)
			if modelVal.ID == 0 {
				valService.CreateWithModel(&modelVal, uint64(modelAttr.ID), value)
			}
			*modelVals = append(*modelVals, modelVal)
		}
	}
	if modelCsv.Attribute2Name != "" {
		modelAttr := models.ProductAttributes{}
		service.DB.Where("attribute_name = ?", modelCsv.Attribute2Name).First(&modelAttr)
		if modelAttr.ID == 0 {
			modelAttr.AttributeName = modelCsv.Attribute2Name
			modelAttr.ProductID = productID
			service.DB.Create(&modelAttr)
		}
		value := strings.TrimSpace(modelCsv.Attribute2Values)
		if value != "" {
			modelVal := models.ProductAttributeValues{}
			valService.DB.Where("attribute_id = ? And attribute_value = ?", modelAttr.ID, value).First(&modelVal)
			if modelVal.ID == 0 {
				valService.CreateWithModel(&modelVal, uint64(modelAttr.ID), value)
			}
			*modelVals = append(*modelVals, modelVal)
		}
	}
}
