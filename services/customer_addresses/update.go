package addrsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelAddr *models.Addresses, req *requests.RequestAddress, addressID uint64) {
	service.DB.First(modelAddr, addressID)
	modelAddr.Active = 0
	service.DB.Save(modelAddr)
	modelAddr.CityID = req.CityID
	modelAddr.CountryID = req.CountryID
	modelAddr.RegionID = req.RegionID
	modelAddr.PostalCode = req.PostalCode
	modelAddr.AddressLine1 = req.AddressLine1
	modelAddr.AddressLine2 = req.AddressLine2
	modelAddr.SubUrb = req.SubUrb
	modelAddr.Active = 1
	modelAddr.ID = 0
	service.DB.Create(modelAddr)
}
