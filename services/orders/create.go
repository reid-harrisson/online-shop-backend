package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
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

func (service *Service) Create(modelOrder *models.Orders, modelCartItems []models.CartItemsWithDetail, billingAddressID uint64, shippingAddressID uint64, modelCoupons []models.Coupons, customerID uint64, modelCombo models.Combos) {
	mapCoupon := map[uint64]int{}
	for index, modelCoupon := range modelCoupons {
		mapCoupon[modelCoupon.StoreID] = index
	}

	modelAddr := models.Addresses{}
	addrRepo := repositories.NewRepositoryAddresses(service.DB)
	addrRepo.ReadAddressByID(&modelAddr, shippingAddressID)

	modelTax := models.Taxes{}
	taxRepo := repositories.NewRepositoryTax(service.DB)
	taxRepo.ReadByCountryID(&modelTax, modelAddr.CountryID)

	shipRepo := repositories.NewRepositoryShippingData(service.DB)
	methRepo := repositories.NewRepositoryShippingMethod(service.DB)
	orderService := orditmsvc.NewServiceOrderItem(service.DB)

	modelOrder.CustomerID = customerID
	modelOrder.BillingAddressID = billingAddressID
	modelOrder.ShippingAddressID = shippingAddressID

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

		if modelCombo.ID != 0 {
			switch modelCombo.DiscountType {
			case utils.PercentageOff:
				totalPrice *= (100 - modelCombo.DiscountAmount) / 100
			case utils.FixedAmountOff:
				totalPrice -= modelCombo.DiscountAmount / float64(len(modelCartItems))
			}
		}

		if len(modelCoupons) > 0 {
			couIndex := mapCoupon[modelItem.StoreID]
			switch modelCoupons[couIndex].DiscountType {
			case utils.FixedCartDiscount:
				totalPrice -= modelCoupons[couIndex].CouponAmount / float64(len(modelCartItems))
			case utils.FixedProductDiscount:
				totalPrice -= modelCoupons[couIndex].CouponAmount * modelItem.Quantity
			case utils.PercentageDiscount:
				totalPrice *= (100 - modelCoupons[couIndex].CouponAmount) / 100
			}
		}

		shippingPrice := GetShippingPrice(modelRates, totalPrice, modelItem.Quantity, modelShip)
		if len(modelCoupons) > 0 {
			couIndex := mapCoupon[modelItem.StoreID]
			if modelCoupons[couIndex].AllowFreeShipping == 1 {
				shippingPrice = 0
			}
		}
		itemStatus := utils.StatusOrderPending
		if modelItem.StockLevel < modelItem.Quantity {
			itemStatus = utils.StatusOrderBackOrdered
			orderStatus = utils.StatusOrderBackOrdered
		}

		taxAmount := modelTax.TaxRate * (totalPrice + shippingPrice) / 100

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
