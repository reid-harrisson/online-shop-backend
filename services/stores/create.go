package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStore *models.Stores, req *requests.RequestStore) error {
	modelStore.CompanyID = req.CompanyID
	modelStore.OwnerID = req.OwnerID
	modelStore.ContactPhone = req.ContactPhone
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ShowStockLevelStatus = req.ShowStockLevelStatus
	modelStore.ShowOutOfStockStatus = req.ShowOutOfStockStatus
	modelStore.BackOrderStatus = req.BackOrderStatus
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms

	modelUser := models.Users{}
	service.DB.Table("users").
		Select(`
			users.email As contact_email,
			users.mobile_no As contact_phone
		`).
		Where("users.id = ?", req.OwnerID).
		Where("users.deleted_at Is Null").
		Scan(&modelUser)

	modelStore.ContactEmail = modelUser.ContactEmail
	modelStore.ContactPhone = modelUser.ContactPhone

	return service.DB.Create(&modelStore).Error
}
