package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	linksvc "OnlineStoreBackend/services/links"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteLinks(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductLinksDB(db)

	// Setup
	linkService := linksvc.NewServiceLink(db)

	// Assertions
	assert.NoError(t, linkService.Delete(1))
}
