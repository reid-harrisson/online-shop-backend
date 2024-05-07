package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	storesvc "OnlineStoreBackend/services/stores"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	storeInput = requests.RequestStore{
		Name:                 "Steve",
		ContactPhone:         "+27793267663",
		ContactEmail:         "steve@pockittv.com",
		ShowStockLevelStatus: 0,
		ShowOutOfStockStatus: 0,
		DeliveryPolicy:       "",
		ReturnsPolicy:        "",
		Terms:                "",
	}
)

func TestCreateStore(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetUsersDB(db)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetStoresDB(db)

	// Setup
	var storeService = storesvc.NewServiceStore(db)
	var modelStore = models.Stores{}

	// Assertions
	if assert.NoError(t, storeService.Create(&modelStore, &storeInput, 1)) {
		assert.Equal(t, uint(3), modelStore.ID)
		assert.Equal(t, storeInput.Name, modelStore.Name)
		assert.Equal(t, storeInput.ContactPhone, modelStore.ContactPhone)
		assert.Equal(t, storeInput.ContactEmail, modelStore.ContactEmail)
		assert.Equal(t, storeInput.ShowStockLevelStatus, modelStore.ShowStockLevelStatus)
		assert.Equal(t, storeInput.ShowOutOfStockStatus, modelStore.ShowOutOfStockStatus)
		assert.Equal(t, storeInput.DeliveryPolicy, modelStore.DeliveryPolicy)
		assert.Equal(t, storeInput.ReturnsPolicy, modelStore.ReturnsPolicy)
		assert.Equal(t, storeInput.Terms, modelStore.Terms)
	}
}
