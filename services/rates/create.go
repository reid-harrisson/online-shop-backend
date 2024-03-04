package ratesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelRate *models.ProductCustomerRates, req *requests.RequestProductRate) error {
	if err := service.DB.Where("product_id = ? And customer_id = ?", req.ProductID, req.CustomerID).First(modelRate).Error; err != nil {
		modelRate.CustomerID = req.CustomerID
		modelRate.ProductID = req.ProductID
		modelRate.Rate = req.Rate
		return service.DB.Create(modelRate).Error
	}
	modelRate.Rate = req.Rate
	return service.DB.Save(modelRate).Error
}
