package utils

import (
	"OnlineStoreBackend/pkgs/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvokeData struct {
	CardNumber   string      `json:"card_number"`
	ExpMonth     int64       `json:"exp_month"`
	ExpYear      int64       `json:"exp_year"`
	CVC          string      `json:"cvc"`
	CustomerName string      `json:"customer_name"`
	Country      string      `json:"country"`
	ZipCode      string      `json:"zip_code"`
	Amount       float64     `json:"amount"`
	Currency     string      `json:"currency"`
	PaymentType  PaymentType `json:"payment_type"`
	RequestID    uint64      `json:"request_id"`
}

func HelperInvoke(paymentMethod string, requestUrl string, c echo.Context, invokeData InvokeData) {
	logger, err := logging.NewLogger()

	if err != nil {
		panic(err)
	}

	jsonStr, err := json.Marshal(invokeData)
	if err != nil {
		logger.Error(err.Error())
	}

	responseInvoke, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		logger.Error(err.Error())
	}

	defer responseInvoke.Body.Close()

	if !(responseInvoke.StatusCode == http.StatusOK || responseInvoke.StatusCode == http.StatusCreated) {
		errMsg := fmt.Sprintf("Audit service error: %d %s", responseInvoke.StatusCode, responseInvoke.Status)
		logger.Error(errMsg)
	}
}
