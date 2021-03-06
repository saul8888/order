package employee

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/employee/database"
	"github.com/orderforme/employee/errors"
	"github.com/orderforme/employee/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee interface {
	Done()
	GetEmployeeByID(context echo.Context) error
	GetEmployees(context echo.Context) error
	CreateEmployee(context echo.Context) error
	UpdateEmployee(context echo.Context) error
	DeleteEmployee(context echo.Context) error
	SearchEmployee(context echo.Context) error
}

type Handler struct {
	employeeRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		employeeRepo: &database.Mongodb{},
	}

	err := handler.employeeRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.employeeRepo.DisconnectDB()
}

// GetEmployeeByID method
func (handler *Handler) GetEmployeeByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	employee, err := handler.employeeRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultEmployeeResponse{
		Error:    false,
		Employee: employee,
	})
}

// GetEmployees method
func (handler *Handler) GetEmployees(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		fmt.Println("Get limit")
		return context.JSON(http.StatusBadRequest, err)
	}

	employees, err := handler.employeeRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalEmployees, err := handler.employeeRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, EmployeeList{
		Data:         employees,
		TotalRecords: totalEmployees,
	})
}

// CreateEmployee method
func (handler *Handler) CreateEmployee(context echo.Context) error {
	request := new(model.Employee)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	//err := handler.employeeRepo.ValidateID("location", request.ID, model.Location{})
	err := handler.employeeRepo.ValidateID("location", request.LocationID, model.Employee{})
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err = handler.employeeRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultEmployeeResponse{
		Error:    false,
		Employee: *request,
	})

}

// UpdateCustomer method
func (handler *Handler) UpdateEmployee(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("customerId queryParam is missing"))
	}

	request := new(model.EmployeeUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	employee, err := handler.employeeRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(employee)
	updatedEmployee, err := handler.employeeRepo.Update(
		ID,
		model.Employeeupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultEmployeeResponse{
		Error:    false,
		Employee: updatedEmployee,
	})

}

// DeleteCustomer method
func (handler *Handler) DeleteEmployee(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("employeeId queryParam is missing"))
	}

	employee, err := handler.employeeRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.employeeRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultEmployeeResponse{
		Error:    false,
		Employee: employee,
	})

}

// SearchEmployee method
func (handler *Handler) SearchEmployee(context echo.Context) error {
	params := new(model.EmployeeSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.employeeRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	employees := []model.Employee{}
	if err = cursor.All(ctx, &employees); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, employees)

}
