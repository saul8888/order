package employee

import "github.com/orderforme/example/model"

type EmployeeList struct {
	Data         []model.Employee `json:"data"` //[]*model.Employee
	TotalRecords int              `json:"totalRecords"`
}

// DefaultEmployeeResponse body
type DefaultEmployeeResponse struct {
	Error    bool           `json:"error"`
	Employee model.Employee `json:"employee"`
}
