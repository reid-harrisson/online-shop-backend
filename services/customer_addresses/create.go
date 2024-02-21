package addrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelAddr *models.CustomerAddresses, req *requests.RequestCustomerAddress, customerID uint64) {
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
