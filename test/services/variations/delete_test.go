package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDeleteVariation(t *testing.T) {
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
	if assert.NoError(t, varService.Delete(1)) {
		assert.Equal(t, gorm.ErrRecordNotFound, varRepo.ReadByID(&modelVar, 1))
	}
}
