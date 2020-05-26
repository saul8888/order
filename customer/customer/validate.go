package customer

import (
	"github.com/labstack/echo"
	"github.com/orderforme/customer/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.Customer) error
}

type validate struct {
	createCustomer model.CreateCustomer
	updateCustomer model.CustomerUpdate
}

func CreateValidate(data model.CreateCustomer) Validate {
	return &validate{createCustomer: data}
}

func UpdateValidate(data model.CustomerUpdate) Validate {
	return &validate{updateCustomer: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createCustomer); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateCustomer); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.Customer) error {
	if request.updateCustomer.FirstName == "" {
		model.Customerupdate["firstName"] = customer.FirstName
	} else {
		model.Customerupdate["firstName"] = request.updateCustomer.FirstName
	}

	if request.updateCustomer.LastName == "" {
		model.Customerupdate["lastName"] = customer.LastName
	} else {
		model.Customerupdate["lastName"] = request.updateCustomer.LastName
	}

	if request.updateCustomer.Email == "" {
		model.Customerupdate["email"] = customer.Email
	} else {
		model.Customerupdate["email"] = request.updateCustomer.Email
	}

	if request.updateCustomer.PhoneNumber == "" {
		model.Customerupdate["phoneNumber"] = customer.PhoneNumber
	} else {
		model.Customerupdate["phoneNumber"] = request.updateCustomer.PhoneNumber
	}

	if request.updateCustomer.Addresses == "" {
		model.Customerupdate["addresses"] = customer.Addresses
	} else {
		model.Customerupdate["addresses"] = request.updateCustomer.Addresses
	}

	return nil
}
