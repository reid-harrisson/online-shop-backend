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
	modelProduct.LongDescription = req.LongDescirpiton

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)

	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateMinimumStockLevel(productID uint64, minimumStockLevel float64, modelProduct *models.Products) error {
	modelProduct.MinimumStockLevel = minimumStockLevel
	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateStatus(modelProduct *models.Products, status utils.ProductStatus) error {
	modelProduct.Status = status
	return service.DB.Save(modelProduct).Error
}
