package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orderforme/user/database"
	"github.com/orderforme/user/errors"
	"github.com/orderforme/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	Done()
	GetUserByID(context echo.Context) error
	GetUsers(context echo.Context) error
	CreateUser(context echo.Context) error
	UpdateUser(context echo.Context) error
	DeleteUser(context echo.Context) error
	SearchUser(context echo.Context) error
}

type Handler struct {
	userRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		userRepo: &database.Mongodb{},
	}

	err := handler.userRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.userRepo.DisconnectDB()
}

// GetUserByID method
func (handler *Handler) GetUserByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	user, err := handler.userRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  user,
	})
}

// GetUsers method
func (handler *Handler) GetUsers(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		fmt.Println("Get limit")
		return context.JSON(http.StatusBadRequest, err)
	}

	users, err := handler.userRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalLocations, err := handler.userRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, UserList{
		Data:         users,
		TotalRecords: totalLocations,
	})
}

// CreateUser method
func (handler *Handler) CreateUser(context echo.Context) error {
	request := new(model.User)
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

	err := handler.userRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  *request,
	})

}

// UpdateUser method
func (handler *Handler) UpdateUser(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("userId queryParam is missing"))
	}

	request := new(model.UserUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	user, err := handler.userRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(user)
	updated, err := handler.userRepo.Update(
		ID,
		model.Userupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  updated,
	})

}

// DeleteUser method
func (handler *Handler) DeleteUser(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("userId queryParam is missing"))
	}

	user, err := handler.userRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.userRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  user,
	})

}

// SearchEmployee method
func (handler *Handler) SearchUser(context echo.Context) error {
	params := new(model.UserSearch)
	if err := context.Bind(params); err != nil {
		fmt.Println("no param")
		return context.JSON(http.StatusBadRequest, err)
	}
	cursor, ctx, err := handler.userRepo.Search(params)
	if err != nil {
		return context.JSON(http.StatusForbidden, err)
	}
	users := []model.User{}
	if err = cursor.All(ctx, &users); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, users)

}
