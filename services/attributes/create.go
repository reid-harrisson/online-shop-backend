package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	"strings"
)

func (service *Service) Create(productID uint64, req *requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	modelAttr.AttributeName = req.Name
	modelAttr.Unit = req.Unit
	modelAttr.ProductID = productID
	service.DB.Create(modelAttr)
}

func (service *Service) CreateWithCSV(modelCsv *models.CSVs, productID uint64) {
	valService := prodattrvalsvc.NewServiceProductAttributeValue(service.DB)
	if modelCsv.Attribute1Name != "" {
		modelAttr := models.ProductAttributes{
			AttributeName: modelCsv.Attribute1Name,
			Unit:          "",
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
			Unit:          "",
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
	if modelCsv.Attribute1Name != "" {
		modelAttr := models.ProductAttributes{}
		service.DB.Where("attribute_name = ? And product_id = ?", modelCsv.Attribute1Name, productID).First(&modelAttr)
		if modelAttr.ID == 0 {
			modelAttr.AttributeName = modelCsv.Attribute1Name
			modelAttr.Unit = ""
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
			modelAttr.Unit = ""
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
