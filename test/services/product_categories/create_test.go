package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	prodcateInputs = []models.ProductCategories{
		{
			ProductID:  1,
			CategoryID: 1,
		},
	}
)

func TestCreateProductCategoriesWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCategoriesDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductCategoriesDB(db)

	// Setup
	var prodcateService = prodcatesvc.NewServiceProductCategory(db)
	var modelProdCate = prodcateInputs

	// Assertions
	if assert.NoError(t, prodcateService.CreateWithCSV(&modelProdCate, []string{"1:1"}, map[string]int{"1:1": 0})) {
		assert.Equal(t, modelProdCate, prodcateInputs)
	}
}
