package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	tagsvc "OnlineStoreBackend/services/tags"

	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	reqUpdateTag = requests.RequestTag{
		Name: "1",
	}
)

func TestUpdateTag(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)

	// Setup
	tagService := tagsvc.NewServiceTag(db)
	modelTag := models.Tags{}

	// Assertions
	assert.NoError(t, tagService.Update(1, &modelTag, &reqUpdateTag))
}
