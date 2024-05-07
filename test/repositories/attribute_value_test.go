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
	readAttributeValue = []models.AttributeValuesWithDetail{
		{
			AttributeValues: models.AttributeValues{
				AttributeID:    1,
				AttributeValue: "125g",
			},
			AttributeName: "weight",
		},
	}
)

func TestReadByProductIDAttributeValue(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)
	test_utils.ResetAttributeValuesDB(db)

	// Setup
	modelAttr := []models.AttributeValuesWithDetail{}
	attrValueRepo := repositories.NewRepositoryAttributeValue(db)

	// Assertions
	if assert.NoError(t, attrValueRepo.ReadByProductID(&modelAttr, 1)) {
		if assert.Equal(t, len(readAttributeValue), len(modelAttr)) {
			readAttributeValue[0].Model.ID = modelAttr[0].Model.ID
			readAttributeValue[0].CreatedAt = modelAttr[0].CreatedAt
			readAttributeValue[0].UpdatedAt = modelAttr[0].UpdatedAt

			assert.Equal(t, readAttributeValue[0], modelAttr[0])
		}
	}
}

func TestReadAttributeValuesByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)
	test_utils.ResetAttributeValuesDB(db)

	// Setup
	modelVal := models.AttributeValuesWithDetail{}
	valRepo := repositories.NewRepositoryAttributeValue(db)

	// Assertions
	if assert.NoError(t, valRepo.ReadByAttrValID(&modelVal, 1)) {
		readAttributeValue[0].CreatedAt = modelVal.CreatedAt
		readAttributeValue[0].UpdatedAt = modelVal.UpdatedAt
		readAttributeValue[0].ID = 1
		assert.Equal(t, readAttributeValue[0], modelVal)
	}
}

func TestReadAttributeValueByAttrID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)
	test_utils.ResetAttributeValuesDB(db)

	// Setup
	modelVals := []models.AttributeValuesWithDetail{}
	valRepo := repositories.NewRepositoryAttributeValue(db)

	// Assertions
	if assert.NoError(t, valRepo.ReadByAttrID(&modelVals, 1)) {
		readAttributeValue[0].CreatedAt = modelVals[0].CreatedAt
		readAttributeValue[0].UpdatedAt = modelVals[0].UpdatedAt
		readAttributeValue[0].ID = 1
		assert.Equal(t, readAttributeValue[0], modelVals[0])
	}
}

func TestReadAttributeValuesByIDs(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)
	test_utils.ResetAttributeValuesDB(db)

	// Setup
	modelVals := []models.AttributeValuesWithDetail{}
	valRepo := repositories.NewRepositoryAttributeValue(db)

	// Assertions
	if assert.NoError(t, valRepo.ReadByIDs(&modelVals, []uint64{1})) {
		readAttributeValue[0].CreatedAt = modelVals[0].CreatedAt
		readAttributeValue[0].UpdatedAt = modelVals[0].UpdatedAt
		readAttributeValue[0].ID = 1
		assert.Equal(t, readAttributeValue, modelVals)
	}
}
