package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
	"encoding/json"
	"fmt"
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
	varRepo.ReadByValueIDs(modelVar, req.AttributeValueIDs, productID)

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

func (service *Service) CreateWithCSV(modelNewVars *[]models.ProductVariations, varMatches []string, varIndices map[string]int) {
	modelCurVars := []models.ProductVariations{}
	service.DB.Where("Concat(product_id, ':', sku) In (?)", varMatches).Find(&modelCurVars)
	for _, modelVar := range modelCurVars {
		match := fmt.Sprintf("%d:%s", modelVar.ProductID, modelVar.Sku)
		index := varIndices[match]
		(*modelNewVars)[index].ID = modelVar.ID
	}
	service.DB.Save(modelNewVars)
}
