package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
)

func (service *Service) Update(modelVar *models.Variations, req *requests.RequestVariation) {
	modelValues := make([]models.AttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryAttributeValue(service.DB)
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
	modelVar.BackOrderStatus = utils.SimpleStatuses(req.BackOrderAllowed)

	service.DB.Save(modelVar)

	detService := prodvardetsvc.NewServiceVariationDetail(service.DB)
	detService.Update(uint64(modelVar.ID), req.AttributeValueIDs)
}

func (service *Service) UpdateBackOrder(modelVar *models.Variations) {
	switch modelVar.BackOrderStatus {
	case utils.Disabled:
		modelVar.BackOrderStatus = utils.Enabled
	case utils.Enabled:
		modelVar.BackOrderStatus = utils.Disabled
	}
	service.DB.Save(modelVar)
}

func (service *Service) UpdateStockLevel(modelVar *models.Variations, stockLevel float64) {
	modelVar.StockLevel = stockLevel
	service.DB.Save(modelVar)
}
