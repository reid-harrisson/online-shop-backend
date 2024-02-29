package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	prodattrvalsvc "OnlineStoreBackend/services/product_attribute_values"
	"strings"
)

func (service *Service) Create(productID uint64, req *requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	modelAttr.AttributeName = req.Name
	modelAttr.Unit = req.Unit
	modelAttr.ProductID = productID
	service.DB.Create(modelAttr)
}

func (service *Service) CreateWithCSV(modelAttrs *[]models.ProductAttributes, modelCsv *models.CSVs, productID uint64) {
	valService := prodattrvalsvc.NewServiceProductAttributeValue(service.DB)
	if modelCsv.Attribute1Name != "" {
		modelAttr := models.ProductAttributes{
			AttributeName: modelCsv.Attribute1Name,
			Unit:          "",
			ProductID:     productID,
		}
		service.DB.Create(&modelAttr)
		*modelAttrs = append(*modelAttrs, modelAttr)
		values := strings.Split(modelCsv.Attribute1Values, ",")
		for _, value := range values {
			valService.Create(uint64(modelAttr.ID), strings.TrimSpace(value))
		}
	}
	if modelCsv.Attribute2Name != "" {
		modelAttr := models.ProductAttributes{
			AttributeName: modelCsv.Attribute2Name,
			Unit:          "",
			ProductID:     productID,
		}
		service.DB.Create(&modelAttr)
		*modelAttrs = append(*modelAttrs, modelAttr)
		values := strings.Split(modelCsv.Attribute1Values, ",")
		for _, value := range values {
			valService.Create(uint64(modelAttr.ID), strings.TrimSpace(value))
		}
	}
}
