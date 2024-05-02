package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	catesvc "OnlineStoreBackend/services/categories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryInputs = []models.Categories{
		{
			StoreID:  1,
			ParentID: 0,
			Name:     "Sauces",
		},
	}
	reqCategory = requests.RequestCategory{
		Name:     "Sauces",
		ParentID: 0,
	}
)

func TestCreateCategoriesWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetCategoriesDB(db)

	// Setup
	var cateService = catesvc.NewServiceCategory(db)
	var modelCategories = categoryInputs

	// Assertions
	if assert.NoError(t, cateService.CreateWithCSV(&modelCategories, []string{"Sauces"}, map[string]string{}, map[string]int{"Sauces": 0})) {
		assert.Equal(t, modelCategories, categoryInputs)
	}
}

func TestCreateCategories(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetCategoriesDB(db)

	// Setup
	var cateService = catesvc.NewServiceCategory(db)
	modelCategory := models.Categories{}

	// Assertions
	if assert.NoError(t, cateService.Create(&modelCategory, &reqCategory, 1)) {
		categoryInputs[0].ID = modelCategory.ID
		categoryInputs[0].CreatedAt = modelCategory.CreatedAt
		categoryInputs[0].UpdatedAt = modelCategory.UpdatedAt

		assert.Equal(t, categoryInputs[0], modelCategory)
	}
}
