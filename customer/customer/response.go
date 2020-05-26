package customer

import "github.com/orderforme/customer/model"

type CustomerList struct {
	Data         []model.Customer `json:"data"` //[]*model.Employee
	TotalRecords int              `json:"totalRecords"`
}

// DefaultCustomerResponse body
type DefaultCustomerResponse struct {
	Error    bool           `json:"error"`
	Customer model.Customer `json:"customer"`
}
