package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	"encoding/json"
	"strings"
)

func (service *Service) Update(modelProduct *models.Products, req *requests.RequestProduct) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton
	modelProduct.SKU = utils.GenerateSKU(req.Title, req.StoreID)
	modelProduct.UnitPriceRegular = req.UnitPriceRegular
	modelProduct.UnitPriceSale = req.UnitPriceRegular
	modelProduct.StockQuantity = req.StockQuantity
	modelProduct.Active = req.Active
	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateMinimumStockLevel(productID uint64, req *requests.RequestMinimumStockLevel, modelProduct *models.Products) error {
	modelProduct.StockQuantity = req.Level
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

func (service *Service) UpdateStatus(productID uint64, productStatus string) error {
	status := models.StatusProductPending
	switch strings.ToLower(productStatus) {
	case "pending":
		status = models.StatusProductPending
	case "approved":
		status = models.StatusProductApproved
	}
	return service.DB.Model(models.Products{}).Where("id = ?", productID).Update("product_status", status).Error
}
