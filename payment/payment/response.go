package payment

import "github.com/orderforme/payment/model"

type PaymentList struct {
	Data         []model.Payment `json:"data"` //[]*model.Employee
	TotalRecords int             `json:"totalRecords"`
}

// DefaultPaymentResponse body
type DefaultPaymentResponse struct {
	Error   bool          `json:"error"`
	Payment model.Payment `json:"payment"`
}
