package location

import "github.com/orderforme/location/model"

type LocationList struct {
	Data         []model.Location `json:"data"` //[]*model.Employee
	TotalRecords int              `json:"totalRecords"`
}

// DefaultLocationResponse body
type DefaultLocationResponse struct {
	Error    bool           `json:"error"`
	Location model.Location `json:"location"`
}
