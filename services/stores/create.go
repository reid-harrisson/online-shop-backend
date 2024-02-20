package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStore *models.Stores, req *requests.RequestStore) error {
	modelStore.CompanyID = req.CompanyID
	modelStore.OwnerID = req.OwnerID
	modelStore.ShowStockLevelStatus = req.ShowStockLevelStatus
	modelStore.ShowOutOfStockStatus = req.ShowOutOfStockStatus
	modelStore.IsBackOrder = req.IsBackOrder
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms
	modelStore.Active = req.Active

	modelUser := models.Users{}
	service.DB.Table("users As u").
		Select(`
			u.email As contact_email,
			u.mobile_no As contact_phone
		`).
		Where("u.id = ?", req.OwnerID).
		Where("u.deleted_at Is Null").
		Scan(&modelUser)

	modelStore.ContactEmail = modelUser.ContactEmail
	modelStore.ContactPhone = modelUser.ContactPhone

	return service.DB.Create(&modelStore).Error
}
