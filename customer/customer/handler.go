package customer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/customer/database"
	"github.com/orderforme/customer/errors"
	"github.com/orderforme/customer/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer interface {
	Done()
	GetCustomerByID(context echo.Context) error
	GetCustomers(context echo.Context) error
	CreateCustomer(context echo.Context) error
	UpdateCustomer(context echo.Context) error
	DeleteCustomer(context echo.Context) error
	SearchCustomer(context echo.Context) error
}

type Handler struct {
	customerRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		customerRepo: &database.Mongodb{},
	}

	err := handler.customerRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.customerRepo.DisconnectDB()
}

// GetEmployeeByID method
func (handler *Handler) GetCustomerByID(context echo.Context) error {

	customerID := context.QueryParam("Id")
	find, err := handler.customerRepo.GetByID(customerID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	customer := model.Customer{}
	if err = find.Decode(&customer); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCustomerResponse{
		Error:    false,
		Customer: customer,
	})
}

// GetCustomers method
func (handler *Handler) GetCustomers(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		fmt.Println("Get limit")
		return context.JSON(http.StatusBadRequest, err)
	}

	cursor, ctx, err := handler.customerRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	customers := []model.Customer{}
	if err = cursor.All(ctx, &customers); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalCustomers, err := handler.customerRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, CustomerList{
		Data:         customers,
		TotalRecords: totalCustomers,
	})
}

// CreateCustomer method
func (handler *Handler) CreateCustomer(context echo.Context) error {
	request := new(model.CreateCustomer)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	// Dates Mongodb
	request.CustomerID = primitive.NewObjectID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	_, err := handler.customerRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	//show new customer
	find, err := handler.customerRepo.GetByID(request.CustomerID.Hex())
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	customer := model.Customer{}
	if err = find.Decode(&customer); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCustomerResponse{
		Error:    false,
		Customer: customer,
	})

}

//  UpdateCustomer method
func (handler *Handler) UpdateCustomer(context echo.Context) error {

	customerID := context.QueryParam("customerId")
	if len(customerID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("customerId queryParam is missing"))
	}

	request := new(model.CustomerUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	find, err := handler.customerRepo.GetByID(customerID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	currentCustomer := model.Customer{}
	if err = find.Decode(&currentCustomer); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(currentCustomer)
	updatedCustomer, err := handler.customerRepo.Update(
		customerID,
		model.Customerupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCustomerResponse{
		Error:    false,
		Customer: updatedCustomer,
	})

}

// DeleteCustomer method
func (handler *Handler) DeleteCustomer(context echo.Context) error {

	customerID := context.QueryParam("customerId")

	if len(customerID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("customerId queryParam is missing"))
	}

	find, err := handler.customerRepo.GetByID(customerID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	customer := model.Customer{}
	if err = find.Decode(&customer); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.customerRepo.Delete(customerID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultCustomerResponse{
		Error:    false,
		Customer: customer,
	})

}

// SearchCustomer
func (handler *Handler) SearchCustomer(context echo.Context) error {
	params := new(model.CustomerSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.customerRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	customers := []model.Customer{}
	if err = cursor.All(ctx, &customers); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, customers)

}
