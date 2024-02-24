package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvarsvc "OnlineStoreBackend/services/product_variation_details"
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
	price := req.Price
	switch req.Type {
	case utils.PercentageOff:
		price = price - price*req.Discount/100
	case utils.FixedAmountOff:
		price = price - req.Discount
	}

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
		modelVar.Price = price
		modelVar.StockLevel = req.StockLevel

		service.DB.Create(&modelVar)
		detailService := prodvarsvc.NewServiceProductVariationDetail(service.DB)
		detailService.Create(uint64(modelVar.ID), req.AttributeValueIDs)
	}
}
