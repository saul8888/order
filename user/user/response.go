package user

import "github.com/orderforme/user/model"

type UserList struct {
	Data         []model.User `json:"data"` //[]*model.Employee
	TotalRecords int          `json:"totalRecords"`
}

// DefaultLocationResponse body
type DefaultUserResponse struct {
	Error bool       `json:"error"`
	User  model.User `json:"location"`
}
