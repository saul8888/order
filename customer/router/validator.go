package router

import "gopkg.in/go-playground/validator.v9"

type ValiDB interface {
	Validate(params interface{}) error
}

type Validb struct {
	validator *validator.Validate
}

func NewValiDB() ValiDB {
	return &Validb{validator: validator.New()}
}

func (v *Validb) Validate(params interface{}) error {
	return v.validator.Struct(params)
}
