package handlers

import (
	"OnlineStoreBackend/models"
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

func ChangeToDraft(db *gorm.DB, modelProduct *models.Products) {
	if modelProduct.Status == utils.Approved || modelProduct.Status == utils.Rejected {
		prodService := prodsvc.NewServiceProduct(db)
		prodService.UpdateStatus(uint64(modelProduct.ID), utils.Draft)
	}
}

func CheckProduct(db *gorm.DB, modelProduct *models.Products, productID uint64) string {
	prodRepo := repositories.NewRepositoryProduct(db)
	err := prodRepo.ReadByID(modelProduct, productID)

	if err != gorm.ErrRecordNotFound {
		return "Product doesn't exist at this ID."
	}
	if modelProduct.Status == utils.Pending {
		return "This product is on pending status."
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
// @Router /store/api/v1/product [post]
func (h *HandlersProducts) Create(c echo.Context) error {
	req := new(requests.RequestProductWithDetail)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.ProductsWithDetail{}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	prodService.Create(&modelProduct.Products, req)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadDetail(&modelProduct, uint64(modelProduct.ID))

	prodService.UpdateStatus(uint64(modelProduct.ID), utils.Draft)
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
// @Router /store/api/v1/product/{id} [get]
func (h *HandlersProducts) ReadByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	modelProduct := models.ProductsWithDetail{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadDetail(&modelProduct, productID)
	return responses.NewResponseProductWithDetail(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Read all products
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int false "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} []responses.ResponseProduct
// @Router /store/api/v1/product [get]
func (h *HandlersProducts) ReadAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelProducts := make([]models.Products, 0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadAll(&modelProducts, storeID, keyword)

	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Read approved products
// @Tags Product Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param page query int false "Page"
// @Param count query int false "Count"
// @Success 200 {object} responses.ResponseProductApprovedPaging
// @Router /store/api/v1/product/approved  [get]
func (h *HandlersProducts) ReadApproved(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	count, _ := strconv.Atoi(c.QueryParam("count"))

	exchangeRate, currencyCode := 0.0, "$"
	modelProducts := make([]models.ProductsApproved, 0)

	totalCount := int64(0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadApproved(&modelProducts, storeID, customerID, page, count, &totalCount)

	taxRepo := repositories.NewRepositoryTax(h.server.DB)
	taxRepo.ReadCurrency(&currencyCode, &exchangeRate, customerID)

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
// @Router /store/api/v1/product/category [get]
func (h *HandlersProducts) ReadByCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	cateID, _ := strconv.ParseUint(c.QueryParam("category_id"), 10, 64)

	modelProducts := make([]models.Products, 0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByCategory(&modelProducts, storeID, cateID)

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
// @Router /store/api/v1/product/search [get]
func (h *HandlersProducts) ReadSearch(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	keyword := c.QueryParam("keyword")
	tags := strings.Split(c.QueryParam("tags"), ",")

	modelProducts := make([]models.Products, 0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByTags(&modelProducts, storeID, tags, keyword)

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
// @Router /store/api/v1/product/paging [get]
func (h *HandlersProducts) ReadPaging(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	count, _ := strconv.ParseInt(c.QueryParam("count"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	totalCount := int64(0)
	modelProducts := make([]models.Products, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadPaging(&modelProducts, int(page), int(count), storeID, keyword, &totalCount)
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
// @Router /store/api/v1/product/{id} [put]
func (h *HandlersProducts) Update(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	req := new(requests.RequestProduct)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	if err := prodService.Update(&modelProduct, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/approve/{id} [put]
func (h *HandlersProducts) Approve(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, productID)

	if modelProduct.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at this ID.")
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	prodService.UpdateStatus(uint64(modelProduct.ID), utils.Approved)

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
// @Router /store/api/v1/product/reject/{id} [put]
func (h *HandlersProducts) Reject(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, productID)

	if modelProduct.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at this ID.")
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	prodService.UpdateStatus(uint64(modelProduct.ID), utils.Rejected)

	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Publish product
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/publish/{id} [put]
func (h *HandlersProducts) Publish(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, productID)

	if modelProduct.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at this ID.")
	}

	modelVars := make([]models.VariationsWithAttributeValue, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByProduct(&modelVars, productID)
	if len(modelVars) > 0 {
		prodService := prodsvc.NewServiceProduct(h.server.DB)
		prodService.UpdateStatus(uint64(modelProduct.ID), utils.Pending)
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
// @Router /store/api/v1/product/{id} [delete]
func (h *HandlersProducts) Delete(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	prodService.Delete(productID)
	ChangeToDraft(h.server.DB, &modelProduct)
	return responses.MessageResponse(c, http.StatusOK, "Product successfully deleted.")
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
// @Router /store/api/v1/product/category/{id} [put]
func (h *HandlersProducts) UpdateCategories(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	req := new(requests.RequestProductCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelCategories := make([]models.ProductCategoriesWithName, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByProductID(&modelCategories, productID)

	cateService := prodcatesvc.NewServiceProductCategory(h.server.DB)
	cateService.Update(&modelCategories, req, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/channel/{id} [put]
func (h *HandlersProducts) UpdateRelatedChannels(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductChannel)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	chanService := chansvc.NewServiceProductChannel(h.server.DB)
	chanService.Update(productID, req)

	modelChannels := make([]models.ProductChannelsWithName, 0)
	chanRepo := repositories.NewRepositoryProductChannel(h.server.DB)
	chanRepo.ReadByProductID(&modelChannels, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/content/{id} [put]
func (h *HandlersProducts) UpdateRelatedContents(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductContent)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	contService := contsvc.NewServiceProductContent(h.server.DB)
	contService.Update(productID, req)

	modelContents := make([]models.ProductContentsWithTitle, 0)
	contRepo := repositories.NewRepositoryProductContent(h.server.DB)
	contRepo.ReadByProductID(&modelContents, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/tag/{id} [put]
func (h *HandlersProducts) UpdateTags(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelTags := make([]models.ProductTagsWithName, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByProductID(&modelTags, productID)

	tagService := prodtagsvc.NewServiceProductTag(h.server.DB)
	tagService.Update(&modelTags, req, &modelProduct)

	ChangeToDraft(h.server.DB, &modelProduct)
	return responses.NewResponseProductTags(c, http.StatusOK, modelTags)
}

// Refresh godoc
// @Summary Add attributes
// @Tags Product Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestAttribute true "Attributes"
// @Success 201 {object} []responses.ResponseAttribute
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/attribute/{id} [post]
func (h *HandlersProducts) CreateAttributes(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByName(&modelAttr, req.Name)

	if modelAttr.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute already exists in the product.")
	}

	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	attrService.Create(productID, req, &modelAttr)

	modelAttrs := make([]models.Attributes, 0)
	attrRepo.ReadByProductID(&modelAttrs, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/attribute/{id} [put]
func (h *HandlersProducts) UpdateAttributes(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByID(&modelAttr, attributeID)

	if modelAttr.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute doesn't exists in the product.")
	}

	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	attrService.Update(attributeID, req, &modelAttr)

	modelAttrs := make([]models.Attributes, 0)
	attrRepo.ReadByProductID(&modelAttrs, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/attribute/{id} [delete]
func (h *HandlersProducts) DeleteAttributes(c echo.Context) error {
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByID(&modelAttr, attributeID)

	if modelAttr.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute doesn't exists in the product.")
	}

	attrService := prodattrsvc.NewServiceAttribute(h.server.DB)
	attrService.Delete(attributeID)

	ChangeToDraft(h.server.DB, &modelProduct)
	return responses.NewResponseAttribute(c, http.StatusOK, modelAttr)
}

func (h *HandlersProducts) UpdateAttributeValues(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	req := new(requests.RequestAttributeValue)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByID(&modelAttr, attributeID)

	if modelAttr.ID == 0 || modelAttr.ProductID != productID {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute doesn't exists in the product.")
	}

	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	valService.Update(attributeID, req)

	modelVals := make([]models.AttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	valRepo.ReadByID(&modelVals, attributeID)
	valRepo.ReadByProductID(&modelVals, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
	return responses.NewResponseAttributeValueByProduct(c, http.StatusOK, modelVals)
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
// @Success 200 {object} []responses.ResponseAttributeValue
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/attribute-value/{id} [post]
func (h *HandlersProducts) CreateAttributeValueByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	value := c.QueryParam("value")

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	if err := valService.Create(attributeID, value); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute value doesn't exist.")
	}

	modelValues := make([]models.AttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	valRepo.ReadByProductID(&modelValues, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/attribute-value/{id} [put]
func (h *HandlersProducts) UpdateAttributeValueByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeValueID, _ := strconv.ParseUint(c.QueryParam("attribute_value_id"), 10, 64)
	value := c.QueryParam("value")

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	if err := valService.UpdateByID(attributeValueID, value); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute value doesn't exist.")
	}

	modelValues := make([]models.AttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	valRepo.ReadByProductID(&modelValues, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/attribute-value/{id} [delete]
func (h *HandlersProducts) DeleteAttributeValueByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeValueID, _ := strconv.ParseUint(c.QueryParam("attribute_value_id"), 10, 64)

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	valService := prodattrvalsvc.NewServiceAttributeValue(h.server.DB)
	if err := valService.DeleteByID(attributeValueID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute value doesn't exist.")
	}

	modelValues := make([]models.AttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	valRepo.ReadByProductID(&modelValues, productID)

	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/shipping/{id} [post]
func (h *HandlersProducts) CreateShippingData(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	shipRepo.ReadByVariationID(&modelShipData, productID)
	if modelShipData.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping data already exists in this product.")
	}
	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Create(productID, req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/shipping/{id} [put]
func (h *HandlersProducts) UpdateShippingData(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	shipRepo.ReadByVariationID(&modelShipData, productID)
	if modelShipData.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping data doesn't exist in this product.")
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Update(productID, req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	ChangeToDraft(h.server.DB, &modelProduct)
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
// @Router /store/api/v1/product/shipping/{id} [delete]
func (h *HandlersProducts) DeleteShippingData(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	shipRepo.ReadByVariationID(&modelShipData, productID)
	if modelShipData.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping data doesn't exist in this product.")
	}
	if modelProduct.Status == utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This product is on pending status.")
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Delete(productID); err != nil {
		return responses.MessageResponse(c, http.StatusOK, "Failed to delete shipping data")
	}
	ChangeToDraft(h.server.DB, &modelProduct)
	return responses.MessageResponse(c, http.StatusOK, "Shipping data is successfully deleted")
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
// @Router /store/api/v1/product/linked [post]
func (h *HandlersProducts) CreateLinkedProduct(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	linkID, _ := strconv.ParseUint(c.QueryParam("link_id"), 10, 64)
	sellType := c.QueryParam("is_up_cross")

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	linkService := linksvc.NewServiceLink(h.server.DB)
	if err := linkService.Create(productID, linkID, utils.SellTypesFromString(sellType)); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadLinkedProducts(&modelProducts, productID)

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
// @Router /store/api/v1/product/linked [get]
func (h *HandlersProducts) ReadLinkedProduct(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadLinkedProducts(&modelProducts, productID)

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
// @Router /store/api/v1/product/linked/{id} [delete]
func (h *HandlersProducts) DeleteLinkedProduct(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	modelProduct := models.Products{}
	if message := CheckProduct(h.server.DB, &modelProduct, productID); message != "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	linkService := linksvc.NewServiceLink(h.server.DB)
	linkService.Delete(id)

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadLinkedProducts(&modelProducts, productID)

	return responses.NewResponseLinkedProducts(c, http.StatusOK, modelProducts)
}
