package revsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelReview *models.ProductReviews, requestProductReview *requests.RequestProductReview, customerID uint64, productID uint64) error {
	modelReview.CustomerID = customerID
	modelReview.ProductID = productID
	modelReview.Comment = requestProductReview.Comment
	modelReview.Rate = requestProductReview.Rate
	modelReview.Status = 0
	service.DB.Create(modelReview)
	return nil
}
