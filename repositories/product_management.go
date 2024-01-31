package repositories

import (
	"OnlineStoreBackend/models"
	"strings"

	"github.com/jinzhu/gorm"
)

type RepositoryProduct struct {
	DB *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *RepositoryProduct {
	return &RepositoryProduct{DB: db}
}

func (repository *RepositoryProduct) Read(modelProduct *models.Products, productID uint64) error {
	return repository.DB.First(&modelProduct, productID).Error
}

func (repository *RepositoryProduct) ReadDetail(modelProdDetail *models.ProductDetails, productID uint64) error {
	modelProduct := models.Products{}
	if err := repository.DB.First(&modelProduct, productID).Error; err != nil {
		return err
	}
	modelProdDetail.Products = modelProduct
	attributes := make(map[string]string)
	tags := make([]string, 0)
	channels := make([]string, 0)
	contents := make([]string, 0)
	modelAttrs := make([]models.Attributes, 0)
	modelTags := make([]models.Tags, 0)
	modelChannels := make([]models.ProductChannelWithName, 0)
	modelContents := make([]models.ProductContentWithName, 0)
	modelShipData := models.ShippingData{}
	modelReviews := make([]models.ProductReviews, 0)
	query := repository.DB.Where("store_product_id = ?", productID)

	// attributes and tags
	query.Find(&modelAttrs)
	for _, modelAttr := range modelAttrs {
		attributes[modelAttr.Attribute] = modelAttr.Value
	}
	query.Find(&modelTags)
	for _, modelTag := range modelTags {
		tags = append(tags, modelTag.Tag)
	}
	modelProdDetail.Attributes = attributes
	modelProdDetail.Tags = tags

	// channels and content
	query.Table("store_product_related_channels As prodChans").
		Select("prodChans.id As id, prodChans.store_product_id As store_product_id, channels.name As channel_name").
		Joins("Left Join channels On channels.id = prodChans.channel_id").
		Where("prodChans.deleted_at Is Null And channels.deleted_at Is Null").
		Scan(&modelChannels)
	query.Table("store_product_related_contents As prodConts").
		Select("prodConts.id As id, prodConts.store_product_id As store_product_id, contents.title As content_name").
		Joins("Left Join contents On contents.id = prodConts.content_id").
		Where("prodConts.deleted_at Is Null And contents.deleted_at Is Null").
		Scan(&modelContents)
	for _, modelChannel := range modelChannels {
		channels = append(channels, modelChannel.ChannelName)
	}
	for _, modelContent := range modelContents {
		contents = append(contents, modelContent.ContentName)
	}
	query.First(&modelShipData)
	query.Find(&modelReviews)
	modelProdDetail.RelatedChannels = channels
	modelProdDetail.RelatedContents = contents
	modelProdDetail.ShipData = modelShipData
	modelProdDetail.Reviews = modelReviews
	return nil
}

func (repository *RepositoryProduct) ReadPaging(modelProducts *[]models.Products, page uint64, count uint64, storeID uint64, keyword string, totalCount *uint64) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.Model(modelProducts).Where("? = 0 Or store_id = ?", storeID, storeID).
		Where("Lower(name) Like ? Or Lower(brief) Like ? Or Lower(description) Like ?", keyword, keyword, keyword).
		Count(totalCount).Offset(page).Limit(count).Find(&modelProducts).Error
}

func (repository *RepositoryProduct) ReadAll(modelProducts *[]models.Products, storeID uint64, keyword string) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.Where("? = 0 Or store_id = ?", storeID, storeID).
		Where("Lower(name) Like ? Or Lower(brief) Like ? Or Lower(description) Like ?", keyword, keyword, keyword).
		Find(&modelProducts).Error
}
