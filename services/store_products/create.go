package prod

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	"encoding/json"
)

func (service *Service) Create(modelProduct *models.Products, req *requests.RequestProduct) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Name = req.Name
	modelProduct.Brief = req.Brief
	modelProduct.Description = req.Description
	modelProduct.SKU = utils.GenerateSKU(req.Name, req.StoreID)
	modelProduct.UnitPriceRegular = req.UnitPriceRegular
	modelProduct.UnitPriceSale = req.UnitPriceSale
	modelProduct.StockQuantity = req.StockQuantity
	modelProduct.Active = req.Active
	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	return service.DB.Create(modelProduct).Error
}
