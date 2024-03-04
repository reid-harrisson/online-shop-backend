package linksvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
)

func (service *Service) Create(productID uint64, linkID uint64, isUpCross utils.SellTypes) error {
	return service.DB.Where("link_id = ? And product_id = ?", linkID, productID).
		FirstOrCreate(&models.ProductLinks{
			ProductID: productID,
			LinkID:    linkID,
			IsUpCross: isUpCross,
		}).Error
}
