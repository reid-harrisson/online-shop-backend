package prodtagsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(tagID uint64, productID uint64) {
	service.DB.Create(models.ProductTags{
		TagID:     tagID,
		ProductID: productID,
	})
}
