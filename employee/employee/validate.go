package employee

import (
	"github.com/labstack/echo"
	"github.com/orderforme/employee/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.Employee) error
}

type validate struct {
	createEmployee model.Employee
	updateEmployee model.EmployeeUpdate
}

func CreateValidate(data model.Employee) Validate {
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
	if request.updateEmployee.FirstName == "" {
		model.Employeeupdate["firstName"] = customer.FirstName
	} else {
		model.Employeeupdate["firstName"] = request.updateEmployee.FirstName
	}

	if request.updateEmployee.LastName == "" {
		model.Employeeupdate["lastName"] = customer.LastName
	} else {
		model.Employeeupdate["lastName"] = request.updateEmployee.LastName
	}

	if request.updateEmployee.Email == "" {
		model.Employeeupdate["email"] = customer.Email
	} else {
		model.Employeeupdate["email"] = request.updateEmployee.Email
	}

	if request.updateEmployee.PhoneNumber == "" {
		model.Employeeupdate["phoneNumber"] = customer.PhoneNumber
	} else {
		model.Employeeupdate["phoneNumber"] = request.updateEmployee.PhoneNumber
	}

	if request.updateEmployee.Addresses == "" {
		model.Employeeupdate["addresses"] = customer.Addresses
	} else {
		model.Employeeupdate["addresses"] = request.updateEmployee.Addresses
	}

	return nil
}

/*
func (request *validate) Populate(customer model.Employee) error {
	fmt.Println(request.updateEmployee.FirstName)
	if request.updateEmployee.FirstName == "" {
		request.updateEmployee.FirstName = customer.FirstName
	}

	if request.updateEmployee.LastName == "" {
		request.updateEmployee.LastName = customer.FirstName
	}

	if request.updateEmployee.Email == "" {
		request.updateEmployee.Email = customer.Email
	}

	if request.updateEmployee.PhoneNumber == "" {
		request.updateEmployee.PhoneNumber = customer.PhoneNumber
	}

	if request.updateEmployee.Addresses == "" {
		request.updateEmployee.Addresses = customer.Addresses
	}

	return nil
}
*/
