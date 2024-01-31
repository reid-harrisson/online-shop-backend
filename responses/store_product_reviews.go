package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductReview struct {
	ID         uint64 `json:"id"`
	Comment    string `json:"comment"`
	CustomerID uint64 `json:"customer_id"`
	ProductID  uint64 `json:"store_product_id"`
	Status     string `json:"status"`
}

func NewResponseReview(c echo.Context, statusCode int, modelReview models.ProductReviews) error {
	responseReview := ResponseProductReview{
		ID:         uint64(modelReview.ID),
		Comment:    modelReview.Comment,
		CustomerID: modelReview.CustomerID,
		ProductID:  modelReview.ProductID,
		Status:     modelReview.Status,
	}
	return Response(c, statusCode, responseReview)
}

func NewResponseProductReviews(c echo.Context, statusCode int, modelReviews []models.ProductReviews) error {
	responseReviews := make([]ResponseProductReview, 0)
	for _, modelReview := range modelReviews {
		responseReviews = append(responseReviews, ResponseProductReview{
			ID:         uint64(modelReview.ID),
			Comment:    modelReview.Comment,
			CustomerID: modelReview.CustomerID,
			ProductID:  modelReview.ProductID,
			Status:     modelReview.Status,
		})
	}
	return Response(c, statusCode, responseReviews)
}
