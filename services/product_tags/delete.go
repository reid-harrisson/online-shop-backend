package prodtagsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(tagID uint64) {
	service.DB.Where("tag_id = ?", tagID).Delete(models.ProductTags{})
}
