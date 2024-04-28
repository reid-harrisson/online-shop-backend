package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	prodsvc "OnlineStoreBackend/services/products"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	updateProdInput = requests.RequestProduct{
		StoreID:          1,
		Title:            "Kefir",
		ShortDescription: "Short description of Kefir",
		LongDescription:  "Long description of Kefir",
		ImageUrls:        []string{"https://kefir-front.jpg", "https://kefir-side.jpg"},
	}
)

func TestUpdateStatus(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)

	// Assertions
	assert.NoError(t, prodService.UpdateStatus(1, utils.Pending))
}

func TestUpdate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)
	var modelProduct = models.Products{Model: gorm.Model{ID: 1}}

	// Assertions
	if assert.NoError(t, prodService.Update(&modelProduct, &updateProdInput)) {
		assert.Equal(t, updateProdInput.StoreID, modelProduct.StoreID)
		assert.Equal(t, updateProdInput.Title, modelProduct.Title)
		assert.Equal(t, updateProdInput.ShortDescription, modelProduct.ShortDescription)
		assert.Equal(t, updateProdInput.LongDescription, modelProduct.LongDescription)
		imageUrls, _ := json.Marshal(updateProdInput.ImageUrls)
		assert.Equal(t, string(imageUrls), modelProduct.ImageUrls)
	}
}
