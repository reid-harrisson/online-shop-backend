package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	prodcatesvc "OnlineStoreBackend/services/product_categories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	prodcateRequest = requests.RequestProductCategory{
		CategoryIDs: []uint64{2},
	}
	prodcateUpdates = []models.ProductCategoriesWithName{
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
		{
			ProductCategories: models.ProductCategories{
				Model: gorm.Model{
					ID: 3,
				},
				ProductID:  1,
				CategoryID: 2,
			},
			CategoryName: "Kimchi",
		},
	}
)

func TestUpdateProductCategories(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCategoriesDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductCategoriesDB(db)

	// Setup
	var prodcateService = prodcatesvc.NewServiceProductCategory(db)
	var modelProdCates = []models.ProductCategoriesWithName{}

	// Assertions
	if assert.NoError(t, prodcateService.Update(&modelProdCates, &prodcateRequest, 1)) {
		if assert.Equal(t, len(prodcateUpdates), len(modelProdCates)) {
			for index := range modelProdCates {
				prodcateUpdates[index].CreatedAt = modelProdCates[index].CreatedAt
				prodcateUpdates[index].UpdatedAt = modelProdCates[index].UpdatedAt
				assert.Equal(t, prodcateUpdates[index], modelProdCates[index])
			}
		}
	}
}
