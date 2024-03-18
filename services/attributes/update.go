package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(attributeID uint64, req *requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	modelAttr.AttributeName = req.Name
	modelAttr.ProductID = attributeID
	service.DB.Save(modelAttr)
}
