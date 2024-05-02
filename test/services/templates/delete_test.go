package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	etsvc "OnlineStoreBackend/services/email_templates"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTag(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTemplatesDB(db)

	// Setup
	temService := etsvc.NewServiceEmailTemplate(db)

	// Assertions
	assert.NoError(t, temService.Delete(1))
}
