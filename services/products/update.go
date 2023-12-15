package product

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"encoding/json"
	"strconv"
	"strings"
)

func (service *Service) Update(productID uint64, requestProduct *requests.RequestProduct) error {
	modelProduct := models.Products{}
	if err := service.DB.First(&modelProduct, productID).Error; err != nil {
		return err
	}
	modelProduct.CompanyID = requestProduct.CompanyID
	modelProduct.UserID = requestProduct.UserID
	modelProduct.Name = requestProduct.Name
	imageUrls, _ := json.Marshal(requestProduct.Images)
	modelProduct.ImageUrls = string(imageUrls)
	modelProduct.Brief = requestProduct.Brief
	modelProduct.Description = requestProduct.Description
	modelProduct.SKU = requestProduct.SKU
	modelProduct.UnitPriceRegular = requestProduct.UnitPriceRegular
	modelProduct.UnitPriceSale = requestProduct.UnitPriceSale
	modelProduct.StockQuantity = requestProduct.StockQuantity
	attributes, _ := json.Marshal(requestProduct.Attributes)
	modelProduct.Attributes = string(attributes)
	modelProduct.Tags = requestProduct.Tags
	modelProduct.Active = requestProduct.Active
	// ShippingData
	service.DB.Delete(models.ShippingData{}, modelProduct.ShippingDataID)
	modelShipping := models.ShippingData{
		Weight:         requestProduct.ShippingWeight,
		Dimension:      requestProduct.ShippingDimension,
		Classification: requestProduct.ShippingClassification,
	}
	service.DB.Create(&modelShipping)
	modelProduct.ShippingDataID = uint64(modelShipping.ID)
	// Linked Products
	linkedProductIDs := make([]string, 0)
	linkedProducts := strings.Split(requestProduct.LinkedProducts, ",")
	for _, linkedProduct := range linkedProducts {
		modelLinkedProduct := models.Products{}
		if err := service.DB.Where("name = ?", linkedProduct).First(&modelLinkedProduct).Error; err == nil {
			linkedProductIDs = append(linkedProductIDs, strconv.FormatUint(uint64(modelLinkedProduct.ID), 10))
		}
	}
	modelProduct.LinkedProductIDs = strings.Join(linkedProductIDs, ",")
	service.DB.Save(&modelProduct)
	// Product Channels
	service.DB.Where("store_product_id = ?", productID).Delete(models.ProductChannels{})
	channels := strings.Split(requestProduct.Channels, ",")
	for _, channel := range channels {
		modelChannel := models.ProductChannels{}
		if err := service.DB.Table("channels").Select(" id As channel_id, ? As store_product_id", productID).
			Where("name = ?", channel).Limit(1).Scan(&modelChannel).Error; err == nil {
			service.DB.Create(&modelChannel)
		}
	}
	service.DB.Where("store_product_id = ?", productID).Delete(models.ProductContents{})
	contents := strings.Split(requestProduct.Contents, ",")
	for _, content := range contents {
		modelContent := models.ProductContents{}
		if err := service.DB.Table("contents").Select("id As content_id, ? As store_product_id", productID).
			Where("title = ?", content).Limit(1).Scan(&modelContent).Error; err == nil {
			service.DB.Create(&modelContent)
		}
	}
	return nil
}
