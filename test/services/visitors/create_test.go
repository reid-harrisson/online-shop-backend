package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	vistsvc "OnlineStoreBackend/services/visitors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	visitorInput = requests.RequestVisitor{
		StoreID:     1,
		ProductID:   2,
		IpAddress:   "111.111.111.113",
		Page:        "Cart",
		Bounce:      2,
		LoadingTime: 0,
	}
)

func TestCreateVisitor(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetVisitorsDB(db)

	// Setup
	var visitService = vistsvc.NewServiceVisitor(db)
	var modelVisitor = models.Visitors{}

	// Assertions
	if assert.NoError(t, visitService.Create(&modelVisitor, &visitorInput)) {
		assert.Equal(t, uint(3), modelVisitor.ID)
		assert.Equal(t, visitorInput.StoreID, modelVisitor.StoreID)
		assert.Equal(t, visitorInput.ProductID, modelVisitor.ProductID)
		assert.Equal(t, visitorInput.IpAddress, modelVisitor.IpAddress)
		assert.Equal(t, utils.CartPage, modelVisitor.Page)
		assert.Equal(t, visitorInput.Bounce, modelVisitor.Bounce)
		assert.Equal(t, visitorInput.LoadingTime, modelVisitor.LoadingTime)
	}
}
