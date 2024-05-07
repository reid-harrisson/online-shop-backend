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

func (repository *RepositoryProduct) ReadDetail(modelDetail *models.ProductsWithDetail, productID uint64) error {
	err := repository.ReadByID(&modelDetail.Products, productID)
	if err != nil {
		return err
	}

	cateRepo := NewRepositoryCategory(repository.DB)
	err = cateRepo.ReadByProductID(&modelDetail.Categories, productID)
	if err != nil {
		return err
	}

	attrRepo := NewRepositoryAttribute(repository.DB)
	err = attrRepo.ReadByProductID(&modelDetail.Attributes, productID)
	if err != nil {
		return err
	}

	tagRepo := NewRepositoryTag(repository.DB)
	err = tagRepo.ReadByProductID(&modelDetail.Tags, productID)
	if err != nil {
		return err
	}

	chanRepo := NewRepositoryProductChannel(repository.DB)
	err = chanRepo.ReadByProductID(&modelDetail.RelatedChannels, productID)
	if err != nil {
		return err
	}

	contRepo := NewRepositoryProductContent(repository.DB)
	err = contRepo.ReadByProductID(&modelDetail.RelatedContents, productID)
	if err != nil {
		return err
	}

	attrValRepo := NewRepositoryAttributeValue(repository.DB)
	err = attrValRepo.ReadByProductID(&modelDetail.AttributeValues, productID)
	if err != nil {
		return err
	}

	return err
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

func (repository *RepositoryProduct) ReadApproved(modelProducts *[]models.ProductsApproved, storeID uint64, page int, count int, totalCount *int64) error {
	query := repository.DB.Table("store_product_variations As vars").
		Select(`
			prods.id,
			prods.title,
			(Select Avg(revs.rate) From store_product_reviews As revs Where revs.product_id = prods.id) As rating,
			Min(Case When vars.discount_type = 0 Then (vars.price - vars.price * vars.discount_amount / 100) Else (vars.price - vars.discount_amount) End) As minimum_price,
			Max(Case When vars.discount_type = 0 Then (vars.price - vars.price * vars.discount_amount / 100) Else (vars.price - vars.discount_amount) End) As maximum_price,
			vars.price As regular_price,
			prods.image_urls
		`).
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Where("prods.store_id = ?", storeID).
		Group("prods.id").
		Count(totalCount)
	if page != 0 || count != 0 {
		return query.Offset(page).
			Limit(count).
			Find(modelProducts).Error
	}
	return query.Find(modelProducts).Error
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

func (repository *RepositoryProduct) ReadByTags(modelProducts *[]models.Products, storeID uint64, tags []string, keyword string) error {
	if len(tags) == 1 && tags[0] == "" {
		tags = []string{}
	}

	modelTags := []models.Tags{}
	repository.DB.Where("name Not In (?)", tags).Find(&modelTags)
	keyword = "%" + keyword + "%"
	tagIDs := []uint64{}
	for _, modelTag := range modelTags {
		tagIDs = append(tagIDs, uint64(modelTag.ID))
	}

	return repository.DB.Table("store_products As prods").
		Select("prods.*").
		Joins("Left Join store_product_tags As tags On prods.id = tags.product_id").
		Where("? = 0 Or prods.store_id = ?", storeID, storeID).
		Where("? = '%%' Or prods.title Like ?", keyword, keyword).
		Where("? = 0 Or tags.tag_id Not In (?)", len(tags), tagIDs).
		Group("prods.id").
		Having("? = 0 Or Count(tags.tag_id) = ?", len(tags), len(tags)).
		Scan(modelProducts).
		Error
}

func (repository *RepositoryProduct) GetMinimumStockLevel(minimumStockLevel *float64, productID uint64) error {
	return repository.DB.Table("store_products").
		Select("minimum_stock_level").
		Where("id = ?", productID).
		Limit(1).
		Scan(minimumStockLevel).
		Error
}
