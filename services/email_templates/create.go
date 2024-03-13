package etsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelTemplate *models.EmailTemplates, req *requests.RequestEmailTemplate, storeID uint64) error {
	modelTemplate.StoreID = storeID
	modelTemplate.OrderStatus = req.OrderStatus
	modelTemplate.CompanyName = req.CompanyName
	modelTemplate.CompanyLink = req.CompanyLink
	modelTemplate.CompanyLogoUrl = req.CompanyLogoUrl
	modelTemplate.CompanyPrimaryColor = req.CompanyPrimaryColor
	modelTemplate.EmailPretext = req.EmailPretext
	modelTemplate.HeaderPosterSloganTitle = req.HeaderPosterSloganTitle
	modelTemplate.HeaderPosterSloganSubtitle = req.HeaderPosterSloganSubtitle
	modelTemplate.BodyGreeting = req.BodyGreeting
	modelTemplate.FirstName = req.FirstName
	modelTemplate.BodyMessage = req.BodyMessage
	modelTemplate.BodyCtaBtnLink = req.BodyCtaBtnLink
	modelTemplate.BodyCtaBtnLabel = req.BodyCtaBtnLabel
	modelTemplate.BodySecondaryMessage = req.BodySecondaryMessage
	modelTemplate.UnsubscribeLink = req.UnsubscribeLink
	modelTemplate.UnsubscribeSafeLink = req.UnsubscribeSafeLink

	return service.DB.Create(&modelTemplate).Error
}
