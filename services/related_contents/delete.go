package contsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(ContentID uint64) {
	service.DB.Where("content_id = ?", ContentID).Delete(models.ProductContents{})
}
