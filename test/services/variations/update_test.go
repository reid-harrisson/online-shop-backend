package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateVariation(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetVariationDetailsDB(db)
	test_utils.ResetAttributesDB(db)
	test_utils.ResetAttributeValuesDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var varService = prodvarsvc.NewServiceVariation(db)
	var varRepo = repositories.NewRepositoryVariation(db)
	var modelVar = models.Variations{}

	// Assertions
	if assert.NoError(t, varRepo.ReadByID(&modelVar, 1)) {
		if assert.NoError(t, varService.Update(&modelVar, &varRequest)) {
			varOutput.CreatedAt = modelVar.CreatedAt
			varOutput.UpdatedAt = modelVar.UpdatedAt
			varOutput.ID = 1
			assert.Equal(t, varOutput, modelVar)
		}
	}
}
