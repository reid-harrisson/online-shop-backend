package prodRev

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestProductReview, modelReview *models.ProductReviews) error {
	modelReview.Comment = req.Comment
	modelReview.ProductID = productID
	modelReview.CustomerID = req.CustomerID
	modelReview.Status = "archieved"
	service.DB.Create(modelReview)
	return nil
}
