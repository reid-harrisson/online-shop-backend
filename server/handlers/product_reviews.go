package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	revsvc "OnlineStoreBackend/services/product_reviews"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersProductReviews struct {
	server *s.Server
}

func NewHandlersProductReviews(server *s.Server) *HandlersProductReviews {
	return &HandlersProductReviews{server: server}
}

// Refresh godoc
// @Summary Create product review
// @Tags Product Review
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Param customer_id query int true "Customer ID"
// @Param params body requests.RequestProductReview true "Review Info"
// @Success 201 {object} responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review [post]
func (h *HandlersProductReviews) CreateReview(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	requestProductReview := new(requests.RequestProductReview)

	if err := c.Bind(requestProductReview); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := requestProductReview.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProductReview := models.ProductReviews{}

	serviceProductReview := revsvc.NewServiceProductReview(h.server.DB)

	if err := serviceProductReview.Create(&modelProductReview, requestProductReview, customerID, productID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseReview(c, http.StatusCreated, modelProductReview)
}

// Refresh godoc
// @Summary Moderate product review
// @Tags Product Review
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} []responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/publish/{id} [get]
func (h *HandlersProductReviews) ReadPublishedReviews(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var modelProductPublishedReviews []models.ProductReviews

	repositoryProductReview := repositories.NewRepositoryReview(h.server.DB)
	repositoryProductReview.ReadPublishReviews(&modelProductPublishedReviews, id)

	return responses.NewResponseProductReviews(c, http.StatusOK, modelProductPublishedReviews)
}

// Refresh godoc
// @Summary Read all product reviews
// @Tags Product Review
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} []responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/{id} [get]
func (h *HandlersProductReviews) ReadAll(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var modelProductReviews []models.ProductReviews

	repositoryProductReview := repositories.NewRepositoryReview(h.server.DB)
	repositoryProductReview.ReadReviews(&modelProductReviews, id)

	return responses.NewResponseProductReviews(c, http.StatusOK, modelProductReviews)
}

// Refresh godoc
// @Summary Moderate product review
// @Tags Product Review
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param id path int true "Product Review ID"
// @Param status query string true "Status" default(Approved)
// @Success 200 {object} responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/moderate/{id} [put]
func (h *HandlersProductReviews) ModerateReview(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	status := c.QueryParam("status")

	modelProductReview := models.ProductReviews{}

	serviceProductReview := revsvc.NewServiceProductReview(h.server.DB)

	if err := serviceProductReview.UpdateStatus(id, &modelProductReview, status); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseReview(c, http.StatusOK, modelProductReview)
}
