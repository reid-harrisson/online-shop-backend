package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	revsvc "OnlineStoreBackend/services/reviews"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersReviews struct {
	server *s.Server
}

func NewHandlersReviews(server *s.Server) *HandlersReviews {
	return &HandlersReviews{server: server}
}

// Refresh godoc
// @Summary Create product review
// @Tags Review Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Param params body requests.RequestReview true "Review Info"
// @Success 201 {object} responses.ResponseReview
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/review [post]
func (h *HandlersReviews) CreateReview(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	requestReview := new(requests.RequestReview)

	if err := c.Bind(requestReview); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := requestReview.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelReview := models.Reviews{}

	serviceReview := revsvc.NewServiceReview(h.server.DB)

	if err := serviceReview.Create(&modelReview, requestReview, customerID, productID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseReview(c, http.StatusCreated, modelReview)
}

// Refresh godoc
// @Summary Moderate product review
// @Tags Review Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} []responses.ResponseReview
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/review/publish/{id} [get]
func (h *HandlersReviews) ReadPublishedReviews(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var modelPublishedReviews []models.Reviews

	repositoryReview := repositories.NewRepositoryReview(h.server.DB)
	repositoryReview.ReadPublishReviews(&modelPublishedReviews, id)

	return responses.NewResponseReviews(c, http.StatusOK, modelPublishedReviews)
}

// Refresh godoc
// @Summary Read all product reviews
// @Tags Review Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} []responses.ResponseReview
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/review/{id} [get]
func (h *HandlersReviews) ReadAll(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var modelReviews []models.Reviews

	repositoryReview := repositories.NewRepositoryReview(h.server.DB)
	repositoryReview.ReadReviews(&modelReviews, id)

	return responses.NewResponseReviews(c, http.StatusOK, modelReviews)
}

// Refresh godoc
// @Summary Moderate product review
// @Tags Review Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param id path int true "Product Review ID"
// @Param status query string true "Status" default(Approved)
// @Success 200 {object} responses.ResponseReview
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/review/moderate/{id} [put]
func (h *HandlersReviews) ModerateReview(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	status := c.QueryParam("status")

	modelReview := models.Reviews{}

	serviceReview := revsvc.NewServiceReview(h.server.DB)

	if err := serviceReview.UpdateStatus(id, &modelReview, status); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseReview(c, http.StatusOK, modelReview)
}
