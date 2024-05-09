package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
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
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/review [post]
func (h *HandlersReviews) CreateReview(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	requestReview := new(requests.RequestReview)

	err = c.Bind(requestReview)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	err = requestReview.Validate()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelReview := models.Reviews{}

	serviceReview := revsvc.NewServiceReview(h.server.DB)

	err = serviceReview.Create(&modelReview, requestReview, customerID, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/review/publish/{id} [get]
func (h *HandlersReviews) ReadPublishedReviews(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var modelPublishedReviews []models.Reviews

	repositoryReview := repositories.NewRepositoryReview(h.server.DB)
	err = repositoryReview.ReadPublishReviews(&modelPublishedReviews, id)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseReviews(c, http.StatusOK, modelPublishedReviews)
}

// Refresh godoc
// @Summary Read all product reviews
// @Tags Review Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} []responses.ResponseReview
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/review/{id} [get]
func (h *HandlersReviews) ReadAll(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var modelReviews []models.Reviews

	repositoryReview := repositories.NewRepositoryReview(h.server.DB)
	err = repositoryReview.ReadReviews(&modelReviews, id)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseReviews(c, http.StatusOK, modelReviews)
}

// Refresh godoc
// @Summary Moderate product review
// @Tags Review Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product Review ID"
// @Param status query string true "Status" default(Approved)
// @Success 200 {object} responses.ResponseReview
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/review/moderate/{id} [put]
func (h *HandlersReviews) ModerateReview(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	status := c.QueryParam("status")

	modelReview := models.Reviews{}

	serviceReview := revsvc.NewServiceReview(h.server.DB)

	err = serviceReview.UpdateStatus(id, &modelReview, status)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseReview(c, http.StatusOK, modelReview)
}
