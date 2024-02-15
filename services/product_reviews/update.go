package revsvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateStatus(id uint64, modelReview *models.ProductReviews, reviewStatus string) error {
	if err := service.DB.First(&modelReview, id).Error; err != nil {
		return err
	}

	var status int8

	switch reviewStatus {
	case "Approved":
		status = 1
	case "Blocked":
		status = 2
	}

	modelReview.Status = status

	return service.DB.Save(&modelReview).Error
}
