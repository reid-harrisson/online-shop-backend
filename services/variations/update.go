package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
)

func (service *Service) Update(modelVar *models.Variations, req *requests.RequestVariation) error {
	modelValues := make([]models.AttributeValuesWithDetail, 0)

	valRepo := repositories.NewRepositoryAttributeValue(service.DB)
	if err := valRepo.ReadByIDs(&modelValues, req.AttributeValueIDs); err != nil {
		return err
	}

	modelProduct := models.Products{}

	prodRepo := repositories.NewRepositoryProduct(service.DB)
	if err := prodRepo.ReadByID(&modelProduct, modelVar.ProductID); err != nil {
		return err
	}

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

	modelVar.Title = title
	modelVar.Sku = utils.CleanSpecialLetters(sku)
	modelVar.Price = req.Price
	modelVar.StockLevel = req.StockLevel
	modelVar.DiscountAmount = req.DiscountAmount
	modelVar.DiscountType = req.DiscountType
	modelVar.Description = req.Description
	modelVar.BackOrderStatus = utils.SimpleStatuses(req.BackOrderAllowed)

	if err := service.DB.Save(modelVar).Error; err != nil {
		return err
	}

	detService := prodvardetsvc.NewServiceVariationDetail(service.DB)
	if err := detService.Update(uint64(modelVar.ID), req.AttributeValueIDs); err != nil {
		return err
	}

	return nil
}

func (service *Service) UpdateBackOrder(modelVar *models.Variations) error {
	switch modelVar.BackOrderStatus {
	case utils.Disabled:
		modelVar.BackOrderStatus = utils.Enabled
	case utils.Enabled:
		modelVar.BackOrderStatus = utils.Disabled
	}

	return service.DB.Save(modelVar).Error
}

func (service *Service) UpdateStockLevel(modelVar *models.Variations, stockLevel float64) error {
	modelVar.StockLevel = stockLevel

	return service.DB.Save(modelVar).Error
}
