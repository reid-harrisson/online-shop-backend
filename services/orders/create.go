package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	addrsvc "OnlineStoreBackend/services/customer_addresses"
	orditmsvc "OnlineStoreBackend/services/order_items"
)

func (service *Service) Create(modelOrder *models.Orders, modelCarts []models.CartItemsWithDetail, modelTax models.TaxSettings, customerID uint64) {
	modelOrder.CustomerID = customerID
	modelOrder.Status = models.StatusOrderPending

	modelAddr := models.CustomerAddresses{}
	addrRepo := repositories.NewRepositoryCustomer(service.DB)
	addrRepo.ReadAddressByCustomerID(&modelAddr, customerID)

	if modelAddr.ID == 0 {
		addrService := addrsvc.NewServiceCustomerAddress(service.DB)
		addrService.CreateFromUser(&modelAddr, customerID)
	}

	modelOrder.BillingAddressID = uint64(modelAddr.ID)
	modelOrder.ShippingAddressID = uint64(modelAddr.ID)
	service.DB.Create(&modelOrder)

	orderID := modelOrder.ID
	orderService := orditmsvc.NewServiceOrderItem(service.DB)
	modelItems := make([]*models.OrderItems, 0)
	for _, modelCart := range modelCarts {
		modelItems = append(modelItems, &models.OrderItems{
			OrderID:          uint64(orderID),
			StoreID:          modelCart.StoreID,
			VariationID:      modelCart.ProductID,
			Price:            modelCart.UnitPrice,
			Quantity:         modelCart.Quantity,
			SubTotalPrice:    modelCart.TotalPrice,
			TaxRate:          modelTax.TaxRate,
			TaxAmount:        utils.Round(modelTax.TaxRate * modelCart.TotalPrice / 100),
			ShippingMethodID: 0,
			ShippingPrice:    float64(0),
			TotalPrice:       utils.Round(modelCart.TotalPrice + (modelCart.TotalPrice * modelTax.TaxRate / 100)),
			Status:           models.StatusOrderPending,
		})
	}
	orderService.Create(modelItems)
}
