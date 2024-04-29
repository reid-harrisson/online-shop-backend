package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	addrsvc "OnlineStoreBackend/services/addresses"
	ordsvc "OnlineStoreBackend/services/orders"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HandlersCheckout struct {
	server *s.Server
}

func NewHandlersCheckout(server *s.Server) *HandlersCheckout {
	return &HandlersCheckout{server: server}
}

// Refresh godoc
// @Summary Create address to customer
// @Tags Checkout Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestAddress true "Address"
// @Success 201 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/checkout/address [post]
func (h *HandlersCheckout) CreateAddress(c echo.Context) error {
	req := new(requests.RequestAddress)

	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Create address
	modelAddr := models.Addresses{}
	addrService := addrsvc.NewServiceAddress(h.server.DB)
	err = addrService.Create(&modelAddr, req, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAddress(c, http.StatusCreated, modelAddr)
}

// Refresh godoc
// @Summary Read addresses
// @Tags Checkout Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/checkout/address [get]
func (h *HandlersCheckout) ReadAddresses(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read address by customer id
	modelAddrs := make([]models.Addresses, 0)
	addrRepo := repositories.NewRepositoryAddresses(h.server.DB)
	err = addrRepo.ReadAddressesByCustomerID(&modelAddrs, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAddresses(c, http.StatusOK, modelAddrs)
}

// Refresh godoc
// @Summary Read coupon
// @Tags Checkout Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param code query string true "Coupon code"
// @Success 200 {object} responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/checkout/coupon [get]
func (h *HandlersCheckout) ReadCoupon(c echo.Context) error {
	code := c.QueryParam("code")

	modelCoupon := models.Coupons{}

	// Read by code
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	err := couRepo.ReadByCode(&modelCoupon, code)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCoupon(c, http.StatusOK, modelCoupon)
}

// Refresh godoc
// @Summary Update address to customer
// @Tags Checkout Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Address ID"
// @Param params body requests.RequestAddress true "Customer Address"
// @Success 200 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/checkout/address/{id} [put]
func (h *HandlersCheckout) UpdateAddress(c echo.Context) error {
	req := new(requests.RequestAddress)

	addressID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelAddr := models.Addresses{}

	// Update address
	addrService := addrsvc.NewServiceAddress(h.server.DB)
	err = addrService.Update(&modelAddr, req, addressID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseAddress(c, http.StatusOK, modelAddr)
}

// Refresh godoc
// @Summary Read checkout
// @Tags Checkout Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 200 {object} []responses.ResponseCheckout
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout [post]
func (h *HandlersCheckout) Read(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestCheckout)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read detailed cart item
	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	err = cartRepo.ReadDetail(&modelItems, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read address by id
	modelAddr := models.Addresses{}
	addrRepo := repositories.NewRepositoryAddresses(h.server.DB)
	err = addrRepo.ReadByID(&modelAddr, req.ShippingAddressID, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read coupon by id
	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	err = couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	response, err := GetResponseStores(h.server.DB, modelItems, modelAddr, modelCoupons, models.Combos{})
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCheckout(c, http.StatusOK, response)
}

// Refresh godoc
// @Summary Read checkout from combo
// @Tags Checkout Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param combo_id query int true "Combo ID"
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 200 {object} []responses.ResponseCheckout
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/combo [post]
func (h *HandlersCheckout) ReadCombo(c echo.Context) error {
	req := new(requests.RequestCheckout)

	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	comboID, err := strconv.ParseUint(c.QueryParam("combo_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read combo by id
	modelCombo := models.Combos{}
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadByID(&modelCombo, comboID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read detailed cart item
	modelItems := make([]models.CartItemsWithDetail, 0)
	err = combRepo.ReadDetail(&modelItems, comboID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read address by id
	modelAddr := models.Addresses{}
	addrRepo := repositories.NewRepositoryAddresses(h.server.DB)
	err = addrRepo.ReadByID(&modelAddr, req.ShippingAddressID, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Read coupon by id
	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	err = couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	response, err := GetResponseStores(h.server.DB, modelItems, modelAddr, modelCoupons, modelCombo)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCheckout(c, http.StatusOK, response)
}

func GetResponseStores(db *gorm.DB, modelItems []models.CartItemsWithDetail, modelAddr models.Addresses, modelCoupons []models.Coupons, modelCombo models.Combos) ([]responses.ResponseCheckoutStore, error) {
	mapStore := map[uint64]int{}
	mapVar := map[uint64]int{}
	mapCoupon := map[uint64]int{}
	storeIDs := []uint64{}
	responseStores := []responses.ResponseCheckoutStore{}

	for index, modelCoupon := range modelCoupons {
		mapCoupon[modelCoupon.StoreID] = index
	}

	for index, modelItem := range modelItems {
		storeID := modelItem.StoreID

		if mapStore[storeID] == 0 {
			responseStores = append(responseStores, responses.ResponseCheckoutStore{
				StoreID: storeID,
			})
			mapStore[storeID] = len(responseStores)
		}

		storeIDs = append(storeIDs, modelItem.StoreID)
		mapVar[uint64(modelItem.ID)] = index
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
	taxRepo := repositories.NewRepositoryTax(db)
	err := taxRepo.ReadByCountryID(&modelTax, modelAddr.CountryID)
	if err != nil {
		return nil, err
	}

	mapRates := map[uint64][]models.ShippingTableRates{}
	mapMeth := map[uint64]models.ShippingMethods{}

	methRepo := repositories.NewRepositoryShippingMethod(db)
	err = methRepo.ReadMethodAndTableRatesByStoreIDs(&mapRates, &mapMeth, storeIDs)
	if err != nil {
		return nil, err
	}

	for index := range responseStores {
		subTotal := 0.0
		shippingPrice := 0.0
		quantity := 0.0
		shippingMethod := utils.ShippingMethodsToString(utils.TableRate)
		size := len(responseStores[index].Variations)
		couIndex := mapCoupon[responseStores[index].StoreID]

		for _, responseVar := range responseStores[index].Variations {
			totalPrice := responseVar.TotalPrice

			if modelCombo.ID != 0 {
				switch modelCombo.DiscountType {
				case utils.PercentageOff:
					totalPrice *= (100 - modelCombo.DiscountAmount) / 100
				case utils.FixedAmountOff:
					totalPrice -= modelCombo.DiscountAmount / float64(size)
				}
			}

			if len(modelCoupons) > 0 {
				switch modelCoupons[couIndex].DiscountType {
				case utils.FixedCartDiscount:
					totalPrice -= modelCoupons[couIndex].CouponAmount / float64(size)
				case utils.FixedProductDiscount:
					totalPrice -= modelCoupons[couIndex].CouponAmount * quantity
				case utils.PercentageDiscount:
					totalPrice *= (100 - modelCoupons[couIndex].CouponAmount) / 100
				}
			}

			subTotal += totalPrice
			shippingPrice += ordsvc.GetShippingPrice(mapRates[uint64(mapMeth[responseStores[index].StoreID].ID)], totalPrice, responseVar.Quantity, modelItems[mapVar[responseVar.ID]].Weight)
			quantity += responseVar.Quantity
		}
		if len(modelCoupons) > 0 {
			if modelCoupons[couIndex].AllowFreeShipping == 1 {
				shippingPrice = 0
				shippingMethod = utils.ShippingMethodsToString(utils.FreeShipping)
			}
		}
		taxAmount := modelTax.TaxRate * (subTotal + shippingPrice) / 100
		responseStores[index].ShippingPrice = utils.Round(shippingPrice)
		responseStores[index].SubTotal = utils.Round(subTotal)
		responseStores[index].TaxRate = utils.Round(modelTax.TaxRate)
		responseStores[index].TaxAmount = utils.Round(taxAmount)
		responseStores[index].TotalPrice = utils.Round(shippingPrice + subTotal + taxAmount)
		responseStores[index].ShippingMethod = shippingMethod
	}
	return responseStores, nil
}
