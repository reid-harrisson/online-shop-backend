package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	"encoding/json"
)

func (service *Service) Update(modelProduct *models.Products, req *requests.RequestProduct) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescription

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)

	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateMinimumStockLevel(productID uint64, minimumStockLevel float64) error {
	return service.DB.Model(&models.Products{}).Where("id = ?", productID).Update("minimum_stock_level", minimumStockLevel).Error
}

func (service *Service) UpdateStatus(productID uint64, status utils.ProductStatus) error {
	return service.DB.Model(&models.Products{}).Where("id = ?", productID).Update("status", status).Error
}
