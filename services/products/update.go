package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	"encoding/json"
)

func (service *Service) Update(modelProduct *models.Products, req *requests.RequestProduct) error {
	productID := uint64(modelProduct.ID)
	if err := service.DB.First(&modelProduct, productID).Error; err != nil {
		return err
	}
	modelProduct.StoreID = req.StoreID
	modelProduct.Name = req.Name
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton
	modelProduct.SKU = utils.GenerateSKU(req.Name, req.StoreID)
	modelProduct.UnitPriceRegular = req.UnitPriceRegular
	modelProduct.StockQuantity = req.StockQuantity
	modelProduct.Active = req.Active
	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateStockQuantity(productID uint64, req *requests.RequestProductQuantity, modelProduct *models.Products) error {
	if err := service.DB.First(modelProduct, productID).Error; err != nil {
		return err
	}
	modelProduct.StockQuantity = req.Quantity
	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdatePrice(productID uint64, req *requests.RequestProductPrice, modelProduct *models.Products) error {
	if err := service.DB.First(modelProduct, productID).Error; err != nil {
		return err
	}
	modelProduct.UnitPriceSale = req.Price * modelProduct.UnitPriceSale / modelProduct.UnitPriceRegular
	modelProduct.UnitPriceRegular = req.Price
	return service.DB.Save(modelProduct).Error
}
