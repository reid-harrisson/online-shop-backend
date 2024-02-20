package revsvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateStatus(id uint64, modelReview *models.ProductReviews, reviewStatus string) error {
	if err := service.DB.First(&modelReview, id).Error; err != nil {
		return err
	}

	var status models.ReviewStatuses

	switch reviewStatus {
	case "Approved":
		status = models.StatusReviewApproved
	case "Blocked":
		status = models.StatusReviewBlocked
	}

	modelReview.Status = status

	return service.DB.Save(&modelReview).Error
}
