package revsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelReview *models.Reviews, requestReview *requests.RequestReview, customerID uint64, productID uint64) error {
	modelReview.CustomerID = customerID
	modelReview.ProductID = productID

	modelReview.Comment = requestReview.Comment
	modelReview.Rate = requestReview.Rate
	modelReview.Status = utils.StatusReviewPending

	return service.DB.Create(modelReview).Error
}
