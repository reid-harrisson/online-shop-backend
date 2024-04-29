package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	tagsvc "OnlineStoreBackend/services/tags"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	tagInputs = []models.Tags{
		{
			Model: gorm.Model{
				ID: 1,
			},
			StoreID: 1,
			Name:    "sauces",
		},
	}
)

func TestCreateTags(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)

	// Setup
	var tagService = tagsvc.NewServiceTag(db)
	var modelTags = tagInputs[0]

	// Assertions
	if assert.NoError(t, tagService.Create(&modelTags, "sauces", 1)) {
		tagInputs[0].Model.ID = modelTags.Model.ID
		tagInputs[0].CreatedAt = modelTags.CreatedAt
		tagInputs[0].UpdatedAt = modelTags.UpdatedAt

		assert.Equal(t, tagInputs[0], modelTags)
	}
}

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
