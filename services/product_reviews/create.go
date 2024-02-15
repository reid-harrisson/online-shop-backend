package revsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelReview *models.ProductReviews, req *requests.RequestProductReview) error {
	modelReview.Comment = req.Comment
	modelReview.ProductID = req.ProductID
	modelReview.CustomerID = req.CustomerID
	modelReview.Status = "archieved"
	service.DB.Create(modelReview)
	return nil
}
