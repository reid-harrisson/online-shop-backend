package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	etsvc "OnlineStoreBackend/services/email_templates"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	reqEmailTemplate = requests.RequestEmailTemplate{
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
	}
	templateInputs = []models.EmailTemplates{
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

func TestCreateTemplates(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTemplatesDB(db)

	// Setup
	temService := etsvc.NewServiceEmailTemplate(db)
	modelTemplates := templateInputs[0]

	// Assertions
	if assert.NoError(t, temService.Create(&modelTemplates, &reqEmailTemplate, 1)) {
		templateInputs[0].Model.ID = modelTemplates.Model.ID
		templateInputs[0].CreatedAt = modelTemplates.CreatedAt
		templateInputs[0].UpdatedAt = modelTemplates.UpdatedAt

		assert.Equal(t, templateInputs[0], modelTemplates)
	}
}
