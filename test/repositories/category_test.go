package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	categoryOutputs = []models.Categories{
		{
			StoreID:  1,
			ParentID: 0,
			Name:     "Kefir",
		},
	}
	categoryWithChildrenOutputs = []models.CategoriesWithChildren{
		{
			Categories: models.Categories{
				StoreID:  1,
				ParentID: 0,
				Name:     "Kefir",
			},
			ChildrenIDs: []uint64{},
		},
	}
	prodcateWithNameOutputs = []models.ProductCategoriesWithName{
		{
			ProductCategories: models.ProductCategories{
				Model: gorm.Model{
					ID: 1,
				},
				ProductID:  1,
				CategoryID: 1,
			},
			CategoryName: "Kefir",
		},
	}
)

func TestCategoryReadByStoreID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCategoriesDB(db)

	// Setup
	cateRepo := repositories.NewRepositoryCategory(db)
	modelCategories := make([]models.CategoriesWithChildren, 0)

	// Assertions
	if assert.NoError(t, cateRepo.ReadByStoreID(&modelCategories, 1)) {
		categoryWithChildrenOutputs[0].ID = modelCategories[0].ID
		categoryWithChildrenOutputs[0].CreatedAt = modelCategories[0].CreatedAt
		categoryWithChildrenOutputs[0].UpdatedAt = modelCategories[0].UpdatedAt

		assert.Equal(t, categoryWithChildrenOutputs[0], modelCategories[0])
	}
}

func TestCategoryReadByName(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCategoriesDB(db)

	// Setup
	cateRepo := repositories.NewRepositoryCategory(db)
	modelCategory := models.Categories{}

	// Assertions
	if assert.NoError(t, cateRepo.ReadByName(&modelCategory, "Kefir", uint64(1))) {
		categoryOutputs[0].ID = modelCategory.ID
		categoryOutputs[0].CreatedAt = modelCategory.CreatedAt
		categoryOutputs[0].UpdatedAt = modelCategory.UpdatedAt

		assert.Equal(t, categoryOutputs[0], modelCategory)
	}
}

func TestCategoryReadByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCategoriesDB(db)

	// Setup
	cateRepo := repositories.NewRepositoryCategory(db)
	modelCategory := models.Categories{}

	// Assertions
	if assert.NoError(t, cateRepo.ReadByID(&modelCategory, uint64(1))) {
		categoryOutputs[0].ID = modelCategory.ID
		categoryOutputs[0].CreatedAt = modelCategory.CreatedAt
		categoryOutputs[0].UpdatedAt = modelCategory.UpdatedAt

		assert.Equal(t, categoryOutputs[0], modelCategory)
	}
}

func TestReadCategoriesByProductID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCategoriesDB(db)
	test_utils.ResetProductCategoriesDB(db)

	// Setup
	cateRepo := repositories.NewRepositoryCategory(db)
	modelCategories := []models.ProductCategoriesWithName{}

	// Assertions
	if assert.NoError(t, cateRepo.ReadByProductID(&modelCategories, 1)) {
		prodcateWithNameOutputs[0].CreatedAt = modelCategories[0].CreatedAt
		prodcateWithNameOutputs[0].UpdatedAt = modelCategories[0].UpdatedAt
		assert.Equal(t, prodcateWithNameOutputs, modelCategories)
	}
}
