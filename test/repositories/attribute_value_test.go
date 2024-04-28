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
				AttributeValue: "100g",
			},
			AttributeName: "weight",
		},
	}
)

func TestReadByProductIDAttributeValue(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetAttributesDB(db)

	// Setup
	modelAttr := []models.AttributeValuesWithDetail{}
	attrValueRepo := repositories.NewRepositoryAttributeValue(db)

	// Assertions
	if assert.NoError(t, attrValueRepo.ReadByProductID(&modelAttr, 1)) {
		readAttributeValue[0].Model.ID = modelAttr[0].Model.ID
		readAttributeValue[0].CreatedAt = modelAttr[0].CreatedAt
		readAttributeValue[0].UpdatedAt = modelAttr[0].UpdatedAt

		assert.Equal(t, readAttributeValue[0], modelAttr[0])
	}
}
