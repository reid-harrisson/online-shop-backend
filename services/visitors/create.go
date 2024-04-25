package vistsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelVisitor *models.Visitors, req *requests.RequestVisitor) error {
	modelVisitor.StoreID = req.StoreID
	modelVisitor.ProductID = req.ProductID
	modelVisitor.IpAddress = req.IpAddress
	modelVisitor.Page = utils.PageTypeFromString(req.Page)
	modelVisitor.Bounce = req.Bounce
	modelVisitor.LoadingTime = req.LoadingTime
	return service.DB.Create(modelVisitor).Error
}
