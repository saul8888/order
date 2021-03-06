package payment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/payment/database"
	"github.com/orderforme/payment/errors"
	"github.com/orderforme/payment/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//payment Payment
type Payment interface {
	Done()
	GetPaymentByID(context echo.Context) error
	GetPayments(context echo.Context) error
	CreatePayment(context echo.Context) error
	UpdatePayment(context echo.Context) error
	DeletePayment(context echo.Context) error
	SearchPayment(context echo.Context) error
}

type Handler struct {
	paymentRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		paymentRepo: &database.Mongodb{},
	}

	err := handler.paymentRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.paymentRepo.DisconnectDB()
}

// GetPaymentByID
func (handler *Handler) GetPaymentByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	payment, err := handler.paymentRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultPaymentResponse{
		Error:   false,
		Payment: payment,
	})
}

// GetPayments method
func (handler *Handler) GetPayments(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		fmt.Println("Get limit")
		return context.JSON(http.StatusBadRequest, err)
	}

	payments, err := handler.paymentRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalPayments, err := handler.paymentRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, PaymentList{
		Data:         payments,
		TotalRecords: totalPayments,
	})
}

// CreatePayment method
func (handler *Handler) CreatePayment(context echo.Context) error {
	request := new(model.Payment)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	//validate ID externt
	err := handler.paymentRepo.ValidateID("location", request.LocationID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	//validate ID externt
	err = handler.paymentRepo.ValidateID("customer", request.CustomerID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err = handler.paymentRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultPaymentResponse{
		Error:   false,
		Payment: *request,
	})

}

// UpdatePayment
func (handler *Handler) UpdatePayment(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("paymentId queryParam is missing"))
	}

	request := new(model.PaymentUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	payment, err := handler.paymentRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(payment)
	updatedPayment, err := handler.paymentRepo.Update(
		ID,
		model.Paymentupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultPaymentResponse{
		Error:   false,
		Payment: updatedPayment,
	})

}

// DeletePayment method
func (handler *Handler) DeletePayment(context echo.Context) error {

	paymentID := context.QueryParam("paymentId")

	if len(paymentID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("paymentId queryParam is missing"))
	}

	payment, err := handler.paymentRepo.GetByID(paymentID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.paymentRepo.Delete(paymentID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	//return context.JSON(http.StatusOK, payment)
	return context.JSON(http.StatusOK, DefaultPaymentResponse{
		Error:   false,
		Payment: payment,
	})

}

// SearchPayment method
func (handler *Handler) SearchPayment(context echo.Context) error {
	params := new(model.PaymentSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.paymentRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	payments := []model.Payment{}
	if err = cursor.All(ctx, &payments); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, payments)

}
