package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	addrsvc "OnlineStoreBackend/services/customer_addresses"
	ordsvc "OnlineStoreBackend/services/orders"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCheckoutProcess struct {
	server *s.Server
}

func NewHandlersCheckoutProcess(server *s.Server) *HandlersCheckoutProcess {
	return &HandlersCheckoutProcess{server: server}
}

// Refresh godoc
// @Summary Create address to customer
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param params body requests.RequestAddress true "Address"
// @Success 201 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/address [post]
func (h *HandlersCheckoutProcess) CreateAddress(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	req := new(requests.RequestAddress)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAddr := models.Addresses{}
	addrService := addrsvc.NewServiceAddress(h.server.DB)
	addrService.Create(&modelAddr, req, customerID)

	return responses.NewResponseAddress(c, http.StatusCreated, modelAddr)
}

// Refresh godoc
// @Summary Read addresses
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/address [get]
func (h *HandlersCheckoutProcess) ReadAddresses(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	modelAddrs := make([]models.Addresses, 0)
	addrRepo := repositories.NewRepositoryAddresses(h.server.DB)
	addrRepo.ReadAddressesByCustomerID(&modelAddrs, customerID)

	return responses.NewResponseAddresses(c, http.StatusOK, modelAddrs)
}

// Refresh godoc
// @Summary Read coupon
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param code query string true "Coupon code"
// @Success 200 {object} responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/coupon [get]
func (h *HandlersCheckoutProcess) ReadCoupon(c echo.Context) error {
	code := c.QueryParam("code")

	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	couRepo.ReadByCode(&modelCoupon, code)

	return responses.NewResponseCoupon(c, http.StatusOK, modelCoupon)
}

// Refresh godoc
// @Summary Update address to customer
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Address ID"
// @Param params body requests.RequestAddress true "Customer Address"
// @Success 200 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/address/{id} [put]
func (h *HandlersCheckoutProcess) UpdateAddress(c echo.Context) error {
	addressID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestAddress)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAddr := models.Addresses{}
	addrService := addrsvc.NewServiceAddress(h.server.DB)
	addrService.Update(&modelAddr, req, addressID)

	return responses.NewResponseAddress(c, http.StatusOK, modelAddr)
}

// Refresh godoc
// @Summary Read checkout
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 200 {object} []responses.ResponseCheckout
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout [post]
func (h *HandlersCheckoutProcess) Read(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	req := new(requests.RequestCheckout)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadDetail(&modelItems, customerID)

	modelAddr := models.Addresses{}
	addrRepo := repositories.NewRepositoryAddresses(h.server.DB)
	addrRepo.ReadAddressByID(&modelAddr, req.ShippingAddressID)

	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)

	mapStore := map[uint64]int{}
	mapCoupon := map[uint64]int{}
	responseStores := []responses.ResponseCheckoutStore{}
	for index, modelCoupon := range modelCoupons {
		mapCoupon[modelCoupon.StoreID] = index
	}
	for _, modelItem := range modelItems {
		storeID := modelItem.StoreID
		if mapStore[storeID] == 0 {
			responseStores = append(responseStores, responses.ResponseCheckoutStore{
				StoreID: storeID,
			})
			mapStore[storeID] = len(responseStores)
		}
		storeIndex := mapStore[storeID] - 1
		imageUrls := make([]string, 0)
		json.Unmarshal([]byte(modelItem.ImageUrls), &imageUrls)
		categories := make([]string, 0)
		json.Unmarshal([]byte("["+modelItem.Categories+"]"), &categories)
		salePrice := ordsvc.GetSalePrice(modelItem)
		responseStores[storeIndex].Variations = append(responseStores[storeIndex].Variations, responses.ResponseCheckoutVariation{
			ID:            uint64(modelItem.ID),
			VariationID:   modelItem.VariationID,
			VariationName: modelItem.VariationName,
			ImageUrls:     imageUrls,
			Categories:    categories,
			SalePrice:     salePrice,
			RegularPrice:  modelItem.Price,
			Quantity:      modelItem.Quantity,
			StockLevel:    modelItem.StockLevel,
			TotalPrice:    modelItem.Quantity * salePrice,
		})
	}
	modelTax := models.Taxes{}
	taxRepo := repositories.NewRepositoryTax(h.server.DB)
	taxRepo.ReadByCountryID(&modelTax, modelAddr.CountryID)
	tableRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	for index := range responseStores {
		modelRates := []models.ShippingTableRates{}
		tableRepo.ReadRates(&modelRates, responseStores[index].StoreID)
		subTotal := 0.0
		shippingPrice := 0.0
		quantity := 0.0

		for _, responseVar := range responseStores[index].Variations {
			modelShip := models.ShippingData{}
			shipRepo := repositories.NewRepositoryShippingData(h.server.DB)
			shipRepo.ReadByVariationID(&modelShip, responseVar.VariationID)

			subTotal += responseVar.TotalPrice
			shippingPrice += ordsvc.GetShippingPrice(modelRates, responseVar.TotalPrice, responseVar.Quantity, modelShip)
			quantity += responseVar.Quantity
		}
		if len(modelCoupons) > 0 {
			couIndex := mapCoupon[responseStores[index].StoreID]
			if modelCoupons[couIndex].AllowFreeShipping == 1 {
				shippingPrice = 0
			}
			switch modelCoupons[couIndex].DiscountType {
			case utils.FixedCartDiscount:
				subTotal -= modelCoupons[couIndex].CouponAmount
			case utils.FixedProductDiscount:
				subTotal -= modelCoupons[couIndex].CouponAmount * quantity
			case utils.PercentageDiscount:
				subTotal *= (100 - modelCoupons[couIndex].CouponAmount) / 100
			}
		}
		taxAmount := modelTax.TaxRate * (subTotal + shippingPrice) / 100
		responseStores[index].ShippingPrice = utils.Round(shippingPrice)
		responseStores[index].SubTotal = utils.Round(subTotal)
		responseStores[index].TaxRate = utils.Round(modelTax.TaxRate)
		responseStores[index].TaxAmount = utils.Round(taxAmount)
		responseStores[index].TotalPrice = utils.Round(shippingPrice + subTotal + taxAmount)
	}

	return responses.NewResponseCheckout(c, http.StatusOK, responseStores)
}
