package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvardetsvc "OnlineStoreBackend/services/product_variation_details"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	"encoding/json"
	"strconv"
	"strings"
)

func GenerateSKU(modelProduct *models.Products, modelValues *[]models.ProductAttributeValuesWithDetail) string {
	sku := modelProduct.Title
	for _, modelValue := range *modelValues {
		lenAttr := len(modelValue.AttributeName)
		lenVal := len(modelValue.AttributeValue)
		if lenAttr > 3 {
			lenAttr = 3
		}
		if lenVal > 3 && modelValue.Unit == "" {
			lenVal = 3
		}
		sku += "-" + modelValue.AttributeName[0:lenAttr] + "-" + modelValue.AttributeValue[0:lenVal]
	}
	return strings.ToUpper(sku)
}

func (service *Service) Create(modelVar *models.ProductVariations, req *requests.RequestProductVariation, productID uint64) {
	modelValues := make([]models.ProductAttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryProductAttributeValue(service.DB)
	valRepo.ReadByIDs(&modelValues, req.AttributeValueIDs)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	prodRepo.ReadByID(&modelProduct, productID)

	sku := GenerateSKU(&modelProduct, &modelValues)

	service.DB.Where("sku = ?", sku).First(&modelVar)

	if modelVar.ID == 0 {
		modelVar.Sku = sku
		modelVar.ProductID = productID
		modelVar.Price = req.Price
		modelVar.DiscountAmount = req.DiscountAmount
		modelVar.DiscountType = req.DiscountType
		modelVar.StockLevel = req.StockLevel

		service.DB.Create(&modelVar)
		detailService := prodvardetsvc.NewServiceProductVariationDetail(service.DB)
		detailService.Create(uint64(modelVar.ID), req.AttributeValueIDs)
	}
}

func (service *Service) CreateSimpleWithCSV(modelVar *models.ProductVariations, modelCsv *models.CSVs, productID uint64) {
	images := strings.Split(modelCsv.Images, ",")
	imageUrls, _ := json.Marshal(images)
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
	modelVar.DiscountType = utils.FixedAmountOff
	modelVar.Description = modelCsv.Description
	modelVar.ImageUrls = string(imageUrls)
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
