package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(attributeID uint64, req *requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	modelAttr.Name = req.Name
	modelAttr.Unit = req.Unit
	modelAttr.ProductID = attributeID
	service.DB.Save(modelAttr)
}
