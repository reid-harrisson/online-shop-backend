package vistsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelVisitor *models.Visitors, req *requests.RequestVisitor) {
	modelVisitor.StoreID = req.StoreID
	modelVisitor.ProductID = req.ProductID
	modelVisitor.IpAddress = req.IpAddress
	modelVisitor.Page = utils.PageTypeFromString(req.Page)
	modelVisitor.Bounce = req.Bounce
	modelVisitor.LoadingTime = req.LoadingTime
	modelVisitor.ErrorCode = req.ErrorCode
	service.DB.Create(modelVisitor)
}
