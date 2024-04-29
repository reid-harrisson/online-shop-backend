package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	reqAttrUpdate = requests.RequestAttribute{
		Name: "update name",
	}
)

func TestUpdateAttribute(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	valService := prodattrsvc.NewServiceAttribute(db)
	modelAttr := models.Attributes{}

	// Assertions
	assert.NoError(t, valService.Update(1, &reqAttrUpdate, &modelAttr))
}
