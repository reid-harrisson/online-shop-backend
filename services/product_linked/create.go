package linkedsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
)

func (service *Service) Create(productID uint64, linkID uint64, isUpCross utils.SellTypes) error {
	modelLink := models.ProductLinked{}
	service.DB.Where("link_id = ? And product_id = ?", linkID, productID).First(&modelLink)
	modelLink.ProductID = productID
	modelLink.LinkID = linkID
	modelLink.IsUpCross = isUpCross
	return service.DB.Save(&modelLink).Error
}
