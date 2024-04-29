package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	storeOutputs = []models.Stores{
		{
			CompanyID:            1,
			OwnerID:              1,
			Name:                 "Steve",
			ContactPhone:         "+27793267663",
			ContactEmail:         "steve@pockittv.com",
			ShowStockLevelStatus: 0,
			ShowOutOfStockStatus: 0,
			DeliveryPolicy:       "",
			ReturnsPolicy:        "",
			Terms:                "",
		},
		{
			CompanyID:            2,
			OwnerID:              2,
			Name:                 "Tade",
			ContactPhone:         "+0824721073",
			ContactEmail:         "jade@pockittv.com",
			ShowStockLevelStatus: 0,
			ShowOutOfStockStatus: 0,
			DeliveryPolicy:       "",
			ReturnsPolicy:        "",
			Terms:                "",
		},
	}
)

func TestReadAll(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetUsersDB(db)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetStoresDB(db)

	// Setup
	var storeRpo = repositories.NewRepositoryStore(db)
	var modelStores = []models.Stores{}

	// Assertions
	if assert.NoError(t, storeRpo.ReadAll(&modelStores)) {
		if assert.Equal(t, 2, len(modelStores)) {
			for index, modelStore := range modelStores {
				storeOutputs[index].CreatedAt = modelStore.CreatedAt
				storeOutputs[index].UpdatedAt = modelStore.UpdatedAt
				storeOutputs[index].ID = modelStore.ID
				assert.Equal(t, storeOutputs[index], modelStore)
			}
		}
	}
}

func TestReadByUser(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetUsersDB(db)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetStoresDB(db)

	// Setup
	var storeRpo = repositories.NewRepositoryStore(db)
	var modelStores = []models.Stores{}

	// Assertions
	if assert.NoError(t, storeRpo.ReadByUser(&modelStores, 1)) {
		if assert.Equal(t, 1, len(modelStores)) {
			for index, modelStore := range modelStores {
				storeOutputs[index].CreatedAt = modelStore.CreatedAt
				storeOutputs[index].UpdatedAt = modelStore.UpdatedAt
				storeOutputs[index].ID = modelStore.ID
				assert.Equal(t, storeOutputs[index], modelStore)
			}
		}
	}
}
