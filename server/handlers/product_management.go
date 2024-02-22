package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodattrsvc "OnlineStoreBackend/services/product_attributes"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	linkedsvc "OnlineStoreBackend/services/product_linked"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	prodvarsvc "OnlineStoreBackend/services/product_variations"
	prodsvc "OnlineStoreBackend/services/products"
	chansvc "OnlineStoreBackend/services/related_channels"
	contsvc "OnlineStoreBackend/services/related_contents"
	shipsvc "OnlineStoreBackend/services/shipping_data"
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
// @Summary Add product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestProduct true "Product Info"
// @Success 201 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product [post]
func (h *HandlersProductManagement) Create(c echo.Context) error {
	req := new(requests.RequestProduct)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}

	serviceProduct := prodsvc.NewServiceProduct(h.server.DB)
	serviceProduct.Create(&modelProduct, req)

	return responses.NewResponseProduct(c, http.StatusCreated, modelProduct)
}

// Refresh godoc
// @Summary Read product by ID
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseProductWithDetail
// @Router /store/api/v1/product/{id} [get]
func (h *HandlersProductManagement) ReadByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	modelProduct := models.ProductsWithDetail{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadDetail(&modelProduct, productID)
	return responses.NewResponseProductWithDetail(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Read all products
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int false "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} []responses.ResponseProduct
// @Router /store/api/v1/product [get]
func (h *HandlersProductManagement) ReadAll(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	modelProducts := make([]models.Products, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadAll(&modelProducts, storeID, keyword)
	return responses.NewResponseProducts(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Read products by pagination
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int true "Page" default(0)
// @Param count query int true "Count" default(100)
// @Param store_id query int false "Store ID"
// @Param keyword query string false "Keyword"
// @Success 200 {object} responses.ResponseProductsPaging
// @Router /store/api/v1/product/paging [get]
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
// @Summary Edit product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProduct true "Product Info"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/{id} [put]
func (h *HandlersProductManagement) Update(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	req := new(requests.RequestProduct)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	modelProduct := models.Products{}
	modelProduct.ID = uint(productID)
	service := prodsvc.NewServiceProduct(h.server.DB)
	if err := service.Update(&modelProduct, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Delete product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/{id} [delete]
func (h *HandlersProductManagement) Delete(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	service := prodsvc.NewServiceProduct(h.server.DB)
	if err := service.Delete(productID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusOK, "Product successfully deleted.")
}

// Refresh godoc
// @Summary Edit categories of product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductCategory true "Product Info"
// @Success 200 {object} []responses.ResponseCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/category/{id} [put]
func (h *HandlersProductManagement) UpdateCategories(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	req := new(requests.RequestProductCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCategories := make([]models.ProductCategoriesWithName, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByProductID(&modelCategories, productID)

	service := prodcatesvc.NewServiceProductCategory(h.server.DB)
	service.Update(&modelCategories, req, productID)
	return responses.NewResponseProductCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Edit related channels
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductChannel true "Product Channel Info"
// @Success 200 {object} []responses.ResponseProductChannel
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/channel/{id} [put]
func (h *HandlersProductManagement) UpdateRelatedChannels(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductChannel)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProdut := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProdut, productID)
	if modelProdut.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at ths ID.")
	}

	modelChannels := make([]models.ProductChannelsWithName, 0)
	chanRepo := repositories.NewRepositoryProductChannel(h.server.DB)
	chanRepo.ReadByProductID(&modelChannels, productID)

	chanService := chansvc.NewServiceProductChannel(h.server.DB)
	chanService.Update(&modelChannels, req, productID)
	return responses.NewResponseProductChannels(c, http.StatusOK, modelChannels)
}

// Refresh godoc
// @Summary Edit related contents
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestProductContent true "Product Content Info"
// @Success 200 {object} []responses.ResponseProductContent
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/content/{id} [put]
func (h *HandlersProductManagement) UpdateRelatedContents(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductContent)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProdut := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProdut, productID)
	if modelProdut.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at ths ID.")
	}

	modelContents := make([]models.ProductContentsWithTitle, 0)
	contRepo := repositories.NewRepositoryProductContent(h.server.DB)
	contRepo.ReadByProductID(&modelContents, productID)

	contService := contsvc.NewServiceProductContent(h.server.DB)
	contService.Update(&modelContents, req, productID)
	return responses.NewResponseProductContents(c, http.StatusOK, modelContents)
}

// Refresh godoc
// @Summary Edit tags
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestTag true "Tags"
// @Success 200 {object} []responses.ResponseTag
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/tag/{id} [put]
func (h *HandlersProductManagement) UpdateTags(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTags := make([]models.ProductTagsWithName, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByProductID(&modelTags, productID)

	tagService := prodtagsvc.NewServiceProductTag(h.server.DB)
	tagService.Update(&modelTags, req, productID)
	return responses.NewResponseProductTags(c, http.StatusOK, modelTags)
}

// Refresh godoc
// @Summary Add attributes
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestAttribute true "Attributes"
// @Success 201 {object} []responses.ResponseProductAttribute
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/attribute/{id} [post]
func (h *HandlersProductManagement) CreateAttributes(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAttr := models.ProductAttributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByName(&modelAttr, req.Name)

	if modelAttr.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute already exists in the product.")
	}

	attrService := prodattrsvc.NewServiceProductAttribute(h.server.DB)
	attrService.Create(productID, req, &modelAttr)

	modelAttrs := make([]models.ProductAttributes, 0)
	attrRepo.ReadByProductID(&modelAttrs, productID)

	return responses.NewResponseProductAttributes(c, http.StatusCreated, modelAttrs)
}

// Refresh godoc
// @Summary Edit attributes
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param attribute_id query int true "Attribute ID"
// @Param params body requests.RequestAttribute true "Attributes"
// @Success 200 {object} []responses.ResponseProductAttribute
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/attribute/{id} [put]
func (h *HandlersProductManagement) UpdateAttributes(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	req := new(requests.RequestAttribute)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAttr := models.ProductAttributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByID(&modelAttr, attributeID)

	if modelAttr.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute doesn't exists in the product.")
	}

	attrService := prodattrsvc.NewServiceProductAttribute(h.server.DB)
	attrService.Update(attributeID, req, &modelAttr)

	modelAttrs := make([]models.ProductAttributes, 0)
	attrRepo.ReadByProductID(&modelAttrs, productID)

	return responses.NewResponseProductAttributes(c, http.StatusOK, modelAttrs)
}

// Refresh godoc
// @Summary Delete attributes
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param attribute_id query int true "Attribute ID"
// @Success 200 {object} []responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/attribute [delete]
func (h *HandlersProductManagement) DeleteAttributes(c echo.Context) error {
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)

	modelAttr := models.ProductAttributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByID(&modelAttr, attributeID)

	if modelAttr.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute doesn't exists in the product.")
	}

	attrService := prodattrsvc.NewServiceProductAttribute(h.server.DB)
	attrService.Delete(attributeID)

	return responses.NewResponseProductAttribute(c, http.StatusOK, modelAttr)
}

// Refresh godoc
// @Summary Edit variations
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param attribute_id query int true "Attribute ID"
// @Param params body requests.RequestProductVariation true "Attributes"
// @Success 200 {object} []responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/variation/{id} [put]
func (h *HandlersProductManagement) UpdateVariations(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	attributeID, _ := strconv.ParseUint(c.QueryParam("attribute_id"), 10, 64)
	req := new(requests.RequestProductVariation)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAttr := models.ProductAttributes{}
	attrRepo := repositories.NewRepositoryAttribute(h.server.DB)
	attrRepo.ReadByID(&modelAttr, attributeID)

	if modelAttr.ID == 0 || modelAttr.ProductID != productID {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This attribute doesn't exists in the product.")
	}

	modelVars := make([]models.ProductVariationsWithName, 0)
	varRepo := repositories.NewRepositoryProductVariation(h.server.DB)
	varRepo.ReadByID(&modelVars, attributeID)

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.Update(attributeID, productID, &modelVars, req)

	varRepo.ReadByProductID(&modelVars, productID)
	return responses.NewResponseProductVariations(c, http.StatusOK, modelVars)
}

// Refresh godoc
// @Summary Set minimum stock level of product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestMinimumStockLevel true "Minimum Stock Level"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/min-stock-level/{id} [put]
func (h *HandlersProductManagement) UpdateMinimumStockLevel(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestMinimumStockLevel)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, productID)
	if modelProduct.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at ths ID.")
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	if err := prodService.UpdateMinimumStockLevel(productID, req, &modelProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Add shipping data
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Shipping Data"
// @Success 201 {object} responses.ResponseShippingData
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/shipping/{id} [post]
func (h *HandlersProductManagement) CreateShippingData(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	shipRepo.ReadByProductID(&modelShipData, productID)
	if modelShipData.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping data already exists in this product.")
	}
	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Create(productID, req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseShippingData(c, http.StatusCreated, modelShipData)
}

// Refresh godoc
// @Summary Edit shipping data
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Review"
// @Success 200 {object} responses.ResponseShippingData
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/shipping/{id} [put]
func (h *HandlersProductManagement) UpdateShippingData(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingData)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	shipRepo.ReadByProductID(&modelShipData, productID)
	if modelShipData.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping data doesn't exist in this product.")
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Update(productID, req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseShippingData(c, http.StatusOK, modelShipData)
}

// Refresh godoc
// @Summary Delete shipping data
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param params body requests.RequestShippingData true "Review"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/product/shipping/{id} [delete]
func (h *HandlersProductManagement) DeleteShippingData(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
	shipRepo.ReadByProductID(&modelShipData, productID)
	if modelShipData.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping data doesn't exist in this product.")
	}

	shipService := shipsvc.NewServiceShippingData(h.server.DB)
	if err := shipService.Delete(productID); err != nil {
		return responses.MessageResponse(c, http.StatusOK, "Failed to delete shipping data")
	}
	return responses.MessageResponse(c, http.StatusOK, "Shipping data is successfully deleted")
}

// Refresh godoc
// @Summary Create linked product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int false "product ID"
// @Param linked_id query int false "linked product ID"
// @Param is_up_cross query int false "is up-sell or cross-sell"
// @Success 201 {object} responses.ResponseProductLinked
// @Router /store/api/v1/product/linked [post]
func (h *HandlersProductManagement) CreateLinkedProduct(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	linkedID, _ := strconv.ParseUint(c.QueryParam("linked_id"), 10, 64)
	isUpCross, _ := strconv.ParseUint(c.QueryParam("is_up_cross"), 10, 64)

	modelProductLinked := models.ProductLinked{}

	linkedService := linkedsvc.NewServiceProductLinked(h.server.DB)
	if err := linkedService.Create(productID, linkedID, isUpCross, &modelProductLinked); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadLinkedProducts(&modelProducts, productID)

	return responses.NewResponseProductLinked(c, http.StatusCreated, modelProducts)
}

// Refresh godoc
// @Summary Read linked products
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int false "product ID"
// @Success 200 {object} responses.ResponseProductLinked
// @Router /store/api/v1/product/linked [get]
func (h *HandlersProductManagement) ReadLinkedProduct(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadLinkedProducts(&modelProducts, productID)

	return responses.NewResponseProductLinked(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Delete linked product
// @Tags Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int false "ID"
// @Param product_id query int false "product ID"
// @Success 200 {object} responses.ResponseProductLinked
// @Router /store/api/v1/product/linked/{id} [delete]
func (h *HandlersProductManagement) DeleteLinkedProduct(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	prodService := linkedsvc.NewServiceProductLinked(h.server.DB)
	prodService.Delete(id)

	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadLinkedProducts(&modelProducts, productID)

	return responses.NewResponseProductLinked(c, http.StatusOK, modelProducts)
}
