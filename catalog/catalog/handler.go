package catalog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/catalog/database"
	"github.com/orderforme/catalog/errors"
	"github.com/orderforme/catalog/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catalog interface {
	Done()
	GetCatalogByID(context echo.Context) error
	GetCatalogs(context echo.Context) error
	CreateCatalog(context echo.Context) error
	UpdateCatalog(context echo.Context) error
	DeleteCatalog(context echo.Context) error
	AddCatalog(context echo.Context) error
	SearchCatalog(context echo.Context) error
}

type Handler struct {
	catalogRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		catalogRepo: &database.Mongodb{},
	}

	err := handler.catalogRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.catalogRepo.DisconnectDB()
}

// GetCatalogByID method
func (handler *Handler) GetCatalogByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	catalog, err := handler.catalogRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCatalogResponse{
		Error:   false,
		Catalog: catalog,
	})
}

// GetCatalogs method
func (handler *Handler) GetCatalogs(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	catalogs, err := handler.catalogRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalCatalogs, err := handler.catalogRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, CatalogList{
		Data:         catalogs,
		TotalRecords: totalCatalogs,
	})
}

// CreateCatalog method
func (handler *Handler) CreateCatalog(context echo.Context) error {
	request := new(model.Catalog)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	//validate location
	err := handler.catalogRepo.ValidateID("location", request.LocationID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	/*
		if request.Marcas == nil {
			request.Marcas.CreatedAt = time.Now()
			request.Marcas.UpdatedAt = time.Now()
		}
	*/
	err = handler.catalogRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCatalogResponse{
		Error:   false,
		Catalog: *request,
	})

}

// UpdateCatalog method
func (handler *Handler) UpdateCatalog(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("catalogId queryParam is missing"))
	}

	request := new(model.CatalogUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	/*
		catalog, err := handler.catalogRepo.GetByID(ID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, errors.NewError(err))
		}
	*/
	//req.Populate(catalog)
	updatedCatalog, err := handler.catalogRepo.Update(
		ID,
		model.Catalogupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCatalogResponse{
		Error:   false,
		Catalog: updatedCatalog,
	})

}

// DeleteCatalog
func (handler *Handler) DeleteCatalog(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("catalogId queryParam is missing"))
	}

	catalog, err := handler.catalogRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.catalogRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	//return context.JSON(http.StatusOK, catalog)
	return context.JSON(http.StatusOK, DefaultCatalogResponse{
		Error:   false,
		Catalog: catalog,
	})

}

// SearchCatalog method
func (handler *Handler) AddCatalog(context echo.Context) error {
	objectID, err := primitive.ObjectIDFromHex("5ed5831ba9de42de4423df36")
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	err = handler.catalogRepo.AddMarcas("hola", objectID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	return context.JSON(http.StatusOK, "hola")
}

// SearchCatalog method
func (handler *Handler) SearchCatalog(context echo.Context) error {
	params := new(model.CatalogSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.catalogRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	catalogs := []model.Catalog{}
	if err = cursor.All(ctx, &catalogs); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, catalogs)

}
