package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodRate "OnlineStoreBackend/services/store_product_rates"
	prodRev "OnlineStoreBackend/services/store_product_reviews"
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
// @Summary Add rate
// @Tags product reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestProductRate true "Product Rate Info"
// @Success 201 {object} responses.ResponseProductCustomerRate
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/rate [post]
func (h *HandlersProductReviews) CreateRate(c echo.Context) error {
	req := new(requests.RequestProductRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelRate := models.ProductCustomerRates{}
	rateService := prodRate.CreateService(h.server.DB)
	if err := rateService.Create(&modelRate, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductCustomerRate(c, http.StatusCreated, modelRate)
}

// Refresh godoc
// @Summary View rate
// @Tags product reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} responses.ResponseProductRate
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/rate [get]
func (h *HandlersProductReviews) ReadRate(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	revRepo := repositories.NewRepositoryReview(h.server.DB)
	modelRate := models.ProductRates{}
	if err := revRepo.ReadRate(&modelRate, productID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductRate(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary View reviews
// @Tags product reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} []responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review [get]
func (h *HandlersProductReviews) ReadReview(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	revRepo := repositories.NewRepositoryReview(h.server.DB)
	modelReviews := make([]models.ProductReviews, 0)
	if err := revRepo.ReadReviews(&modelReviews, productID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductReviews(c, http.StatusOK, modelReviews)
}

// Refresh godoc
// @Summary View published reviews
// @Tags product reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} []responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/publish [get]
func (h *HandlersProductReviews) ReadPublishReview(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	revRepo := repositories.NewRepositoryReview(h.server.DB)
	modelReviews := make([]models.ProductReviews, 0)
	if err := revRepo.ReadPublishReviews(&modelReviews, productID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductReviews(c, http.StatusOK, modelReviews)
}

// Refresh godoc
// @Summary Publish review
// @Tags product reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Review ID"
// @Param params body requests.RequestProductReviewStatus true "Status"
// @Success 200 {object} responses.ResponseProductReview
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/publish/{id} [put]
func (h *HandlersProductReviews) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductReviewStatus)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelReview := models.ProductReviews{}
	reviewService := prodRev.CreateService(h.server.DB)
	if err := reviewService.UpdateStatus(id, &modelReview, "published"); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseReview(c, http.StatusOK, modelReview)
}

// Refresh godoc
// @Summary Delete review
// @Tags product reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Review ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /api/v1/review/{id} [delete]
func (h *HandlersProductReviews) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reviewService := prodRev.CreateService(h.server.DB)
	if err := reviewService.Delete(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusOK, "Review successfully deleted.")
}
