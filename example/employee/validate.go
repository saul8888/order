package employee

import (
	"github.com/labstack/echo"
	"github.com/orderforme/example/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	//ValiCreate(context echo.Context) (model.Employee, error)
	ValiUpdate(context echo.Context) error
	Populate(customer model.Employee) error
}

type validate struct {
	createEmployee model.CreateEmployee
	updateEmployee model.EmployeeUpdate
}

func CreateValidate(data model.CreateEmployee) Validate {
	return &validate{createEmployee: data}
}

func UpdateValidate(data model.EmployeeUpdate) Validate {
	return &validate{updateEmployee: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createEmployee); err != nil {
		return err
	}

	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateEmployee); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.Employee) error {

	if request.updateEmployee.Status == "" {
		model.Employeeupdate["status"] = customer.Status
	} else {
		model.Employeeupdate["status"] = request.updateEmployee.Status
	}

	if request.updateEmployee.Email == "" {
		model.Employeeupdate["email"] = customer.Email
	} else {
		model.Employeeupdate["email"] = request.updateEmployee.Email
	}

	if request.updateEmployee.Addresses.Country == "" {
		model.Employeeupdate["addresses.country"] = customer.Addresses.Country
	} else {
		model.Employeeupdate["addresses.country"] = request.updateEmployee.Addresses.Country
	}

	return nil
}
