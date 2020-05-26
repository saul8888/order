package payment

import (
	"github.com/labstack/echo"
	"github.com/orderforme/payment/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.Payment) error
}

type validate struct {
	createPayment model.CreatePayment
	updatePayment model.PaymentUpdate
}

func CreateValidate(data model.CreatePayment) Validate {
	return &validate{createPayment: data}
}

func UpdateValidate(data model.PaymentUpdate) Validate {
	return &validate{updatePayment: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createPayment); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updatePayment); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.Payment) error {
	if request.updatePayment.Name == "" {
		model.Paymentupdate["Name"] = customer.Name
	} else {
		model.Paymentupdate["Name"] = request.updatePayment.Name
	}

	if request.updatePayment.Status == "" {
		model.Paymentupdate["status"] = customer.Status
	} else {
		model.Paymentupdate["status"] = request.updatePayment.Status
	}

	if request.updatePayment.Currency == "" {
		model.Paymentupdate["currency"] = customer.Currency
	} else {
		model.Paymentupdate["currency"] = request.updatePayment.Currency
	}

	if request.updatePayment.Amount == "" {
		model.Paymentupdate["amount"] = customer.Amount
	} else {
		model.Paymentupdate["amount"] = request.updatePayment.Amount
	}

	return nil
}
