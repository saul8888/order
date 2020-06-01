package user

import (
	"github.com/labstack/echo"
	"github.com/orderforme/user/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.User) error
}

type validate struct {
	createUser model.User
	updateUser model.UserUpdate
}

func CreateValidate(data model.User) Validate {
	return &validate{createUser: data}
}

func UpdateValidate(data model.UserUpdate) Validate {
	return &validate{updateUser: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createUser); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateUser); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.User) error {
	if request.updateUser.UserName == "" {
		model.Userupdate["userName"] = customer.UserName
	} else {
		model.Userupdate["userName"] = request.updateUser.UserName
	}

	if request.updateUser.UserCreate == "" {
		model.Userupdate["userCreate"] = customer.UserCreate
	} else {
		model.Userupdate["userCreate"] = request.updateUser.UserCreate
	}

	if request.updateUser.UserSuper == "" {
		model.Userupdate["userSuper"] = customer.UserSuper
	} else {
		model.Userupdate["userSuper"] = request.updateUser.UserSuper
	}

	if request.updateUser.UserCatalogs == "" {
		model.Userupdate["userCatalogs"] = customer.UserCatalogs
	} else {
		model.Userupdate["userCatalogs"] = request.updateUser.UserCatalogs
	}

	if request.updateUser.Valuntil == "" {
		model.Userupdate["valuntil"] = customer.Valuntil
	} else {
		model.Userupdate["valuntil"] = request.updateUser.Valuntil
	}

	if request.updateUser.UserLimit == "" {
		model.Userupdate["userLimit"] = customer.UserLimit
	} else {
		model.Userupdate["userLimit"] = request.updateUser.UserLimit
	}

	if request.updateUser.Status == "" {
		model.Userupdate["status"] = customer.Status
	} else {
		model.Userupdate["status"] = request.updateUser.Status
	}

	return nil
}
