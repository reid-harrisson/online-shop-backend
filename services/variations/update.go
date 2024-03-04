package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
)

func (service *Service) Update(modelVar *models.ProductVariations, req *requests.RequestProductVariation) {
	modelValues := make([]models.ProductAttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryProductAttributeValue(service.DB)
	valRepo.ReadByIDs(&modelValues, req.AttributeValueIDs)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	prodRepo.ReadByID(&modelProduct, modelVar.ProductID)

	sku := modelProduct.Title
	title := modelProduct.Title
	for index, modelValue := range modelValues {
		sku += modelValue.AttributeValue
		if index == 0 {
			title += " - "
		} else {
			title += ", "
		}
		modelVar.Title += modelValue.AttributeValue
	}

	modelVar.Title = title
	modelVar.Sku = utils.CleanSpecialLetters(sku)
	modelVar.Price = req.Price
	modelVar.StockLevel = req.StockLevel
	modelVar.DiscountAmount = req.DiscountAmount
	modelVar.DiscountType = req.DiscountType
	modelVar.Description = req.Description

	service.DB.Save(modelVar)

	detService := prodvardetsvc.NewServiceProductVariationDetail(service.DB)
	detService.Update(uint64(modelVar.ID), req.AttributeValueIDs)
}

func (service *Service) UpdateStockLevel(modelVar *models.ProductVariations, stockLevel float64) {
	modelVar.StockLevel = stockLevel
	service.DB.Save(modelVar)
}
