package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseEmailTemplate struct {
	ID                         uint64 `json:"id"`
	StoreID                    uint64 `json:"store_id"`
	OrderStatus                uint64 `json:"order_status"`
	CompanyName                string `json:"company_name"`
	CompanyLink                string `json:"company_link"`
	CompanyLogoUrl             string `json:"company_logo_url"`
	CompanyPrimaryColor        string `json:"company_primary_color"`
	EmailPretext               string `json:"email_pretext"`
	HeaderPosterSloganTitle    string `json:"header_poster_slogan_title"`
	HeaderPosterSloganSubtitle string `json:"header_poster_slogan_subtitle"`
	BodyGreeting               string `json:"body_greeting"`
	FirstName                  string `json:"first_name"`
	BodyMessage                string `json:"body_message"`
	BodyCtaBtnLink             string `json:"body_cta_btn_link"`
	BodyCtaBtnLabel            string `json:"body_cta_btn_label"`
	BodySecondaryMessage       string `json:"body_secondary_message"`
	UnsubscribeLink            string `json:"unsubscribe_link"`
	UnsubscribeSafeLink        string `json:"unsubscribe_safe_link"`
}

func NewResponseEmailTemplate(c echo.Context, statusCode int, modelEmailTemplate *models.EmailTemplate) error {
	responseEmailTemplate := ResponseEmailTemplate{
		ID:                         uint64(modelEmailTemplate.ID),
		StoreID:                    modelEmailTemplate.StoreID,
		OrderStatus:                modelEmailTemplate.OrderStatus,
		CompanyName:                modelEmailTemplate.CompanyName,
		CompanyLink:                modelEmailTemplate.CompanyLink,
		CompanyLogoUrl:             modelEmailTemplate.CompanyLogoUrl,
		CompanyPrimaryColor:        modelEmailTemplate.CompanyPrimaryColor,
		EmailPretext:               modelEmailTemplate.EmailPretext,
		HeaderPosterSloganTitle:    modelEmailTemplate.HeaderPosterSloganTitle,
		HeaderPosterSloganSubtitle: modelEmailTemplate.HeaderPosterSloganSubtitle,
		BodyGreeting:               modelEmailTemplate.BodyGreeting,
		FirstName:                  modelEmailTemplate.FirstName,
		BodyMessage:                modelEmailTemplate.BodyMessage,
		BodyCtaBtnLink:             modelEmailTemplate.BodyCtaBtnLink,
		BodyCtaBtnLabel:            modelEmailTemplate.BodyCtaBtnLabel,
		BodySecondaryMessage:       modelEmailTemplate.BodySecondaryMessage,
		UnsubscribeLink:            modelEmailTemplate.UnsubscribeLink,
		UnsubscribeSafeLink:        modelEmailTemplate.UnsubscribeSafeLink,
	}
	return Response(c, statusCode, responseEmailTemplate)
}

func NewResponseEmailTemplates(c echo.Context, statusCode int, modelEmailTemplates []models.EmailTemplate) error {
	responseEmailTemplates := make([]ResponseEmailTemplate, 0)
	for _, modelEmailTemplate := range modelEmailTemplates {
		responseEmailTemplates = append(responseEmailTemplates, ResponseEmailTemplate{
			ID:                         uint64(modelEmailTemplate.ID),
			StoreID:                    modelEmailTemplate.StoreID,
			OrderStatus:                modelEmailTemplate.OrderStatus,
			CompanyName:                modelEmailTemplate.CompanyName,
			CompanyLink:                modelEmailTemplate.CompanyLink,
			CompanyLogoUrl:             modelEmailTemplate.CompanyLogoUrl,
			CompanyPrimaryColor:        modelEmailTemplate.CompanyPrimaryColor,
			EmailPretext:               modelEmailTemplate.EmailPretext,
			HeaderPosterSloganTitle:    modelEmailTemplate.HeaderPosterSloganTitle,
			HeaderPosterSloganSubtitle: modelEmailTemplate.HeaderPosterSloganSubtitle,
			BodyGreeting:               modelEmailTemplate.BodyGreeting,
			FirstName:                  modelEmailTemplate.FirstName,
			BodyMessage:                modelEmailTemplate.BodyMessage,
			BodyCtaBtnLink:             modelEmailTemplate.BodyCtaBtnLink,
			BodyCtaBtnLabel:            modelEmailTemplate.BodyCtaBtnLabel,
			BodySecondaryMessage:       modelEmailTemplate.BodySecondaryMessage,
			UnsubscribeLink:            modelEmailTemplate.UnsubscribeLink,
			UnsubscribeSafeLink:        modelEmailTemplate.UnsubscribeSafeLink,
		})
	}
	return Response(c, statusCode, responseEmailTemplates)
}
