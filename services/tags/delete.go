package tagsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(tagID uint64) {
	service.DB.Where("id = ?", tagID).Delete(&models.Tags{})
}
