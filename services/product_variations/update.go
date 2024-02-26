package prodvarsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	prodvardetsvc "OnlineStoreBackend/services/product_variation_details"
)

func (service *Service) Update(modelVar *models.ProductVariations, req *requests.RequestProductVariation) {
	modelValues := make([]models.ProductAttributeValuesWithDetail, 0)
	valRepo := repositories.NewRepositoryProductAttributeValue(service.DB)
	valRepo.ReadByIDs(&modelValues, req.AttributeValueIDs)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(service.DB)
	prodRepo.ReadByID(&modelProduct, modelVar.ProductID)

	modelVar.Sku = GenerateSKU(&modelProduct, &modelValues)
	modelVar.Price = req.Price
	modelVar.StockLevel = req.StockLevel
	modelVar.DiscountAmount = req.DiscountAmount
	modelVar.DiscountType = req.DiscountType

	service.DB.Save(modelVar)

	detailService := prodvardetsvc.NewServiceProductVariationDetail(service.DB)
	detailService.Create(uint64(modelVar.ID), req.AttributeValueIDs)
}

func (service *Service) UpdateStockLevel(modelVar *models.ProductVariations, stockLevel float64) {
	modelVar.StockLevel = stockLevel
	service.DB.Save(modelVar)
}
