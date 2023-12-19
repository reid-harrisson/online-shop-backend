package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	product "OnlineStoreBackend/services/products"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersProduct struct {
	server *s.Server
}

func NewHandlersProduct(server *s.Server) *HandlersProduct {
	return &HandlersProduct{server: server}
}

// Refresh godoc
// @Summary Create Product
// @Description Perform create product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestProduct true "Product Info"
// @Success 201 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product [post]
func (h *HandlersProduct) Create(c echo.Context) error {
	requestProduct := new(requests.RequestProduct)
	if err := c.Bind(requestProduct); err != nil {
		return err
	} else if err := requestProduct.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Please specify a name for the role.")
	}

	modelProduct := models.Products{}
	serviceProduct := product.NewServiceProduct(h.server.DB)
	if err := serviceProduct.Create(&modelProduct, requestProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to create product in online store.")
	}

	// //Audit
	// auditData := utils.AuditData{
	// 	Description: "User created a online store product for " + modelProduct.Name,
	// 	Code:        33002,
	// 	Category:    "Online Store",
	// 	Type:        "Audit",
	// 	Action:      "Create Online Store Product",
	// 	Event:       "store_product->create()",
	// }
	// utils.HelperAudit(h.server.Config.Audit.Url, c, auditData)

	modelProductDetail := models.ProductDetails{}
	repositoryProduct := repositories.NewRepositoryProduct(serviceProduct.DB)
	repositoryProduct.ReadOne(&modelProductDetail, uint64(modelProduct.ID))
	return responses.NewResponseProduct(c, http.StatusCreated, modelProductDetail)
}

// Refresh godoc
// @Summary Read Product
// @Description Perform read product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 209 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/{id} [get]
func (h *HandlersProduct) ReadOne(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	modelProductDetail := models.ProductDetails{}
	repositoryProduct := repositories.NewRepositoryProduct(h.server.DB)
	if err := repositoryProduct.ReadOne(&modelProductDetail, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No product founded.")
	}
	return responses.NewResponseProduct(c, http.StatusCreated, modelProductDetail)
}

// Refresh godoc
// @Summary Read All Products
// @Description Perform read all products
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product [get]
func (h *HandlersProduct) ReadAll(c echo.Context) error {
	modelProductDetails := make([]models.ProductDetails, 0)
	repositoryProduct := repositories.NewRepositoryProduct(h.server.DB)
	repositoryProduct.ReadAll(&modelProductDetails)
	return responses.NewResponseProducts(c, http.StatusCreated, modelProductDetails)
}

// Refresh godoc
// @Summary Read Active Products
// @Description Perform read active products
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/active [get]
func (h *HandlersProduct) ReadActive(c echo.Context) error {
	modelProductDetails := make([]models.ProductDetails, 0)
	repositoryProduct := repositories.NewRepositoryProduct(h.server.DB)
	repositoryProduct.ReadActive(&modelProductDetails)
	return responses.NewResponseProducts(c, http.StatusCreated, modelProductDetails)
}

// Refresh godoc
// @Summary Read Products by pagination
// @Description Perform read products by pagination
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int true "Page" default(0)
// @Param count query int true "Count" default(100)
// @Param company_id query int false "Company ID"
// @Param user_id query int false "User ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} responses.ResponseProductsPaging
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/paging [get]
func (h *HandlersProduct) ReadPaging(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	page, _ := strconv.ParseUint(c.QueryParam("page"), 10, 64)
	count, _ := strconv.ParseUint(c.QueryParam("page"), 10, 64)
	companyID, _ := strconv.ParseUint(c.QueryParam("page"), 10, 64)
	userID, _ := strconv.ParseUint(c.QueryParam("page"), 10, 64)
	totalCount := uint64(0)
	modelProductDetails := make([]models.ProductDetails, 0)
	repositoryProduct := repositories.NewRepositoryProduct(h.server.DB)
	repositoryProduct.ReadPaging(&modelProductDetails, page, count, companyID, userID, keyword, &totalCount)
	return responses.NewResponseProductsPaging(c, http.StatusCreated, modelProductDetails, totalCount)
}

// Refresh godoc
// @Summary Read Products by searching
// @Description Perform read products by searching
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param franchisee_id query int false "Franchisee ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/search [get]
func (h *HandlersProduct) ReadSearch(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	companyID, _ := strconv.ParseUint(c.QueryParam("franchisee_id"), 10, 64)
	modelProductDetails := make([]models.ProductDetails, 0)
	repositoryProduct := repositories.NewRepositoryProduct(h.server.DB)
	repositoryProduct.ReadSearch(&modelProductDetails, companyID, keyword)
	return responses.NewResponseProducts(c, http.StatusCreated, modelProductDetails)
}

// Refresh godoc
// @Summary Update Product
// @Description Perform update product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProduct true "Product Info"
// @Success 209 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/{id} [put]
func (h *HandlersProduct) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	requestProduct := new(requests.RequestProduct)
	if err := c.Bind(requestProduct); err != nil {
		return err
	} else if err := requestProduct.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Please specify a name for the role.")
	}
	service := product.NewServiceProduct(h.server.DB)
	if err := service.Update(id, requestProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to update product in online store.")
	}

	modelProductDetail := models.ProductDetails{}
	repository := repositories.NewRepositoryProduct(h.server.DB)
	repository.ReadOne(&modelProductDetail, id)
	return responses.NewResponseProduct(c, http.StatusCreated, modelProductDetail)
}

// Refresh godoc
// @Summary Delete Product
// @Description Perform delete product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 209 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/{id} [delete]
func (h *HandlersProduct) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	service := product.NewServiceProduct(h.server.DB)
	if err := service.Delete(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to delete product in online store.")
	}
	return responses.MessageResponse(c, http.StatusOK, "Product successfully deleted.")
}
