package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(req requests.RequestAttribute, modelAttr *models.ProductAttributes) {
	service.DB.Where("name = ?", req.Name).First(modelAttr)
	if modelAttr.ID == 0 {
		modelAttr.Name = req.Name
		modelAttr.Unit = req.Unit
		service.DB.Create(modelAttr)
	}
}
