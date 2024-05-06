package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	linksvc "OnlineStoreBackend/services/links"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	linkInputs = []models.Links{
		{
			ProductID: 1,
			LinkID:    2,
			IsUpCross: utils.UpSell,
		},
		{
			ProductID: 2,
			LinkID:    1,
			IsUpCross: utils.CrossSell,
		},
	}
	linkMatches = []string{}
	linkIndices = map[string]int{}
)

func TestCreateLinks(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductLinksDB(db)

	// Setup
	linkService := linksvc.NewServiceLink(db)

	// Assertions
	assert.NoError(t, linkService.Create(1, 2, utils.CrossSell))
}

func TestCreateLinksWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductLinksDB(db)

	// Setup
	linkService := linksvc.NewServiceLink(db)
	modelLinks := linkInputs

	// Assertions
	if assert.NoError(t, linkService.CreateWithCSV(&modelLinks, linkMatches, linkIndices)) {
		assert.Equal(t, modelLinks[0].ID, uint(3))
		assert.Equal(t, modelLinks[1].ID, uint(4))
	}
}
