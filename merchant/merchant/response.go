package merchant

import "github.com/orderforme/merchant/model"

type MerchantList struct {
	Data         []model.Merchant `json:"data"` //[]*model.Employee
	TotalRecords int              `json:"totalRecords"`
}

// DefaultMerchantResponse body
type DefaultMerchantResponse struct {
	Error    bool           `json:"error"`
	Merchant model.Merchant `json:"merchant"`
}
