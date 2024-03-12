package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	addrsvc "OnlineStoreBackend/services/customer_addresses"
	orditmsvc "OnlineStoreBackend/services/order_items"
)

func GetSalePrice(modelItem models.CartItemsWithDetail) float64 {
	price := modelItem.Price
	switch modelItem.DiscountType {
	case utils.FixedAmountOff:
		price -= modelItem.DiscountAmount
	case utils.PercentageOff:
		price -= modelItem.DiscountAmount * price / 100
	}
	return float64(price)
}

func GetShippingPrice(modelRates []models.ShippingTableRates, totalPrice float64, quantity float64, modelShip models.ShippingData) float64 {
	shippingPrice := float64(0)
	for _, modelRate := range modelRates {
		compare := float64(0)
		switch modelRate.Condition {
		case utils.Price:
			compare = totalPrice
		case utils.Weight:
			compare = modelShip.Weight
		case utils.ItemCount:
			compare = quantity
		}
		if compare >= modelRate.Min && compare <= modelRate.Max {
			shippingPrice += modelRate.CostPerKg*modelShip.Weight*quantity + modelRate.ItemCost*quantity + modelRate.RowCost + modelRate.PercentCost*totalPrice/100
		}
	}

	return shippingPrice
}

func (service *Service) Create(modelOrder *models.Orders, modelCartItems []models.CartItemsWithDetail, modelTax models.Taxes, customerID uint64) {
	modelAddr := models.Addresses{}

	addrRepo := repositories.NewRepositoryAddresses(service.DB)
	shipRepo := repositories.NewRepositoryShippingData(service.DB)
	methRepo := repositories.NewRepositoryShippingMethod(service.DB)
	orderService := orditmsvc.NewServiceOrderItem(service.DB)
	if err := addrRepo.ReadAddressByCustomerID(&modelAddr, customerID); err != nil {
		addrService := addrsvc.NewServiceAddress(service.DB)
		addrService.CreateFromUser(&modelAddr, customerID)
	}

	modelOrder.CustomerID = customerID
	modelOrder.BillingAddressID = uint64(modelAddr.ID)
	modelOrder.ShippingAddressID = uint64(modelAddr.ID)

	orderStatus := utils.StatusOrderPending

	modelItems := make([]models.OrderItems, 0)

	for _, modelItem := range modelCartItems {
		price := GetSalePrice(modelItem)
		totalPrice := price * modelItem.Quantity

		modelShip := models.ShippingData{}
		modelMethod := models.ShippingMethods{}
		modelRates := []models.ShippingTableRates{}

		methRepo.ReadRates(&modelRates, modelItem.StoreID)
		methRepo.ReadTableRateMethodByStoreID(&modelMethod, modelItem.StoreID)
		shipRepo.ReadByVariationID(&modelShip, modelItem.VariationID)

		shippingPrice := GetShippingPrice(modelRates, totalPrice, modelItem.Quantity, modelShip)

		itemStatus := utils.StatusOrderPending
		if modelItem.StockLevel < modelItem.Quantity {
			itemStatus = utils.StatusOrderBackOrdered
			orderStatus = utils.StatusOrderBackOrdered
		}

		taxAmount := modelTax.TaxRate * totalPrice / 100

		modelItems = append(modelItems, models.OrderItems{
			OrderID:          0,
			StoreID:          modelItem.StoreID,
			VariationID:      modelItem.VariationID,
			Price:            price,
			Quantity:         modelItem.Quantity,
			SubTotalPrice:    totalPrice,
			TaxRate:          modelTax.TaxRate,
			TaxAmount:        utils.Round(taxAmount),
			ShippingMethodID: uint64(modelMethod.ID),
			ShippingPrice:    shippingPrice,
			TotalPrice:       utils.Round(totalPrice + taxAmount + shippingPrice),
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
