package repositories

import (
	"OnlineStoreBackend/models"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type RepositoryProduct struct {
	DB *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *RepositoryProduct {
	return &RepositoryProduct{DB: db}
}

func (repository *RepositoryProduct) ReadOne(modelProductDetail *models.ProductDetails, productID uint64) error {
	// Product
	modelProduct := models.Products{}
	if err := repository.DB.First(&modelProduct, productID).Error; err != nil {
		return err
	}
	modelProductDetail.ID = modelProduct.ID
	modelProductDetail.CompanyID = modelProduct.CompanyID
	modelProductDetail.UserID = modelProduct.UserID
	modelProductDetail.Name = modelProduct.Name
	modelProductDetail.Brief = modelProduct.Brief
	modelProductDetail.Description = modelProduct.Description
	modelProductDetail.ImageUrls = modelProduct.ImageUrls
	modelProductDetail.SKU = modelProduct.SKU
	modelProductDetail.Tags = modelProduct.Tags
	modelProductDetail.UnitPriceRegular = modelProduct.UnitPriceRegular
	modelProductDetail.UnitPriceSale = modelProduct.UnitPriceSale
	modelProductDetail.StockQuantity = modelProduct.StockQuantity
	modelProductDetail.ShippingDataID = modelProduct.ShippingDataID
	modelProductDetail.LinkedProductIDs = modelProduct.LinkedProductIDs
	modelProductDetail.Attributes = modelProduct.Attributes
	modelProductDetail.Active = modelProduct.Active
	modelProductDetail.Reviews = modelProduct.Reviews

	// Shipping Data
	modelShipping := models.ShippingData{}
	repository.DB.First(&modelShipping, modelProduct.ShippingDataID)
	modelProductDetail.ShippingInfo = modelShipping

	// Channels
	modelChannels := make([]models.ProductChannels, 0)
	repository.DB.Where("store_product_id = ?", productID).Find(&modelChannels)
	repository.DB.Where("")
	channelIDs := make([]string, 0)
	for _, modelChannel := range modelChannels {
		channelIDs = append(channelIDs, strconv.FormatUint(modelChannel.ChannelID, 10))
	}
	modelProductDetail.ChannelIDs = strings.Join(channelIDs, ",")

	// Contents
	modelContents := make([]models.ProductContents, 0)
	repository.DB.Where("store_product_id = ?", productID).Find(&modelContents)
	contentIDs := make([]string, 0)
	for _, modelContent := range modelContents {
		contentIDs = append(contentIDs, strconv.FormatUint(modelContent.ContentID, 10))
	}
	modelProductDetail.ContentIDs = strings.Join(contentIDs, ",")

	// ImageUrls and Attributes
	_ = json.Unmarshal([]byte(modelProductDetail.ImageUrls), &modelProductDetail.ImgUrls)
	_ = json.Unmarshal([]byte(modelProductDetail.Attributes), &modelProductDetail.Attribs)
	return nil
}

func (repository *RepositoryProduct) ReadAll(modelProductDetails *[]models.ProductDetails) error {
	modelProducts := make([]models.Products, 0)
	repository.DB.Find(&modelProducts)
	for _, modelProduct := range modelProducts {
		productID := modelProduct.ID
		modelProductDetail := models.ProductDetails{}
		modelProductDetail.ID = modelProduct.ID
		modelProductDetail.CompanyID = modelProduct.CompanyID
		modelProductDetail.UserID = modelProduct.UserID
		modelProductDetail.Name = modelProduct.Name
		modelProductDetail.Brief = modelProduct.Brief
		modelProductDetail.Description = modelProduct.Description
		modelProductDetail.ImageUrls = modelProduct.ImageUrls
		modelProductDetail.SKU = modelProduct.SKU
		modelProductDetail.Tags = modelProduct.Tags
		modelProductDetail.UnitPriceRegular = modelProduct.UnitPriceRegular
		modelProductDetail.UnitPriceSale = modelProduct.UnitPriceSale
		modelProductDetail.StockQuantity = modelProduct.StockQuantity
		modelProductDetail.ShippingDataID = modelProduct.ShippingDataID
		modelProductDetail.LinkedProductIDs = modelProduct.LinkedProductIDs
		modelProductDetail.Attributes = modelProduct.Attributes
		modelProductDetail.Active = modelProduct.Active
		modelProductDetail.Reviews = modelProduct.Reviews

		// Shipping Data
		modelShipping := models.ShippingData{}
		repository.DB.First(&modelShipping, modelProduct.ShippingDataID)
		modelProductDetail.ShippingInfo = modelShipping

		// Channels
		modelChannels := make([]models.ProductChannels, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelChannels)
		repository.DB.Where("")
		channelIDs := make([]string, 0)
		for _, modelChannel := range modelChannels {
			channelIDs = append(channelIDs, strconv.FormatUint(modelChannel.ChannelID, 10))
		}
		modelProductDetail.ChannelIDs = strings.Join(channelIDs, ",")

		// Contents
		modelContents := make([]models.ProductContents, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelContents)
		contentIDs := make([]string, 0)
		for _, modelContent := range modelContents {
			contentIDs = append(contentIDs, strconv.FormatUint(modelContent.ContentID, 10))
		}
		modelProductDetail.ContentIDs = strings.Join(contentIDs, ",")

		// ImageUrls and Attributes
		_ = json.Unmarshal([]byte(modelProductDetail.ImageUrls), &modelProductDetail.ImgUrls)
		_ = json.Unmarshal([]byte(modelProductDetail.Attributes), &modelProductDetail.Attribs)
		*modelProductDetails = append(*modelProductDetails, modelProductDetail)
	}
	return nil
}

func (repository *RepositoryProduct) ReadActive(modelProductDetails *[]models.ProductDetails) error {
	modelProducts := make([]models.Products, 0)
	repository.DB.Where("active = 1").Find(&modelProducts)
	for _, modelProduct := range modelProducts {
		productID := modelProduct.ID
		modelProductDetail := models.ProductDetails{}
		modelProductDetail.ID = modelProduct.ID
		modelProductDetail.CompanyID = modelProduct.CompanyID
		modelProductDetail.UserID = modelProduct.UserID
		modelProductDetail.Name = modelProduct.Name
		modelProductDetail.Brief = modelProduct.Brief
		modelProductDetail.Description = modelProduct.Description
		modelProductDetail.ImageUrls = modelProduct.ImageUrls
		modelProductDetail.SKU = modelProduct.SKU
		modelProductDetail.Tags = modelProduct.Tags
		modelProductDetail.UnitPriceRegular = modelProduct.UnitPriceRegular
		modelProductDetail.UnitPriceSale = modelProduct.UnitPriceSale
		modelProductDetail.StockQuantity = modelProduct.StockQuantity
		modelProductDetail.ShippingDataID = modelProduct.ShippingDataID
		modelProductDetail.LinkedProductIDs = modelProduct.LinkedProductIDs
		modelProductDetail.Attributes = modelProduct.Attributes
		modelProductDetail.Active = modelProduct.Active
		modelProductDetail.Reviews = modelProduct.Reviews

		// Shipping Data
		modelShipping := models.ShippingData{}
		repository.DB.First(&modelShipping, modelProduct.ShippingDataID)
		modelProductDetail.ShippingInfo = modelShipping

		// Channels
		modelChannels := make([]models.ProductChannels, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelChannels)
		repository.DB.Where("")
		channelIDs := make([]string, 0)
		for _, modelChannel := range modelChannels {
			channelIDs = append(channelIDs, strconv.FormatUint(modelChannel.ChannelID, 10))
		}
		modelProductDetail.ChannelIDs = strings.Join(channelIDs, ",")

		// Contents
		modelContents := make([]models.ProductContents, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelContents)
		contentIDs := make([]string, 0)
		for _, modelContent := range modelContents {
			contentIDs = append(contentIDs, strconv.FormatUint(modelContent.ContentID, 10))
		}
		modelProductDetail.ContentIDs = strings.Join(contentIDs, ",")

		// ImageUrls and Attributes
		_ = json.Unmarshal([]byte(modelProductDetail.ImageUrls), &modelProductDetail.ImgUrls)
		_ = json.Unmarshal([]byte(modelProductDetail.Attributes), &modelProductDetail.Attribs)
		*modelProductDetails = append(*modelProductDetails, modelProductDetail)
	}
	return nil
}

func (repository *RepositoryProduct) ReadPaging(modelProductDetails *[]models.ProductDetails, page uint64, count uint64, companyID uint64, userID uint64, keyword string, totalCount *uint64) error {
	keyword = strings.ToLower("%" + keyword + "%")
	modelProducts := make([]models.Products, 0)
	repository.DB.Where("? = 0 Or company_id = ?", companyID, companyID).
		Where("? = 0 Or user_id = ?", userID, userID).
		Where("Lower(name) Like ? Or Lower(brief) Like ? Or Lower(description) Like ?", keyword, keyword, keyword).
		Count(totalCount).Offset(page).Limit(count).Find(&modelProducts)

	for _, modelProduct := range modelProducts {
		productID := modelProduct.ID
		modelProductDetail := models.ProductDetails{}
		modelProductDetail.ID = modelProduct.ID
		modelProductDetail.CompanyID = modelProduct.CompanyID
		modelProductDetail.UserID = modelProduct.UserID
		modelProductDetail.Name = modelProduct.Name
		modelProductDetail.Brief = modelProduct.Brief
		modelProductDetail.Description = modelProduct.Description
		modelProductDetail.ImageUrls = modelProduct.ImageUrls
		modelProductDetail.SKU = modelProduct.SKU
		modelProductDetail.Tags = modelProduct.Tags
		modelProductDetail.UnitPriceRegular = modelProduct.UnitPriceRegular
		modelProductDetail.UnitPriceSale = modelProduct.UnitPriceSale
		modelProductDetail.StockQuantity = modelProduct.StockQuantity
		modelProductDetail.ShippingDataID = modelProduct.ShippingDataID
		modelProductDetail.LinkedProductIDs = modelProduct.LinkedProductIDs
		modelProductDetail.Attributes = modelProduct.Attributes
		modelProductDetail.Active = modelProduct.Active
		modelProductDetail.Reviews = modelProduct.Reviews

		// Shipping Data
		modelShipping := models.ShippingData{}
		repository.DB.First(&modelShipping, modelProduct.ShippingDataID)
		modelProductDetail.ShippingInfo = modelShipping

		// Channels
		modelChannels := make([]models.ProductChannels, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelChannels)
		repository.DB.Where("")
		channelIDs := make([]string, 0)
		for _, modelChannel := range modelChannels {
			channelIDs = append(channelIDs, strconv.FormatUint(modelChannel.ChannelID, 10))
		}
		modelProductDetail.ChannelIDs = strings.Join(channelIDs, ",")

		// Contents
		modelContents := make([]models.ProductContents, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelContents)
		contentIDs := make([]string, 0)
		for _, modelContent := range modelContents {
			contentIDs = append(contentIDs, strconv.FormatUint(modelContent.ContentID, 10))
		}
		modelProductDetail.ContentIDs = strings.Join(contentIDs, ",")

		// ImageUrls and Attributes
		_ = json.Unmarshal([]byte(modelProductDetail.ImageUrls), &modelProductDetail.ImgUrls)
		_ = json.Unmarshal([]byte(modelProductDetail.Attributes), &modelProductDetail.Attribs)
		*modelProductDetails = append(*modelProductDetails, modelProductDetail)
	}
	return nil
}

func (repository *RepositoryProduct) ReadSearch(modelProductDetails *[]models.ProductDetails, franchiseeID uint64, keyword string) error {
	keyword = strings.ToLower("%" + keyword + "%")
	modelProducts := make([]models.Products, 0)
	repository.DB.Where("active = 1").
		Where("? = 0 Or company_id = ?", franchiseeID, franchiseeID).
		Where("Lower(name) Like ? Or Lower(brief) Like ? Or Lower(description) Like ? Or Lower(tags) Like ?", keyword, keyword, keyword, keyword).
		Find(&modelProducts)

	for _, modelProduct := range modelProducts {
		productID := modelProduct.ID
		modelProductDetail := models.ProductDetails{}
		modelProductDetail.ID = modelProduct.ID
		modelProductDetail.CompanyID = modelProduct.CompanyID
		modelProductDetail.UserID = modelProduct.UserID
		modelProductDetail.Name = modelProduct.Name
		modelProductDetail.Brief = modelProduct.Brief
		modelProductDetail.Description = modelProduct.Description
		modelProductDetail.ImageUrls = modelProduct.ImageUrls
		modelProductDetail.SKU = modelProduct.SKU
		modelProductDetail.Tags = modelProduct.Tags
		modelProductDetail.UnitPriceRegular = modelProduct.UnitPriceRegular
		modelProductDetail.UnitPriceSale = modelProduct.UnitPriceSale
		modelProductDetail.StockQuantity = modelProduct.StockQuantity
		modelProductDetail.ShippingDataID = modelProduct.ShippingDataID
		modelProductDetail.LinkedProductIDs = modelProduct.LinkedProductIDs
		modelProductDetail.Attributes = modelProduct.Attributes
		modelProductDetail.Active = modelProduct.Active
		modelProductDetail.Reviews = modelProduct.Reviews

		// Shipping Data
		modelShipping := models.ShippingData{}
		repository.DB.First(&modelShipping, modelProduct.ShippingDataID)
		modelProductDetail.ShippingInfo = modelShipping

		// Channels
		modelChannels := make([]models.ProductChannels, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelChannels)
		repository.DB.Where("")
		channelIDs := make([]string, 0)
		for _, modelChannel := range modelChannels {
			channelIDs = append(channelIDs, strconv.FormatUint(modelChannel.ChannelID, 10))
		}
		modelProductDetail.ChannelIDs = strings.Join(channelIDs, ",")

		// Contents
		modelContents := make([]models.ProductContents, 0)
		repository.DB.Where("store_product_id = ?", productID).Find(&modelContents)
		contentIDs := make([]string, 0)
		for _, modelContent := range modelContents {
			contentIDs = append(contentIDs, strconv.FormatUint(modelContent.ContentID, 10))
		}
		modelProductDetail.ContentIDs = strings.Join(contentIDs, ",")

		// ImageUrls and Attributes
		_ = json.Unmarshal([]byte(modelProductDetail.ImageUrls), &modelProductDetail.ImgUrls)
		_ = json.Unmarshal([]byte(modelProductDetail.Attributes), &modelProductDetail.Attribs)
		*modelProductDetails = append(*modelProductDetails, modelProductDetail)
	}
	return nil
}
