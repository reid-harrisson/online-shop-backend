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

func (repository *RepositoryProduct) ReadByID(modelProduct *models.Products, productID uint64) {
	repository.DB.First(modelProduct, productID)
}

func (repository *RepositoryProduct) ReadLinkedProducts(modelProducts *[]models.ProductsWithLink, productID uint64) {
	repository.DB.Table("store_products As prods").
		Select("prods.*, links.is_up_cross As is_up_cross").
		Joins("Join store_product_links As links On links.link_id = prods.id").
		Where("links.product_id = ?", productID).
		Where("links.deleted_at Is Null And prods.deleted_at Is Null").
		Scan(modelProducts)
}

func (repository *RepositoryProduct) ReadDetail(modelDetail *models.ProductsWithDetail, productID uint64) {
	repository.ReadByID(&modelDetail.Products, productID)

	cateRepo := NewRepositoryCategory(repository.DB)
	cateRepo.ReadByProductID(&modelDetail.Categories, productID)

	attrRepo := NewRepositoryAttribute(repository.DB)
	attrRepo.ReadByProductID(&modelDetail.Attributes, productID)

	tagRepo := NewRepositoryTag(repository.DB)
	tagRepo.ReadByProductID(&modelDetail.Tags, productID)

	chanRepo := NewRepositoryProductChannel(repository.DB)
	chanRepo.ReadByProductID(&modelDetail.RelatedChannels, productID)

	contRepo := NewRepositoryProductContent(repository.DB)
	contRepo.ReadByProductID(&modelDetail.RelatedContents, productID)

	attrValRepo := NewRepositoryProductAttributeValue(repository.DB)
	attrValRepo.ReadByProductID(&modelDetail.AttributeValues, productID)
}

func (repository *RepositoryProduct) ReadPaging(modelProducts *[]models.Products, page uint64, count uint64, storeID uint64, keyword string, totalCount *uint64) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.Model(models.Products{}).
		Where("? = 0 Or store_id = ?", storeID, storeID).
		Where("Lower(title) Like ?", keyword).
		Count(totalCount).Offset(page).Limit(count).Find(modelProducts).Error
}

func (repository *RepositoryProduct) ReadAll(modelProducts *[]models.Products, storeID uint64, keyword string) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.
		Where("? = 0 Or store_id = ?", storeID, storeID).
		Where("Lower(title) Like ?", keyword).
		Find(modelProducts).Error
}

func (repository *RepositoryProduct) ReadCurrencyID(modelProduct *models.Products, storeID uint64) error {
	return repository.DB.
		Model(models.Stores{}).
		Select("cu.id As currency_id").
		Joins("LEFT JOIN users AS u ON u.id = owner_id").
		Joins("LEFT JOIN countries AS ca ON ca.id = u.country_id").
		Joins("LEFT JOIN currencies AS cu ON cu.`code` = ca.currency_code ").
		Where("stores.id = ?", storeID).
		Scan(modelProduct).
		Error
}
