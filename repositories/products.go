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

func (repository *RepositoryProduct) ReadOne(modelProduct *models.ProductsWithCategory, productID uint64) {
	repository.DB.Table("store_products As prods").Select("prods.*, cates.name As category").
		Joins("Left Join store_categories As cates On cates.id = prods.category_id").
		Where("prods.id = ?", productID).
		Where("prods.deleted_at Is Null And cates.deleted_at Is Null").
		Limit(1).
		Scan(modelProduct)
}

func (repository *RepositoryProduct) ReadDetail(modelDetail *models.ProductsWithDetail, productID uint64) {
	repository.ReadOne(&modelDetail.ProductsWithCategory, productID)

	attrRepo := NewRepositoryAttribute(repository.DB)
	attrRepo.ReadByProductID(&modelDetail.Attributes, productID)

	tagRepo := NewRepositoryTag(repository.DB)
	tagRepo.ReadByProductID(&modelDetail.Tags, productID)

	chanRepo := NewRepositoryProductChannel(repository.DB)
	chanRepo.ReadByProductID(&modelDetail.RelatedChannels, productID)

	contRepo := NewRepositoryProductContent(repository.DB)
	contRepo.ReadByProductID(&modelDetail.RelatedContents, productID)

	shipRepo := NewRepositoryShipping(repository.DB)
	shipRepo.ReadByProductID(&modelDetail.ShippingData, productID)

	varRepo := NewRepositoryProductVariation(repository.DB)
	varRepo.ReadByProductID(&modelDetail.Variations, productID)
}

func (repository *RepositoryProduct) ReadPaging(modelProducts *[]models.ProductsWithCategory, page uint64, count uint64, storeID uint64, keyword string, totalCount *uint64) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.Table("store_products As prods").Select("prods.*, cates.name As category").
		Joins("Left Join store_categories As cates On cates.id = prods.category_id").
		Where("? = 0 Or prods.store_id = ?", storeID, storeID).
		Where("Lower(title) Like ?", keyword).
		Where("prods.deleted_at Is Null And cates.deleted_at Is Null").
		Count(totalCount).Offset(page).Limit(count).Find(&modelProducts).Error
}

func (repository *RepositoryProduct) ReadAll(modelProducts *[]models.ProductsWithCategory, storeID uint64, keyword string) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.Table("store_products As prods").Select("prods.*, cates.name As category").
		Joins("Left Join store_categories As cates On cates.id = prods.category_id").
		Where("? = 0 Or prods.store_id = ?", storeID, storeID).
		Where("Lower(title) Like ?", keyword).
		Where("prods.deleted_at Is Null And cates.deleted_at Is Null").
		Scan(&modelProducts).Error
}
