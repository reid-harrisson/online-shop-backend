package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readTags = []models.Tags{
		{
			StoreID: 1,
			Name:    "kefir",
		},
	}
)

func TestReadByNameTag(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)

	// Setup
	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(db)

	// Assertions
	if assert.NoError(t, tagRepo.ReadByName(&modelTag, "kefir", 1)) {
		readTags[0].Model.ID = modelTag.Model.ID
		readTags[0].CreatedAt = modelTag.CreatedAt
		readTags[0].UpdatedAt = modelTag.UpdatedAt

		assert.Equal(t, readTags[0], modelTag)
	}
}
