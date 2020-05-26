package merchant

import (
	"github.com/labstack/echo"
	"github.com/orderforme/merchant/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.Merchant) error
}

//Merchant merchant
type validate struct {
	createMerchant model.CreateMerchant
	updateMerchant model.MerchantUpdate
}

func CreateValidate(data model.CreateMerchant) Validate {
	return &validate{createMerchant: data}
}

func UpdateValidate(data model.MerchantUpdate) Validate {
	return &validate{updateMerchant: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createMerchant); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateMerchant); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.Merchant) error {
	if request.updateMerchant.Name == "" {
		model.Merchantupdate["Name"] = customer.Name
	} else {
		model.Merchantupdate["Name"] = request.updateMerchant.Name
	}

	if request.updateMerchant.Currency == "" {
		model.Merchantupdate["currency"] = customer.Currency
	} else {
		model.Merchantupdate["currency"] = request.updateMerchant.Currency
	}

	if request.updateMerchant.Status == "" {
		model.Merchantupdate["status"] = customer.Status
	} else {
		model.Merchantupdate["status"] = request.updateMerchant.Status
	}

	if request.updateMerchant.Addresses == "" {
		model.Merchantupdate["addresses"] = customer.Addresses
	} else {
		model.Merchantupdate["addresses"] = request.updateMerchant.Addresses
	}

	return nil
}
