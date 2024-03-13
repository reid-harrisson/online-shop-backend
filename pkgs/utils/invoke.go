package utils

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
