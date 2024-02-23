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
	modelProduct.Status = utils.Draft

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)

	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateMinimumStockLevel(productID uint64, req *requests.RequestMinimumStockLevel, modelProduct *models.Products) error {
	modelProduct.MinimumStockLevel = req.Level
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
	return service.DB.Model(models.Products{}).Where("id = ?", productID).Update("status", status).Error
}
