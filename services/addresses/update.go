package addrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelAddr *models.Addresses, req *requests.RequestAddress, addressID uint64) error {
	if err := service.DB.Model(&models.Addresses{}).Where("id = ?", addressID).Update("active", 0).Error; err != nil {
		return err
	}
	customerID := uint64(0)
	if err := service.DB.Model(&models.Addresses{}).Select("customer_id").Where("id = ?", addressID).Scan(&customerID).Error; err != nil {
		return err
	}
	modelAddr.Name = req.Name
	modelAddr.CustomerID = customerID
	modelAddr.CountryID = req.CountryID
	modelAddr.RegionID = req.RegionID
	modelAddr.CityID = req.CityID
	modelAddr.PostalCode = req.PostalCode
	modelAddr.AddressLine1 = req.AddressLine1
	modelAddr.AddressLine2 = req.AddressLine2
	modelAddr.SubUrb = req.SubUrb
	modelAddr.Active = 1
	modelAddr.ID = 0
	return service.DB.Create(modelAddr).Error
}
