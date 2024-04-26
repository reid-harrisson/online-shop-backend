package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryCategory struct {
	DB *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *RepositoryCategory {
	return &RepositoryCategory{DB: db}
}

func (repository *RepositoryCategory) ReadByProductID(modelCategories *[]models.ProductCategoriesWithName, productID uint64) error {
	return repository.DB.Table("store_product_categories As prodcates").
		Select("prodcates.*, cates.name As category_name").
		Joins("Join store_categories As cates On cates.id = prodcates.category_id").
		Where("cates.deleted_at Is Null And prodcates.deleted_at Is Null").
		Where("prodcates.product_id = ?", productID).
		Scan(modelCategories).
		Error
}

func (repository *RepositoryCategory) ReadByName(modelCategory *models.Categories, name string, storeID uint64) {
	repository.DB.Where("name = ? And store_id = ?", name, storeID).First(modelCategory)
}

func (repository *RepositoryCategory) ReadByID(modelCategory *models.Categories, categoryID uint64) error {
	return repository.DB.Where("id = ?", categoryID).Error
}

func (repository *RepositoryCategory) ReadByStoreID(modelStoreCategories *[]models.CategoriesWithChildren, storeID uint64) {
	modelCategories := make([]models.Categories, 0)
	repository.DB.Where("store_id = ?", storeID).Find(&modelCategories)
	allChildrenIDs := make(map[uint64][]uint64)
	for _, modelCategory := range modelCategories {
		parentID := modelCategory.ParentID
		allChildrenIDs[parentID] = append(allChildrenIDs[parentID], uint64(modelCategory.ID))
	}
	for _, modelCategory := range modelCategories {
		childrenIDs := make([]uint64, 0)
		childrenIDs = append(childrenIDs, allChildrenIDs[uint64(modelCategory.ID)]...)
		*modelStoreCategories = append(*modelStoreCategories, models.CategoriesWithChildren{
			Categories:  modelCategory,
			ChildrenIDs: childrenIDs,
		})
	}
}
