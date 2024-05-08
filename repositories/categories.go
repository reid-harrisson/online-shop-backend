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

func (repository *RepositoryCategory) ReadByNameAndStoreIDAndParentID(modelCategory *models.Categories, name string, storeID uint64, parentID uint64) error {
	return repository.DB.Where("name = ? And store_id = ? And parent_id = ?", name, storeID, parentID).First(modelCategory).Error
}

func (repository *RepositoryCategory) ReadByID(modelCategory *models.Categories, categoryID uint64) error {
	return repository.DB.Where("id = ?", categoryID).First(&modelCategory).Error
}

func (repository *RepositoryCategory) ReadByParentID(modelCategory *[]models.Categories, parentID uint64) error {
	return repository.DB.Where("parent_id = ?", parentID).Find(&modelCategory).Error
}

func (repository *RepositoryCategory) ReadByStoreID(modelStoreCategories *[]models.CategoriesWithChildren, storeID uint64) error {
	modelCategories := make([]models.Categories, 0)
	err := repository.DB.Where("store_id = ?", storeID).Find(&modelCategories).Error
	if err != nil {
		return err
	}

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

	return nil
}
