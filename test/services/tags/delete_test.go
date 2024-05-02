package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	tagsvc "OnlineStoreBackend/services/tags"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTag(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)

	// Setup
	var tagService = tagsvc.NewServiceTag(db)

	// Assertions
	assert.NoError(t, tagService.Delete(1))
}
