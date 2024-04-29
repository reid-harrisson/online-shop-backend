package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	prodtagRequest = requests.RequestProductTag{
		Tags: []string{"kefir"},
	}
	prodtagUpdates = []models.ProductTagsWithName{
		{
			ProductTags: models.ProductTags{
				Model: gorm.Model{
					ID: 1,
				},
				ProductID: 1,
				TagID:     1,
			},
			TagName: "kefir",
		},
		{
			ProductTags: models.ProductTags{
				Model: gorm.Model{
					ID: 3,
				},
				ProductID: 1,
				TagID:     1,
			},
			TagName: "kefir",
		},
	}
)

func TestUpdateProductTags(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var prodtagService = prodtagsvc.NewServiceProductTag(db)
	var modelProdTags = []models.ProductTagsWithName{}

	// Assertions
	if assert.NoError(t, prodtagService.Update(&modelProdTags, &prodtagRequest, &models.Products{Model: gorm.Model{ID: 1}, StoreID: 1})) {
		if assert.Equal(t, len(prodtagUpdates), len(modelProdTags)) {
			for index := range modelProdTags {
				prodtagUpdates[index].CreatedAt = modelProdTags[index].CreatedAt
				prodtagUpdates[index].UpdatedAt = modelProdTags[index].UpdatedAt
				assert.Equal(t, prodtagUpdates[index], modelProdTags[index])
			}
		}
	}
}
