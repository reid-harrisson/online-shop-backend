package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
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
