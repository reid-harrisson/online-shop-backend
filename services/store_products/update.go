package prod

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
	modelProduct.Brief = req.Brief
	modelProduct.Description = req.Description
	modelProduct.SKU = utils.GenerateSKU(req.Name, req.StoreID)
	modelProduct.UnitPriceRegular = req.UnitPriceRegular
	modelProduct.UnitPriceSale = req.UnitPriceSale
	modelProduct.StockQuantity = req.StockQuantity
	modelProduct.Active = req.Active
	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelProduct.ImageUrls = string(imageUrls)
	return service.DB.Save(modelProduct).Error
}

func (service *Service) UpdateLinkedProduct(productID uint64, req *requests.RequestProductLinked, modelProduct *models.Products) error {
	if err := service.DB.First(modelProduct, productID).Error; err != nil {
		return err
	}
	linkedProdIDs := make([]uint64, 0)
	for _, linkedProdID := range req.LinkedProductIDs {
		linkedProdIDs = append(linkedProdIDs, linkedProdID)
	}
	linkedProdJson, _ := json.Marshal(linkedProdIDs)
	modelProduct.LinkedProductIDs = string(linkedProdJson)
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
