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

func (repository *RepositoryProduct) ReadByID(modelProduct *models.Products, productID uint64) error {
	return repository.DB.First(&modelProduct, productID).Error
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

	shipRepo := NewRepositoryShippingData(repository.DB)
	shipRepo.ReadByProductID(&modelDetail.ShippingData, productID)
}

func (repository *RepositoryProduct) ReadPaging(modelProducts *[]models.Products, page uint64, count uint64, storeID uint64, keyword string, totalCount *uint64) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.Model(models.Products{}).
		Where("? = 0 Or store_id = ?", storeID, storeID).
		Where("Lower(title) Like ?", keyword).
		Count(totalCount).Offset(page).Limit(count).Find(&modelProducts).Error
}

func (repository *RepositoryProduct) ReadAll(modelProducts *[]models.Products, storeID uint64, keyword string) error {
	keyword = strings.ToLower("%" + keyword + "%")
	return repository.DB.
		Where("? = 0 Or store_id = ?", storeID, storeID).
		Where("Lower(title) Like ?", keyword).
		Find(&modelProducts).Error
}
