package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	readProductsWithDetails = []models.ProductsWithDetail{
		{
			Products: readProducts[0],
		},
	}
	readProducts = []models.Products{
		{
			Model: gorm.Model{
				ID: 1,
			},
			StoreID:           1,
			Title:             "Gochujang - Korean Chilli Pepper Paste",
			ShortDescription:  "Gochujang is a Sticky, Sweet, Savoury &amp; SPICY Chilli Paste.",
			LongDescription:   "Our AMAZING range of products are available nation wide in South Africa at select health stores.",
			ImageUrls:         "https://www.chegourmet.co.za/wp-content/uploads/2019/09/Gochujang-Front-scaled.jpg",
			MinimumStockLevel: 0,
			Status:            0,
			Sku:               "44",
			Type:              1,
			ShippingClass:     "Courier Refrigerated",
		},
	}
	readProductsWithLink = []models.ProductsWithLink{
		{
			Products: models.Products{
				StoreID:           1,
				Title:             "Gochujang - Korean Chilli Pepper Paste",
				ShortDescription:  "Gochujang is a Sticky, Sweet, Savoury &amp; SPICY Chilli Paste.",
				LongDescription:   "Our AMAZING range of products are available nation wide in South Africa at select health stores.",
				ImageUrls:         "https://www.chegourmet.co.za/wp-content/uploads/2019/09/Gochujang-Front-scaled.jpg",
				MinimumStockLevel: 0,
				Status:            0,
				Sku:               "44",
				Type:              1,
				ShippingClass:     "Courier Refrigerated",
			},
			IsUpCross: utils.UpSell,
		},
	}
	readProductsApproved = []models.ProductsApproved{
		{
			ID:           1,
			Title:        "Gochujang - Korean Chilli Pepper Paste",
			MinimumPrice: 76,
			MaximumPrice: 76,
			RegularPrice: 96,
			Rating:       0,
			ImageUrls:    "https://www.chegourmet.co.za/wp-content/uploads/2019/09/Gochujang-Front-scaled.jpg",
		},
	}
)

func TestReadByIDProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	if assert.NoError(t, prodRepo.ReadByID(&modelProduct, 1)) {
		readProducts[0].Model.ID = modelProduct.Model.ID
		readProducts[0].CreatedAt = modelProduct.CreatedAt
		readProducts[0].UpdatedAt = modelProduct.UpdatedAt

		assert.Equal(t, readProducts[0], modelProduct)
	}
}

func TestReadDetailProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProduct := models.ProductsWithDetail{}
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	if assert.NoError(t, prodRepo.ReadDetail(&modelProduct, 1)) {
		readProductsWithDetails[0].CreatedAt = modelProduct.CreatedAt
		readProductsWithDetails[0].UpdatedAt = modelProduct.UpdatedAt
		assert.Equal(t, readProductsWithDetails[0].Products, modelProduct.Products)
	}
}

func TestReadLinkedProductsProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductLinksDB(db)

	// Setup
	modelProducts := make([]models.ProductsWithLink, 0)
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	if assert.NoError(t, prodRepo.ReadLinkedProducts(&modelProducts, 1)) {
		if assert.Equal(t, len(readProductsWithLink), len(modelProducts)) {
			readProductsWithLink[0].Model.ID = modelProducts[0].Model.ID
			readProductsWithLink[0].CreatedAt = modelProducts[0].CreatedAt
			readProductsWithLink[0].UpdatedAt = modelProducts[0].UpdatedAt
			assert.Equal(t, readProductsWithLink[0], modelProducts[0])
		}
	}
}

func TestReadAllProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProducts := []models.Products{}
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	if assert.NoError(t, prodRepo.ReadAll(&modelProducts, 1, "")) {
		readProducts[0].Model.ID = modelProducts[0].Model.ID
		readProducts[0].CreatedAt = modelProducts[0].CreatedAt
		readProducts[0].UpdatedAt = modelProducts[0].UpdatedAt

		assert.Equal(t, readProducts[0], modelProducts[0])
	}
}

func TestReadPagingProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProducts := []models.Products{}
	prodRepo := repositories.NewRepositoryProduct(db)
	totalCount := int64(0)

	// Assertions
	if assert.NoError(t, prodRepo.ReadPaging(&modelProducts, 0, 1, 1, "", &totalCount)) {
		readProducts[0].Model.ID = modelProducts[0].Model.ID
		readProducts[0].CreatedAt = modelProducts[0].CreatedAt
		readProducts[0].UpdatedAt = modelProducts[0].UpdatedAt

		assert.Equal(t, readProducts[0], modelProducts[0])
	}
}

func TestReadApprovedProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProductsApproved := []models.ProductsApproved{}
	prodRepo := repositories.NewRepositoryProduct(db)
	totalCount := int64(0)

	// Assertions
	if assert.NoError(t, prodRepo.ReadApproved(&modelProductsApproved, 1, 1, 0, 1, &totalCount)) {
		assert.Equal(t, readProductsApproved[0], modelProductsApproved[0])
	}
}

func TestReadByCategoryProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProducts := make([]models.Products, 1)
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	if assert.NoError(t, prodRepo.ReadByCategory(&modelProducts, 1, 1)) {
		readProducts[0].Model.ID = modelProducts[0].Model.ID
		readProducts[0].CreatedAt = modelProducts[0].CreatedAt
		readProducts[0].UpdatedAt = modelProducts[0].UpdatedAt

		assert.Equal(t, readProducts[0], modelProducts[0])
	}
}

func TestReadByTagsProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	modelProducts := make([]models.Products, 1)
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	assert.NoError(t, prodRepo.ReadByTags(&modelProducts, 1, []string{}, ""))
}
