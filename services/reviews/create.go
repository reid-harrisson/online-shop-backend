package revsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelProductReview *models.ProductReviews, requestProductReview *requests.RequestProductReview, customerID uint64, productID uint64) error {
	modelProductReview.CustomerID = customerID
	modelProductReview.ProductID = productID

	modelProductReview.Comment = requestProductReview.Comment
	modelProductReview.Rate = requestProductReview.Rate
	modelProductReview.Status = utils.StatusReviewPending

	return service.DB.Create(modelProductReview).Error
}
