package etsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(emailTemplateID uint64, modelEmailTemplate *models.EmailTemplate, requestEmailTemplate *requests.RequestEmailTemplate) error {
	modelEmailTemplate.StoreID = requestEmailTemplate.StoreID
	modelEmailTemplate.OrderStatus = requestEmailTemplate.OrderStatus
	modelEmailTemplate.CompanyName = requestEmailTemplate.CompanyName
	modelEmailTemplate.CompanyLink = requestEmailTemplate.CompanyLink
	modelEmailTemplate.CompanyLogoUrl = requestEmailTemplate.CompanyLogoUrl
	modelEmailTemplate.CompanyPrimaryColor = requestEmailTemplate.CompanyPrimaryColor
	modelEmailTemplate.EmailPretext = requestEmailTemplate.EmailPretext
	modelEmailTemplate.HeaderPosterSloganTitle = requestEmailTemplate.HeaderPosterSloganTitle
	modelEmailTemplate.HeaderPosterSloganSubtitle = requestEmailTemplate.HeaderPosterSloganSubtitle
	modelEmailTemplate.BodyGreeting = requestEmailTemplate.BodyGreeting
	modelEmailTemplate.FirstName = requestEmailTemplate.FirstName
	modelEmailTemplate.BodyMessage = requestEmailTemplate.BodyMessage
	modelEmailTemplate.BodyCtaBtnLink = requestEmailTemplate.BodyCtaBtnLink
	modelEmailTemplate.BodyCtaBtnLabel = requestEmailTemplate.BodyCtaBtnLabel
	modelEmailTemplate.BodySecondaryMessage = requestEmailTemplate.BodySecondaryMessage
	modelEmailTemplate.UnsubscribeLink = requestEmailTemplate.UnsubscribeLink
	modelEmailTemplate.UnsubscribeSafeLink = requestEmailTemplate.UnsubscribeSafeLink

	service.DB.Save(&modelEmailTemplate)
	return nil
}
