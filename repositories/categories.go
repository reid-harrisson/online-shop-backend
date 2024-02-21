package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryCategory struct {
	DB *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *RepositoryCategory {
	return &RepositoryCategory{DB: db}
}

func (repository *RepositoryCategory) ReadByName(modelCategory *models.BaseCategories, name string, storeID uint64) {
	repository.DB.Where("name = ? And store_id = ?", name, storeID).First(modelCategory)
}

func (repository *RepositoryCategory) ReadByCategoryID(modelCategory *models.BaseCategories, categoryID uint64) {
	repository.DB.First(modelCategory, categoryID)
}

func (repository *RepositoryCategory) ReadByStoreID(modelStoreCategories *[]models.StoreCategoriesWithChildren, storeID uint64) {
	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	repository.DB.Where("store_id = ?", storeID).Find(&modelCategories)
	allChildrenIDs := make(map[uint64][]uint64)
	for _, modelCategory := range modelCategories {
		parentID := modelCategory.ParentID
		allChildrenIDs[parentID] = append(allChildrenIDs[parentID], uint64(modelCategory.ID))
	}
	for _, modelCategory := range modelCategories {
		modelCategory.ChildrenIDs = make([]uint64, 0)
		modelCategory.ChildrenIDs = append(modelCategory.ChildrenIDs, allChildrenIDs[uint64(modelCategory.ID)]...)
		*modelStoreCategories = append(*modelStoreCategories, modelCategory)
	}
}
