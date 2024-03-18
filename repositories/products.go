package repositories

import (
	"OnlineStoreBackend/models"
	"strings"

	"gorm.io/gorm"
)

type RepositoryProduct struct {
	DB *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *RepositoryProduct {
	return &RepositoryProduct{DB: db}
}

func (repository *RepositoryProduct) ReadByID(modelProduct *models.Products, productID uint64) error {
	return repository.DB.First(modelProduct, productID).Error
}

func (repository *RepositoryProduct) ReadLinkedProducts(modelProducts *[]models.ProductsWithLink, productID uint64) error {
	return repository.DB.Table("store_products As prods").
		Select("prods.*, links.is_up_cross As is_up_cross").
		Joins("Join store_product_links As links On links.link_id = prods.id").
		Where("links.product_id = ?", productID).
		Where("links.deleted_at Is Null And prods.deleted_at Is Null").
		Scan(modelProducts).Error
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

func (repository *RepositoryProduct) ReadPaging(modelProducts *[]models.Products, page int, count int, storeID uint64, keyword string, totalCount *int64) error {
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

func (repository *RepositoryProduct) ReadByCategory(modelProducts *[]models.Products, storeID uint64, cateID uint64) error {
	return repository.DB.Table("store_product_categories As cates").
		Select("prods.*").
		Joins("Left Join store_products As prods On prods.id = cates.product_id").
		Where("cates.category_id = ?", cateID).
		Where("? = 0 Or prods.store_id = ?", storeID, storeID).
		Group("prods.id").
		Scan(modelProducts).Error
}

func (repository *RepositoryProduct) ReadByTags(modelProducts *[]models.Products, storeID uint64, tags []string, keyword string) {
	modelTags := []models.StoreTags{}
	repository.DB.Where("name Not In (?)", tags).Find(&modelTags)
	keyword = "%" + keyword + "%"
	tagIDs := []uint64{}
	for _, modelTag := range modelTags {
		tagIDs = append(tagIDs, uint64(modelTag.ID))
	}
	repository.DB.Table("store_product_tags As tags").
		Select("prods.*").
		Joins("Left Join store_products As prods On prods.id = tags.product_id").
		Where("? = 0 Or prods.store_id = ?", storeID, storeID).
		Where("? = '%%' Or prods.title Like ?", keyword, keyword).
		Where("tags.tag_id Not In (?)", tagIDs).
		Group("prods.id").
		Having("Count(tags.tag_id) = ?", len(tags)).
		Scan(modelProducts)
}
