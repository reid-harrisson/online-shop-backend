package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	catesvc "OnlineStoreBackend/services/categories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	reqUpdateCategory = requests.RequestCategory{
		Name:     "1",
		ParentID: 1,
	}
)

func TestUpdateCategory(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetCategoriesDB(db)

	// Setup
	cateService := catesvc.NewServiceCategory(db)
	modelCategory := models.Categories{}

	// Assertions
	assert.NoError(t, cateService.Update(1, &modelCategory, &reqUpdateCategory))
}
