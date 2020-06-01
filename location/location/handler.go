package location

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/location/database"
	"github.com/orderforme/location/errors"
	"github.com/orderforme/location/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location interface {
	Done()
	GetLocationByID(context echo.Context) error
	GetLocations(context echo.Context) error
	CreateLocation(context echo.Context) error
	UpdateLocation(context echo.Context) error
	DeleteLocation(context echo.Context) error
	SearchLocation(context echo.Context) error
}

type Handler struct {
	locationRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		locationRepo: &database.Mongodb{},
	}

	err := handler.locationRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.locationRepo.DisconnectDB()
}

// GetLocationByID method
func (handler *Handler) GetLocationByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	location, err := handler.locationRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultLocationResponse{
		Error:    false,
		Location: location,
	})
}

// GetLocations method
func (handler *Handler) GetLocations(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	locations, err := handler.locationRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalLocations, err := handler.locationRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, LocationList{
		Data:         locations,
		TotalRecords: totalLocations,
	})
}

// CreateLocation method
func (handler *Handler) CreateLocation(context echo.Context) error {
	request := new(model.Location)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	//validate ID externt
	err := handler.locationRepo.ValidateID("merchant", request.MerchantID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err = handler.locationRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultLocationResponse{
		Error:    false,
		Location: *request,
	})

}

// UpdateLocation method
func (handler *Handler) UpdateLocation(context echo.Context) error {

	ID := context.QueryParam("locationId")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("locationId queryParam is missing"))
	}

	request := new(model.LocationUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	location, err := handler.locationRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(location)
	updatedLocation, err := handler.locationRepo.Update(
		ID,
		model.Locationupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultLocationResponse{
		Error:    false,
		Location: updatedLocation,
	})

}

// DeleteLocation
func (handler *Handler) DeleteLocation(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("locationId queryParam is missing"))
	}

	location, err := handler.locationRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.locationRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	//return context.JSON(http.StatusOK, location)
	return context.JSON(http.StatusOK, DefaultLocationResponse{
		Error:    false,
		Location: location,
	})

}

// SearchLocation method
func (handler *Handler) SearchLocation(context echo.Context) error {
	params := new(model.LocationSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.locationRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	locations := []model.Location{}
	if err = cursor.All(ctx, &locations); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, locations)

}
