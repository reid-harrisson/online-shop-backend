package prodsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"encoding/json"
)

func (service *Service) Update(modelProduct *models.Products, req *requests.RequestProduct) error {
	productID := uint64(modelProduct.ID)

	if err := service.DB.First(&modelProduct, productID).Error; err != nil {
		return err
	}

	modelProduct.StoreID = req.StoreID
	modelProduct.Title = req.Title
	modelProduct.ShortDescription = req.ShortDescription
	modelProduct.LongDescription = req.LongDescirpiton
	modelProduct.Status = utils.Draft

	modelCurrenyID := models.ProductCurrencyID{}

	productRepository := repositories.NewRepositoryProduct(service.DB)
	productRepository.ReadCurrencyID(&modelCurrenyID, req.StoreID)

	modelProduct.CurrencyID = modelCurrenyID.CurrencyID
	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)

	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateMinimumStockLevel(productID uint64, req *requests.RequestMinimumStockLevel, modelProduct *models.Products) error {
	modelProduct.MinimumStockLevel = req.Level
	return service.DB.Save(modelProduct).Error
}
