package merchant

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/merchant/database"
	"github.com/orderforme/merchant/errors"
	"github.com/orderforme/merchant/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Merchant interface {
	Done()
	GetMerchantByID(context echo.Context) error
	GetMerchants(context echo.Context) error
	CreateLocation(context echo.Context) error
	UpdateMerchant(context echo.Context) error
	DeleteMerchant(context echo.Context) error
	SearchMerchant(context echo.Context) error
}

type Handler struct {
	merchantRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		merchantRepo: &database.Mongodb{},
	}

	err := handler.merchantRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.merchantRepo.DisconnectDB()
}

// GetMerchantByID method
func (handler *Handler) GetMerchantByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	merchant, err := handler.merchantRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultMerchantResponse{
		Error:    false,
		Merchant: merchant,
	})
}

// GetMerchants method
func (handler *Handler) GetMerchants(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		fmt.Println("Get limit")
		return context.JSON(http.StatusBadRequest, err)
	}

	merchants, err := handler.merchantRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalMerchants, err := handler.merchantRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, MerchantList{
		Data:         merchants,
		TotalRecords: totalMerchants,
	})
}

// CreateMerchant method
func (handler *Handler) CreateMerchant(context echo.Context) error {
	request := new(model.Merchant)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err := handler.merchantRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultMerchantResponse{
		Error:    false,
		Merchant: *request,
	})

}

// UpdateMerchant method
func (handler *Handler) UpdateMerchant(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("merchantId queryParam is missing"))
	}

	request := new(model.MerchantUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	merchant, err := handler.merchantRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(merchant)
	updatedMerchant, err := handler.merchantRepo.Update(
		ID,
		model.Merchantupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultMerchantResponse{
		Error:    false,
		Merchant: updatedMerchant,
	})

}

// DeleteMerchant method
func (handler *Handler) DeleteMerchant(context echo.Context) error {

	ID := context.QueryParam("merchantId")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("merchantId queryParam is missing"))
	}

	merchant, err := handler.merchantRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.merchantRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	//return context.JSON(http.StatusOK, merchant)
	return context.JSON(http.StatusOK, DefaultMerchantResponse{
		Error:    false,
		Merchant: merchant,
	})

}

// SearchMerchant method
func (handler *Handler) SearchMerchant(context echo.Context) error {
	params := new(model.MerchantSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.merchantRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	merchants := []model.Merchant{}
	if err = cursor.All(ctx, &merchants); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, merchants)

}
