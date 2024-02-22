package varsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	detailsvc "OnlineStoreBackend/services/product_variation_details"
)

func (service *Service) Create(modelVar *models.StoreProductVariations, req *requests.RequestStoreProductVariation, productID uint64) {
	price := req.Price
	switch req.Type {
	case utils.PercentageOff:
		price = price - price*req.Discount/100
	case utils.FixedAmountOff:
		price = price - req.Discount
	}

	modelVar.Sku = ""
	modelVar.ProductID = productID
	modelVar.Price = price
	modelVar.StockLevel = req.StockLevel

	service.DB.Create(&modelVar)
	detailService := detailsvc.NewServiceProductVariationDetail(service.DB)
	detailService.Create(uint64(modelVar.ID), req.AttributeValueIDs)
}
