package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
)

func (service *Service) Delete(tag string) error {
	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(service.DB)
	err := tagRepo.ReadByName(&modelTag, tag, modelTag.StoreID)
	if err != nil {
		return err
	}
	if modelTag.ID != 0 {
		err = service.DB.Where("tag_id = ?", modelTag.ID).Delete(&models.ProductTags{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
