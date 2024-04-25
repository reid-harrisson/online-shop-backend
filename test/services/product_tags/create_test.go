package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodtagsvc "OnlineStoreBackend/services/product_tags"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	prodtagInputs = []models.ProductTags{
		{
			ProductID: 1,
			TagID:     1,
		},
	}
)

func TestCreateProductTagsWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var prodtagService = prodtagsvc.NewServiceProductTag(db)
	var modelProdTags = prodtagInputs

	// Assertions
	if assert.NoError(t, prodtagService.CreateWithCSV(&modelProdTags, []string{"1:1"}, map[string]int{"1:1": 0})) {
		assert.Equal(t, modelProdTags, prodtagInputs)
	}
}
