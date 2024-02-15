package revsvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateStatus(id uint64, modelReview *models.ProductReviews, status string) error {
	if err := service.DB.First(&modelReview, id).Error; err != nil {
		return err
	}
	modelReview.Status = status
	service.DB.Save(&modelReview)
	return nil
}
