package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
	"encoding/json"
	"strconv"
	"strings"
)

func (service *Service) Create(modelVar *models.ProductVariations, req *requests.RequestProductVariation, productID uint64) {
	modelValues := make([]models.ProductAttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryProductAttributeValue(service.DB)
	valRepo.ReadByIDs(&modelValues, req.AttributeValueIDs)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	prodRepo.ReadByID(&modelProduct, productID)

	sku := modelProduct.Title
	title := modelProduct.Title
	for index, modelValue := range modelValues {
		sku += modelValue.AttributeValue
		if index == 0 {
			title += " - "
		} else {
			title += ", "
		}
		title += modelValue.AttributeValue
	}
	sku = utils.CleanSpecialLetters(sku)
	imageUrls, _ := json.Marshal(req.ImageUrls)

	varRepo := repositories.NewRepositoryVariation(service.DB)
	varRepo.ReadByAttributeValueIDs(modelVar, req.AttributeValueIDs, productID)

	if modelVar.ID == 0 {
		modelVar.Sku = sku
		modelVar.Title = title
		modelVar.ProductID = productID
		modelVar.Price = req.Price
		modelVar.ImageUrls = string(imageUrls)
		modelVar.DiscountAmount = req.DiscountAmount
		modelVar.DiscountType = req.DiscountType
		modelVar.StockLevel = req.StockLevel
		modelVar.Description = req.Description
		modelVar.BackOrderStatus = utils.SimpleStatuses(req.BackOrderAllowed)

		service.DB.Create(&modelVar)
		detService := prodvardetsvc.NewServiceProductVariationDetail(service.DB)
		detService.Create(uint64(modelVar.ID), req.AttributeValueIDs)
	}
}

func (service *Service) CreateSimpleWithCSV(modelVar *models.ProductVariations, modelCsv *models.CSVs, productID uint64) {
	modelVar.ProductID = productID
	modelVar.Sku = modelCsv.Sku
	modelVar.Price, _ = strconv.ParseFloat(modelCsv.RegularPrice, 64)
	modelVar.StockLevel, _ = strconv.ParseFloat(modelCsv.Stock, 64)
	if modelCsv.SalePrice != "" {
		salePrice, _ := strconv.ParseFloat(modelCsv.SalePrice, 64)
		modelVar.DiscountAmount = modelVar.Price - salePrice
	} else {
		modelVar.DiscountAmount = 0
	}
	if modelCsv.BackordersAllowed == "1" {
		modelVar.BackOrderStatus = utils.Enabled
	} else {
		modelVar.BackOrderStatus = utils.Disabled
	}
	modelVar.DiscountType = utils.FixedAmountOff
	modelVar.Title = modelCsv.Name
	service.DB.Create(modelVar)

	if modelCsv.Weight != "" {
		shipService := shipsvc.NewServiceShippingData(service.DB)
		shipService.CreateWithCSV(uint64(modelVar.ID), modelCsv)
	}
}

func (service *Service) CreateVariableWithCSV(modelVar *models.ProductVariations, modelCsv *models.CSVs, productID uint64, modelVals *[]models.ProductAttributeValues) {
	images := strings.Split(modelCsv.Images, ",")
	imageUrls, _ := json.Marshal(images)
	modelVar.Price, _ = strconv.ParseFloat(modelCsv.RegularPrice, 64)
	if modelCsv.SalePrice != "" {
		salePrice, _ := strconv.ParseFloat(modelCsv.SalePrice, 64)
		modelVar.DiscountAmount = modelVar.Price - salePrice
	} else {
		modelVar.DiscountAmount = 0
	}
	modelVar.ProductID = productID
	modelVar.Sku = modelCsv.Sku
	modelVar.StockLevel, _ = strconv.ParseFloat(modelCsv.Stock, 64)
	modelVar.DiscountType = utils.FixedAmountOff
	modelVar.Description = modelCsv.Description
	modelVar.ImageUrls = string(imageUrls)
	modelVar.Title = modelCsv.Name
	service.DB.Create(modelVar)

	attributeValueIDs := make([]uint64, 0)
	for _, modelVal := range *modelVals {
		attributeValueIDs = append(attributeValueIDs, uint64(modelVal.ID))
	}

	detService := prodvardetsvc.NewServiceProductVariationDetail(service.DB)
	detService.Create(uint64(modelVar.ID), attributeValueIDs)

	if modelCsv.Weight != "" {
		shipService := shipsvc.NewServiceShippingData(service.DB)
		shipService.CreateWithCSV(uint64(modelVar.ID), modelCsv)
	}
}
