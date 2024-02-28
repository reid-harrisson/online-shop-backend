package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	addrsvc "OnlineStoreBackend/services/customer_addresses"
	orditmsvc "OnlineStoreBackend/services/order_items"
)

func (service *Service) Create(modelOrder *models.Orders, modelCartItems []models.CartItemsWithDetail, modelTax models.TaxSettings, customerID uint64) {
	modelOrder.CustomerID = customerID
	modelOrder.Status = utils.StatusOrderPending

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
	for _, modelItem := range modelCartItems {
		price := modelItem.Price
		switch modelItem.DiscountType {
		case utils.FixedAmountOff:
			price -= modelItem.DiscountAmount
		case utils.PercentageOff:
			price -= modelItem.DiscountAmount * price / 100
		}
		totalPrice := price * modelItem.Quantity
		modelItems = append(modelItems, &models.OrderItems{
			OrderID:          uint64(orderID),
			StoreID:          modelItem.StoreID,
			VariationID:      modelItem.VariationID,
			Price:            price,
			Quantity:         modelItem.Quantity,
			SubTotalPrice:    totalPrice,
			TaxRate:          modelTax.TaxRate,
			TaxAmount:        utils.Round(modelTax.TaxRate * totalPrice / 100),
			ShippingMethodID: 0,
			ShippingPrice:    float64(0),
			TotalPrice:       utils.Round(totalPrice + (totalPrice * modelTax.TaxRate / 100)),
			Status:           utils.StatusOrderPending,
		})
	}
	orderService.Create(modelItems)
}
