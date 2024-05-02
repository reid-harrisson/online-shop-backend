package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	etsvc "OnlineStoreBackend/services/email_templates"

	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	reqEmailTemplateUpdate = requests.RequestEmailTemplate{
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
)

func TestUpdateTemplate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTemplatesDB(db)

	// Setup
	temService := etsvc.NewServiceEmailTemplate(db)
	modelTemplate := models.EmailTemplates{}

	// Assertions
	assert.NoError(t, temService.Update(1, &modelTemplate, &reqEmailTemplateUpdate))
}
