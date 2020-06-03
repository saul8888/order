package catalog

import (
	"github.com/labstack/echo"
	"github.com/orderforme/catalog/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	//Populate(catalog model.Catalog) error
}

type validate struct {
	createCatalog model.Catalog
	updateCatalog model.CatalogUpdate
}

func CreateValidate(data model.Catalog) Validate {
	return &validate{createCatalog: data}
}

func UpdateValidate(data model.CatalogUpdate) Validate {
	return &validate{updateCatalog: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createCatalog); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateCatalog); err != nil {
		return err
	}
	return nil
}

/*
func (request *validate) Populate(catalog model.Catalog) error {
	if request.updateCatalog.Name == "" {
		model.Catalogupdate["Name"] = catalog.Name
	} else {
		model.Catalogupdate["Name"] = request.updateCatalog.Name
	}

	if request.updateCatalog.Currency == "" {
		model.Catalogupdate["currency"] = catalog.Currency
	} else {
		model.Catalogupdate["currency"] = request.updateCatalog.Currency
	}

	if request.updateCatalog.Email == "" {
		model.Catalogupdate["email"] = catalog.Email
	} else {
		model.Catalogupdate["email"] = request.updateCatalog.Email
	}

	if request.updateCatalog.PhoneNumber == "" {
		model.Catalogupdate["phoneNumber"] = catalog.PhoneNumber
	} else {
		model.Catalogupdate["phoneNumber"] = request.updateCatalog.PhoneNumber
	}

	if request.updateCatalog.Instagram == "" {
		model.Catalogupdate["instagram"] = catalog.Instagram
	} else {
		model.Catalogupdate["instagram"] = request.updateCatalog.Instagram
	}

	if request.updateCatalog.Status == "" {
		model.Catalogupdate["status"] = catalog.Status
	} else {
		model.Catalogupdate["status"] = request.updateCatalog.Status
	}

	return nil
}
*/
