package revsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
)

func (service *Service) UpdateStatus(id uint64, modelReview *models.Reviews, reviewStatus string) error {
	if err := service.DB.First(&modelReview, id).Error; err != nil {
		return err
	}

	var status = utils.ReviewStatusFromString(reviewStatus)
	modelReview.Status = status

	return service.DB.Save(&modelReview).Error
}
