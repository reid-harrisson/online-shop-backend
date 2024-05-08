package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	linksvc "OnlineStoreBackend/services/links"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	prodsvc "OnlineStoreBackend/services/products"
	chansvc "OnlineStoreBackend/services/related_channels"
	contsvc "OnlineStoreBackend/services/related_contents"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HandlersProducts struct {
	server *s.Server
}

func NewHandlersProducts(server *s.Server) *HandlersProducts {
	return &HandlersProducts{server: server}
}

func ChangeToDraft(db *gorm.DB, modelProduct *models.Products) error {
	prodService := prodsvc.NewServiceProduct(db)
	if err := prodService.UpdateStatus(uint64(modelProduct.ID), utils.Draft); err != nil {
		return err
	}
	return nil
}

func CheckProduct(db *gorm.DB, modelProduct *models.Products, productID uint64) string {
	prodRepo := repositories.NewRepositoryProduct(db)
	err := prodRepo.ReadByID(modelProduct, productID)

	if err == gorm.ErrRecordNotFound {
		return constants.NotFound
	}
	if modelProduct.Status == utils.Pending {
		return constants.ProductOnPendingStatus
	}
	return ""
}

// Refresh godoc
// @Summary Add product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestProductWithDetail true "Product Info"
// @Success 201 {object} responses.ResponseProductWithDetail
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product [post]
func (h *HandlersProducts) Create(c echo.Context) error {
	req := new(requests.RequestProductWithDetail)

	err := c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	err = req.Validate()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.ProductsWithDetail{}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.Create(&modelProduct.Products, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadDetail(&modelProduct, uint64(modelProduct.ID))
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	err = prodService.UpdateStatus(uint64(modelProduct.ID), utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProductWithDetail(c, http.StatusCreated, modelProduct)
}

// Refresh godoc
// @Summary Read product by ID
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProductWithDetail
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/{id} [get]
func (h *HandlersProducts) ReadByID(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.ProductsWithDetail{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadDetail(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProductWithDetail(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Read all products
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product [get]
func (h *HandlersProducts) ReadAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProducts := make([]models.Products, 0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadAll(&modelProducts, storeID, keyword)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Read approved products
// @Tags Product Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param page query int true "Page"
// @Param count query int true "Count"
// @Success 200 {object} responses.ResponseProductApprovedPaging
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/approved  [get]
func (h *HandlersProducts) ReadApproved(c echo.Context) error {
	// customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	// }

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	count, err := strconv.Atoi(c.QueryParam("count"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	exchangeRate, currencyCode := 1.0, "$"
	modelProducts := make([]models.ProductsApproved, 0)

	totalCount := int64(0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadApproved(&modelProducts, storeID, page, count, &totalCount)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// taxRepo := repositories.NewRepositoryTax(h.server.DB)
	// err = taxRepo.ReadCurrency(&currencyCode, &exchangeRate, customerID)
	// if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != http.StatusNotFound {
	// 	return responses.ErrorResponse(c, statusCode, message)
	// }

	return responses.NewResponseProductsApprovedPaging(c, http.StatusOK, modelProducts, exchangeRate, currencyCode, totalCount)
}

// Refresh godoc
// @Summary Read products by category
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param category_id query int rue "Category ID"
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/category [get]
func (h *HandlersProducts) ReadByCategory(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	cateID, err := strconv.ParseUint(c.QueryParam("category_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProducts := make([]models.Products, 0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByCategory(&modelProducts, storeID, cateID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Read products by tags and keyword
// @Tags Product Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param store_id query int false "Store ID"
// @Param tags query string false "Tags"
// @Param keyword query string false "Keyword"
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/search [get]
func (h *HandlersProducts) ReadSearch(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	keyword := c.QueryParam("keyword")
	tags := strings.Split(c.QueryParam("tags"), ",")

	modelProducts := make([]models.Products, 0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByTags(&modelProducts, storeID, tags, keyword)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Read products by pagination
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int true "Page" default(0)
// @Param count query int true "Count" default(100)
// @Param store_id query int false "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} responses.ResponseProductsPaging
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/paging [get]
func (h *HandlersProducts) ReadPaging(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	count, err := strconv.ParseInt(c.QueryParam("count"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	totalCount := int64(0)
	modelProducts := make([]models.Products, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)

	err = prodRepo.ReadPaging(&modelProducts, int(page), int(count), storeID, keyword, &totalCount)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProductsPaging(c, http.StatusOK, modelProducts, totalCount)
}

// Refresh godoc
// @Summary Edit product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProduct true "Product Info"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/{id} [put]
func (h *HandlersProducts) Update(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Get request
	req := new(requests.RequestProduct)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Check product status
	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	// Update product
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.Update(&modelProduct, req)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	// Change status of product
	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}
	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Approve product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/approve/{id} [put]
func (h *HandlersProducts) Approve(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	// Approve status
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(uint64(modelProduct.ID), utils.Approved)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	modelProduct.Status = utils.Approved

	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Reject product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/reject/{id} [put]
func (h *HandlersProducts) Reject(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Get product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	// Reject product
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(uint64(modelProduct.ID), utils.Rejected)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Submit product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/publish/{id} [put]
func (h *HandlersProducts) Submit(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Get product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	// Get variations in the product
	modelVars := make([]models.VariationsWithAttributeValue, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByProduct(&modelVars, productID)
	if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
		responses.ErrorResponse(c, statusCoede, message)
	}

	// Submit changes when any variations exist.
	if len(modelVars) > 0 {
		prodService := prodsvc.NewServiceProduct(h.server.DB)
		err = prodService.UpdateStatus(uint64(modelProduct.ID), utils.Pending)
		if statusCoede, message := errhandle.SqlErrorHandler(err); statusCoede != 0 {
			responses.ErrorResponse(c, statusCoede, message)
		}
	}

	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Delete product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/{id} [delete]
func (h *HandlersProducts) Delete(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Checkout product status if it's on pending status
	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	// Delete product by ID
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.Delete(productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.SuccessDeleteProduct)
}

// Refresh godoc
// @Summary Edit categories of product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductCategory true "Product Info"
// @Success 200 {object} []responses.ResponseCategory
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/category/{id} [put]
func (h *HandlersProducts) UpdateCategories(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestProductCategory)
	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelCategories := make([]models.ProductCategoriesWithName, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	err = cateRepo.ReadByProductID(&modelCategories, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	cateService := prodcatesvc.NewServiceProductCategory(h.server.DB)
	err = cateService.Update(&modelCategories, req, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseProductCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Edit related channels
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductChannel true "Product Channel Info"
// @Success 200 {object} []responses.ResponseProductChannel
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/channel/{id} [put]
func (h *HandlersProducts) UpdateRelatedChannels(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestProductChannel)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	// Update product related channels
	chanService := chansvc.NewServiceProductChannel(h.server.DB)
	err = chanService.Update(productID, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read related channels in product
	modelChannels := make([]models.ProductChannelsWithName, 0)
	chanRepo := repositories.NewRepositoryProductChannel(h.server.DB)
	err = chanRepo.ReadByProductID(&modelChannels, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Change product status to draft
	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}
	return responses.NewResponseProductChannels(c, http.StatusOK, modelChannels)
}

// Refresh godoc
// @Summary Edit related contents
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductContent true "Product Content Info"
// @Success 200 {object} []responses.ResponseProductContent
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/content/{id} [put]
func (h *HandlersProducts) UpdateRelatedContents(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestProductContent)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Check product if it's on pending status
	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	// Upate product related content
	contService := contsvc.NewServiceProductContent(h.server.DB)
	err = contService.Update(productID, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read product related content
	modelContents := make([]models.ProductContentsWithTitle, 0)
	contRepo := repositories.NewRepositoryProductContent(h.server.DB)
	err = contRepo.ReadByProductID(&modelContents, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Change product status to draft
	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}
	return responses.NewResponseProductContents(c, http.StatusOK, modelContents)
}

// Refresh godoc
// @Summary Edit tags
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestTag true "Tags"
// @Success 200 {object} []responses.ResponseTag
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/tag/{id} [put]
func (h *HandlersProducts) UpdateTags(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestProductTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Check product if it is on pending status
	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	// Read tags of the product
	modelTags := make([]models.ProductTagsWithName, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	err = tagRepo.ReadByProductID(&modelTags, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update tags on the product
	tagService := prodtagsvc.NewServiceProductTag(h.server.DB)
	err = tagService.Update(&modelTags, req, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Change product status to draft
	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}
	return responses.NewResponseProductTags(c, http.StatusOK, modelTags)
}

// Refresh godoc
// @Summary Add attribute
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestAttribute true "Attributes"
// @Success 201 {object} []responses.ResponseAttribute
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/attribute/{id} [post]
func (h *HandlersProducts) CreateAttribute(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Check duplicate attribute in product
	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	err = attrRepo.ReadByProductIDAndName(&modelAttr, productID, req.Name)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != 404 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedProductAttribute)
	}

	// Create attribute
	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	err = attrService.Create(productID, req, &modelAttr)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read all attributes in product
	modelAttrs := make([]models.Attributes, 0)
	err = attrRepo.ReadByProductID(&modelAttrs, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update product status to draft
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(productID, utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAttributes(c, http.StatusCreated, modelAttrs)
}

// Refresh godoc
// @Summary Edit attributes
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param attribute_id query int true "Attribute ID"
// @Param params body requests.RequestAttribute true "Attributes"
// @Success 200 {object} []responses.ResponseAttribute
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/attribute/{id} [put]
func (h *HandlersProducts) UpdateAttributes(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	attributeID, err := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Check duplicate attribute in product
	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	err = attrRepo.ReadByProductIDAndName(&modelAttr, productID, req.Name)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != http.StatusNotFound {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedProductAttribute)
	}

	// Read attribute by id
	err = attrRepo.ReadByID(&modelAttr, attributeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if modelAttr.ProductID != productID {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Update attribute
	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	err = attrService.Update(attributeID, req, &modelAttr)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read all attribute in product
	modelAttrs := make([]models.Attributes, 0)
	err = attrRepo.ReadByProductID(&modelAttrs, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update product status to draft
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(productID, utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAttributes(c, http.StatusOK, modelAttrs)
}

// Refresh godoc
// @Summary Delete attributes
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param attribute_id query int true "Attribute ID"
// @Param id path int true "Product ID"
// @Success 200 {object} []responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/attribute/{id} [delete]
func (h *HandlersProducts) DeleteAttributes(c echo.Context) error {
	// Get attribute ID
	attributeID, err := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read attribute by ID
	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	err = attrRepo.ReadByID(&modelAttr, attributeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if modelAttr.ProductID != productID {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Delete attribute by ID
	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	err = attrService.Delete(attributeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update product status to draft
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(productID, utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAttribute(c, http.StatusOK, modelAttr)
}

// Refresh godoc
// @Summary Add attribute value
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param attribute_id query int true "Attribute ID"
// @Param value query string true "Attribute Value"
// @Success 201 {object} []responses.ResponseAttributeValue
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/attribute-value/{id} [post]
func (h *HandlersProducts) CreateAttributeValueByID(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	attributeID, err := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	value := c.QueryParam("value")

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Check duplicate attribute value
	modelValue := models.AttributeValuesWithDetail{}
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	err = valRepo.ReadByAttrIDAndValue(&modelValue, attributeID, value)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != 404 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedProductAttribute)
	}

	// Create attribute value
	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	err = valService.Create(attributeID, value)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read all attribute values in product
	modelValues := make([]models.AttributeValuesWithDetail, 0)
	err = valRepo.ReadByProductID(&modelValues, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update product status to draft
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(productID, utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAttributeValueByProduct(c, http.StatusCreated, modelValues)
}

// Refresh godoc
// @Summary Edit attribute value
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param attribute_value_id query int true "Attribute Value ID"
// @Param value query string true "Attribute Value"
// @Success 200 {object} []responses.ResponseAttributeValue
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/attribute-value/{id} [put]
func (h *HandlersProducts) UpdateAttributeValueByID(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	attributeValueID, err := strconv.ParseUint(c.QueryParam("attribute_value_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	value := c.QueryParam("value")

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Check duplicate attribute value
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	modelVal := models.AttributeValuesWithDetail{}
	err = valRepo.ReadByAttrValID(&modelVal, attributeValueID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if modelVal.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read attribute by attribute id in attribute value
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	modelAttr := models.Attributes{}
	err = attrRepo.ReadByID(&modelAttr, uint64(modelVal.AttributeID))
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if modelAttr.ProductID != productID {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Check duplicate attribute value
	modelVal.ID = 0
	err = valRepo.ReadByAttrIDAndValue(&modelVal, modelVal.AttributeID, value)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != http.StatusNotFound {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedProductAttributeValue)
	}

	// Update attribute value
	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	err = valService.UpdateByID(attributeValueID, value)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read all attribute value on product
	modelValues := make([]models.AttributeValuesWithDetail, 0)
	repositories.NewRepositoryAttributeValue(h.server.DB)
	err = valRepo.ReadByProductID(&modelValues, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update product status to draft
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(productID, utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAttributeValueByProduct(c, http.StatusOK, modelValues)
}

// Refresh godoc
// @Summary Delete attribute value
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param attribute_value_id query int true "Attribute Value ID"
// @Success 200 {object} []responses.ResponseAttributeValue
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/attribute-value/{id} [delete]
func (h *HandlersProducts) DeleteAttributeValueByID(c echo.Context) error {
	// Get product ID
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Get attribute value ID
	attributeValueID, err := strconv.ParseUint(c.QueryParam("attribute_value_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read product by ID
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read attribute value by ID
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	modelVal := models.AttributeValuesWithDetail{}
	err = valRepo.ReadByAttrValID(&modelVal, attributeValueID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if modelVal.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read attribute by attribute ID in attribute value
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	modelAttr := models.Attributes{}
	err = attrRepo.ReadByID(&modelAttr, uint64(modelVal.AttributeID))
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if modelAttr.ProductID != productID {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Delete attribute value by ID
	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	err = valService.DeleteByID(attributeValueID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read all attribute values in the product
	modelValues := make([]models.AttributeValuesWithDetail, 0)
	err = valRepo.ReadByProductID(&modelValues, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update product status to draft
	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateStatus(productID, utils.Draft)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAttributeValueByProduct(c, http.StatusOK, modelValues)
}

// Refresh godoc
// @Summary Add shipping data
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Shipping Data"
// @Success 201 {object} responses.ResponseShippingData
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/shipping/{id} [post]
func (h *HandlersProducts) CreateShippingData(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	err = shipRepo.ReadByVariationID(&modelShipData, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != 404 {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedProductShippingData)
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Create(productID, req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseShippingData(c, http.StatusCreated, modelShipData)
}

// Refresh godoc
// @Summary Edit shipping data
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Review"
// @Success 200 {object} responses.ResponseShippingData
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/shipping/{id} [put]
func (h *HandlersProducts) UpdateShippingData(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	err = shipRepo.ReadByVariationID(&modelShipData, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != 404 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	err = shipService.Update(productID, req, &modelShipData)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 && statusCode != 404 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}
	return responses.NewResponseShippingData(c, http.StatusOK, modelShipData)
}

// Refresh godoc
// @Summary Delete shipping data
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Review"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/shipping/{id} [delete]
func (h *HandlersProducts) DeleteShippingData(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	err = shipRepo.ReadByVariationID(&modelShipData, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	err = shipService.Delete(productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	err = ChangeToDraft(h.server.DB, &modelProduct)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}
	return responses.MessageResponse(c, http.StatusOK, constants.SuccessDeleteShippingData)
}

// Refresh godoc
// @Summary Create linked product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Param link_id query int true "Linked product ID"
// @Param is_up_cross query string true "Is Up-Sell or Cross-Sell"
// @Success 201 {object} responses.ResponseLinkedProducts
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/linked [post]
func (h *HandlersProducts) CreateLinkedProduct(c echo.Context) error {
	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	linkID, err := strconv.ParseUint(c.QueryParam("link_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	sellType := c.QueryParam("is_up_cross")

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	linkService := linksvc.NewServiceLink(h.server.DB)
	err = linkService.Create(productID, linkID, utils.SellTypesFromString(sellType))
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadLinkedProducts(&modelProducts, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseLinkedProducts(c, http.StatusCreated, modelProducts)
}

// Refresh godoc
// @Summary Read linked products
// @Tags Product Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} responses.ResponseLinkedProducts
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/product/linked [get]
func (h *HandlersProducts) ReadLinkedProduct(c echo.Context) error {
	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)

	err = prodRepo.ReadLinkedProducts(&modelProducts, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseLinkedProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Delete linked product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int false "ID"
// @Param product_id query int true "Product ID"
// @Success 200 {object} responses.ResponseLinkedProducts
// @Success 400 {object} responses.Error
// @Success 404 {object} responses.Error
// @Success 500 {object} responses.Error
// @Router /store/api/v1/product/linked/{id} [delete]
func (h *HandlersProducts) DeleteLinkedProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	linkService := linksvc.NewServiceLink(h.server.DB)
	err = linkService.Delete(id)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadLinkedProducts(&modelProducts, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseLinkedProducts(c, http.StatusOK, modelProducts)
}
