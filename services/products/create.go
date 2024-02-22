package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"encoding/json"
)

func (service *Service) Create(modelProduct *models.Products, req *requests.RequestProduct) error {
	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton
	modelProduct.Status = int8(models.Draft)

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	return service.DB.Create(modelProduct).Error
}
