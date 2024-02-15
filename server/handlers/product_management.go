package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodAttr "OnlineStoreBackend/services/product_attributes"
	prodtag "OnlineStoreBackend/services/product_tags"
	prod "OnlineStoreBackend/services/products"
	prodChan "OnlineStoreBackend/services/related_channels"
	prodCont "OnlineStoreBackend/services/related_contents"
	shipData "OnlineStoreBackend/services/shipping_data"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersProductManagement struct {
	server *s.Server
}

func NewHandlersProductManagement(server *s.Server) *HandlersProductManagement {
	return &HandlersProductManagement{server: server}
}

// Refresh godoc
// @Summary Create product
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestProduct true "Product Info"
// @Success 201 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product [post]
func (h *HandlersProductManagement) Create(c echo.Context) error {
	requestProduct := new(requests.RequestProduct)
	if err := c.Bind(requestProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := requestProduct.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	serviceProduct := prod.CreateService(h.server.DB)
	if err := serviceProduct.Create(&modelProduct, requestProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseProduct(c, http.StatusCreated, modelProduct)
}

// Refresh godoc
// @Summary Get product by ID
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProductWithDetail
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/{id} [get]
func (h *HandlersProductManagement) ReadByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	modelProduct := models.ProductsWithDetail{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadDetail(&modelProduct, productID)
	return responses.NewResponseProductWithDetail(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Get all products
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int false "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} []responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product [get]
func (h *HandlersProductManagement) ReadAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	modelProducts := make([]models.Products, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadAll(&modelProducts, storeID, keyword)
	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Get products by pagination
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int true "Page" default(0)
// @Param count query int true "Count" default(100)
// @Param store_id query int false "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} responses.ResponseProductsPaging
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/paging [get]
func (h *HandlersProductManagement) ReadPaging(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	page, _ := strconv.ParseUint(c.QueryParam("page"), 10, 64)
	count, _ := strconv.ParseUint(c.QueryParam("count"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	totalCount := uint64(0)
	modelProducts := make([]models.Products, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadPaging(&modelProducts, page, count, storeID, keyword, &totalCount)
	return responses.NewResponseProductsPaging(c, http.StatusOK, modelProducts, totalCount)
}

// Refresh godoc
// @Summary Update product
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProduct true "Product Info"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/{id} [put]
func (h *HandlersProductManagement) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	requestProduct := new(requests.RequestProduct)
	if err := c.Bind(requestProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := requestProduct.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	modelProduct := models.Products{}
	modelProduct.ID = uint(id)
	service := prod.CreateService(h.server.DB)
	if err := service.Update(&modelProduct, requestProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Delete product
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/{id} [delete]
func (h *HandlersProductManagement) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	service := prod.CreateService(h.server.DB)
	if err := service.Delete(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusOK, "Product successfully deleted.")
}

// Refresh godoc
// @Summary Add related channels
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductChannel true "Product Channel Info"
// @Success 201 {object} responses.ResponseProductChannel
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/channel/{id} [post]
func (h *HandlersProductManagement) CreateRelatedChannels(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductChannel)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelChannels := make([]models.ProductChannels, 0)
	chanService := prodChan.CreateService(h.server.DB)
	if err := chanService.Create(id, req, &modelChannels); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductChannels(c, http.StatusCreated, modelChannels)
}

// Refresh godoc
// @Summary Add related contents
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductContent true "Product Content Info"
// @Success 201 {object} responses.ResponseProductContent
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/content/{id} [post]
func (h *HandlersProductManagement) CreateRelatedContents(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductContent)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelContents := make([]models.ProductContents, 0)
	contService := prodCont.CreateService(h.server.DB)
	if err := contService.Create(id, req, &modelContents); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductContents(c, http.StatusCreated, modelContents)
}

// Refresh godoc
// @Summary Add Tags
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestTags true "Tags"
// @Success 201 {object} []responses.ResponseTag
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/tag/{id} [post]
func (h *HandlersProductManagement) CreateTags(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTags := make([]models.ProductTags, 0)
	tagService := prodtag.CreateService(h.server.DB)
	tagService.Create(productID, req, &modelTags)

	modelTagsWithName := make([]models.ProductTagsWithName, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByProductID(&modelTagsWithName, productID)
	return responses.NewResponseProductTags(c, http.StatusOK, modelTagsWithName)
}

// Refresh godoc
// @Summary Add attributes
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestAttributes true "Attributes"
// @Success 201 {object} []responses.ResponseProductAttribute
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/attribute/{id} [post]
func (h *HandlersProductManagement) CreateAttributes(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAttrs := make([]models.ProductAttributes, 0)
	attrService := prodAttr.CreateService(h.server.DB)
	attrService.Create(id, req, &modelAttrs)

	modelAttrsWithName := make([]models.ProductAttributesWithName, 0)
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByProductID(&modelAttrsWithName, id)
	return responses.NewResponseProductAttributes(c, http.StatusCreated, modelAttrsWithName)
}

// Refresh godoc
// @Summary Edit stock quantity of product
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductQuantity true "Stock Quantity"
// @Success 201 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/quantity/{id} [put]
func (h *HandlersProductManagement) UpdateStockQuantity(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductQuantity)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	prodService := prod.CreateService(h.server.DB)
	if err := prodService.UpdateStockQuantity(id, req, &modelProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Edit unit price of product
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductPrice true "Unit Price"
// @Success 201 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/price/{id} [put]
func (h *HandlersProductManagement) UpdatePrice(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductPrice)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	prodService := prod.CreateService(h.server.DB)
	if err := prodService.UpdatePrice(id, req, &modelProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Add shipping data
// @Tags product management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Review"
// @Success 201 {object} responses.ResponseShippingData
// @Failure 400 {object} responses.Error
// @Router /api/v1/product/shipping/{id} [post]
func (h *HandlersProductManagement) CreateShippingData(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelShipData := models.ShippingData{}
	shipDataService := shipData.CreateService(h.server.DB)
	if err := shipDataService.Create(id, req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseShippingData(c, http.StatusOK, modelShipData)
}
