package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	addrsvc "OnlineStoreBackend/services/customer_addresses"
	orditmsvc "OnlineStoreBackend/services/order_items"
)

func (service *Service) Create(modelOrder *models.Orders, modelCartItems []models.CartItemsWithDetail, modelTax models.TaxSettings, customerID uint64) {
	modelAddr := models.CustomerAddresses{}

	addrRepo := repositories.NewRepositoryCustomer(service.DB)
	shipRepo := repositories.NewRepositoryShippingData(service.DB)
	methRepo := repositories.NewRepositoryShippingMethod(service.DB)
	orderService := orditmsvc.NewServiceOrderItem(service.DB)
	if err := addrRepo.ReadAddressByCustomerID(&modelAddr, customerID); err != nil {
		addrService := addrsvc.NewServiceCustomerAddress(service.DB)
		addrService.CreateFromUser(&modelAddr, customerID)
	}

	modelOrder.CustomerID = customerID
	modelOrder.BillingAddressID = uint64(modelAddr.ID)
	modelOrder.ShippingAddressID = uint64(modelAddr.ID)

	orderStatus := utils.StatusOrderPending

	modelItems := make([]models.OrderItems, 0)

	for _, modelItem := range modelCartItems {
		price := modelItem.Price
		switch modelItem.DiscountType {
		case utils.FixedAmountOff:
			price -= modelItem.DiscountAmount
		case utils.PercentageOff:
			price -= modelItem.DiscountAmount * price / 100
		}
		totalPrice := price * modelItem.Quantity

		modelShip := models.ShippingData{}
		modelMethod := models.ShippingMethods{}

		methRepo.ReadDefault(&modelMethod, modelItem.StoreID)
		shipRepo.ReadByVariationID(&modelShip, modelItem.VariationID)

		shippingPrice := float64(0)

		if modelMethod.ID != 0 {
			switch modelMethod.Method {
			case utils.FlatRate:
				shippingPrice = 0
			case utils.TableRate:
				shippingPrice = 0
			}
		}

		itemStatus := utils.StatusOrderPending
		if modelItem.StockLevel < modelItem.Quantity {
			itemStatus = utils.StatusOrderBackOrdered
			orderStatus = utils.StatusOrderBackOrdered
		}

		modelItems = append(modelItems, models.OrderItems{
			OrderID:          0,
			StoreID:          modelItem.StoreID,
			VariationID:      modelItem.VariationID,
			Price:            price,
			Quantity:         modelItem.Quantity,
			SubTotalPrice:    totalPrice,
			TaxRate:          modelTax.TaxRate,
			TaxAmount:        utils.Round(modelTax.TaxRate * totalPrice / 100),
			ShippingMethodID: uint64(modelMethod.ID),
			ShippingPrice:    shippingPrice,
			TotalPrice:       utils.Round(totalPrice + (totalPrice * modelTax.TaxRate / 100)),
			Status:           itemStatus,
		})
	}
	modelOrder.Status = orderStatus
	service.DB.Create(&modelOrder)
	for index := range modelItems {
		modelItems[index].OrderID = uint64(modelOrder.ID)
	}
	orderService.Create(&modelItems)
}
