package location

import (
	"github.com/labstack/echo"
	"github.com/orderforme/location/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.Location) error
}

type validate struct {
	createLocation model.CreateLocation
	updateLocation model.LocationUpdate
}

func CreateValidate(data model.CreateLocation) Validate {
	return &validate{createLocation: data}
}

func UpdateValidate(data model.LocationUpdate) Validate {
	return &validate{updateLocation: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createLocation); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateLocation); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.Location) error {
	if request.updateLocation.Name == "" {
		model.Locationupdate["Name"] = customer.Name
	} else {
		model.Locationupdate["Name"] = request.updateLocation.Name
	}

	if request.updateLocation.Currency == "" {
		model.Locationupdate["currency"] = customer.Currency
	} else {
		model.Locationupdate["currency"] = request.updateLocation.Currency
	}

	if request.updateLocation.Email == "" {
		model.Locationupdate["email"] = customer.Email
	} else {
		model.Locationupdate["email"] = request.updateLocation.Email
	}

	if request.updateLocation.PhoneNumber == "" {
		model.Locationupdate["phoneNumber"] = customer.PhoneNumber
	} else {
		model.Locationupdate["phoneNumber"] = request.updateLocation.PhoneNumber
	}

	if request.updateLocation.Instagram == "" {
		model.Locationupdate["instagram"] = customer.Instagram
	} else {
		model.Locationupdate["instagram"] = request.updateLocation.Instagram
	}

	return nil
}
