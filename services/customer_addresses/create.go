package addrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelAddr *models.Addresses, req *requests.RequestAddress, customerID uint64) {
	modelAddr.CityID = req.CityID
	modelAddr.CountryID = req.CountryID
	modelAddr.RegionID = req.RegionID
	modelAddr.CustomerID = customerID
	modelAddr.PostalCode = req.PostalCode
	modelAddr.AddressLine1 = req.AddressLine1
	modelAddr.AddressLine2 = req.AddressLine2
	modelAddr.SubUrb = req.SubUrb
	modelAddr.Active = 1
	service.DB.Create(modelAddr)
}

func (service *Service) CreateFromUser(modelAddr *models.Addresses, customerID uint64) {
	service.DB.
		Table("users").
		Select(`
			users.address_line1 As address_line1,
			users.address_line2 As address_line2,
			users.suburb As suburb,
			users.id As customer_id,
			users.country_id As country_id,
			users.region_id As region_id,
			users.city_id As city_id,
			users.postal_code As postal_code
		`).
		Where("id = ?", customerID).
		Scan(modelAddr)
	modelAddr.Active = 1
	service.DB.Create(modelAddr)
}
