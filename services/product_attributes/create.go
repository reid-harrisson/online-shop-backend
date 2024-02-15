package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	modelAttr.Name = req.Name
	modelAttr.Unit = req.Unit
	modelAttr.ProductID = productID
	service.DB.Create(modelAttr)
}
