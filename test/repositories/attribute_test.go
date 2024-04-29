package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readAttributes = []models.Attributes{
		{
			ProductID:     1,
			AttributeName: "weight",
		},
	}
)

func TestReadByNameAttribute(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetAttributesDB(db)

	// Setup
	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(db)

	// Assertions
	if assert.NoError(t, attrRepo.ReadByName(&modelAttr, "weight")) {
		readAttributes[0].Model.ID = modelAttr.Model.ID
		readAttributes[0].CreatedAt = modelAttr.CreatedAt
		readAttributes[0].UpdatedAt = modelAttr.UpdatedAt

		assert.Equal(t, readAttributes[0], modelAttr)
	}
}

func TestReadByProductIDAttribute(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetAttributesDB(db)

	// Setup
	modelAttr := []models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(db)

	// Assertions
	if assert.NoError(t, attrRepo.ReadByProductID(&modelAttr, 1)) {
		readAttributes[0].Model.ID = modelAttr[0].Model.ID
		readAttributes[0].CreatedAt = modelAttr[0].CreatedAt
		readAttributes[0].UpdatedAt = modelAttr[0].UpdatedAt

		assert.Equal(t, readAttributes[0], modelAttr[0])
	}
}

func TestReadByIDAttribute(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetAttributesDB(db)

	// Setup
	modelAttr := models.Attributes{}
	attrRepo := repositories.NewRepositoryAttribute(db)

	// Assertions
	if assert.NoError(t, attrRepo.ReadByID(&modelAttr, 1)) {
		readAttributes[0].Model.ID = modelAttr.Model.ID
		readAttributes[0].CreatedAt = modelAttr.CreatedAt
		readAttributes[0].UpdatedAt = modelAttr.UpdatedAt

		assert.Equal(t, readAttributes[0], modelAttr)
	}
}
