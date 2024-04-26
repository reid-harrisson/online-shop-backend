package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	tagsvc "OnlineStoreBackend/services/tags"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tagInputs = []models.Tags{
		{
			StoreID: 1,
			Name:    "sauces",
		},
	}
)

func TestCreateTagsWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)

	// Setup
	var tagService = tagsvc.NewServiceTag(db)
	var modelTags = tagInputs

	// Assertions
	if assert.NoError(t, tagService.CreateWithCSV(&modelTags, []string{"sauces"}, map[string]int{"sauces": 0})) {
		assert.Equal(t, modelTags, tagInputs)
	}
}
