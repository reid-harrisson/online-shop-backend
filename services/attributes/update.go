package prodattrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(attributeID uint64, req *requests.RequestAttribute, modelAttr *models.Attributes) error {
	modelAttr.AttributeName = req.Name
	return service.DB.Save(modelAttr).Error
}
