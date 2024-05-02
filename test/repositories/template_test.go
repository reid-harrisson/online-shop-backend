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
	readTemplates = []models.EmailTemplates{
		{
			StoreID:                    1,
			OrderStatus:                1,
			CompanyName:                "company1",
			CompanyLink:                "",
			CompanyLogoUrl:             "",
			CompanyPrimaryColor:        "",
			EmailPretext:               "",
			HeaderPosterSloganTitle:    "",
			HeaderPosterSloganSubtitle: "",
			BodyGreeting:               "",
			FirstName:                  "",
			BodyMessage:                "",
			BodyCtaBtnLink:             "",
			BodyCtaBtnLabel:            "",
			BodySecondaryMessage:       "",
			UnsubscribeLink:            "",
			UnsubscribeSafeLink:        "",
		},
	}
)

func TestReadByStoreIDTemplate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTemplatesDB(db)

	// Setup
	modelTemplates := make([]models.EmailTemplates, 0)
	temRepo := repositories.NewRepositoryEmailTemplate(db)

	// Assertions
	if assert.NoError(t, temRepo.ReadByStoreID(&modelTemplates, 1)) {
		readTemplates[0].Model.ID = modelTemplates[0].Model.ID
		readTemplates[0].CreatedAt = modelTemplates[0].CreatedAt
		readTemplates[0].UpdatedAt = modelTemplates[0].UpdatedAt

		assert.Equal(t, readTemplates[0], modelTemplates[0])
	}
}
