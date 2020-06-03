package catalog

import "github.com/orderforme/catalog/model"

type CatalogList struct {
	Data         []model.Catalog `json:"data"` //[]*model.Employee
	TotalRecords int             `json:"totalRecords"`
}

// DefaultCatalogResponse body
type DefaultCatalogResponse struct {
	Error   bool          `json:"error"`
	Catalog model.Catalog `json:"catalog"`
}
