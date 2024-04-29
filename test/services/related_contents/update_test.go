package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	contsvc "OnlineStoreBackend/services/related_contents"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	contentRequest = requests.RequestProductContent{
		ContentIDs: []uint64{2},
	}
)

func TestUpdateRelatedContents(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetRelatedContentsDB(db)

	// Setup
	var contService = contsvc.NewServiceProductContent(db)

	// Assertions
	assert.NoError(t, contService.Update(1, &contentRequest))
}
