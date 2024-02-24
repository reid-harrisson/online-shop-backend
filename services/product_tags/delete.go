package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
)

func (service *Service) Delete(tag string) {
	modelTag := models.StoreTags{}
	tagRepo := repositories.NewRepositoryTag(service.DB)
	tagRepo.ReadByName(&modelTag, tag)
	if modelTag.ID != 0 {
		service.DB.Where("tag_id = ?", modelTag.ID).Delete(models.ProductTags{})
	}
}
