package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address data model
type AddressesEm struct {
	Country    string `bson:"country" json:"country"` //"US",
	City       string `bson:"city" json:"city"`
	PostalCode string `bson:"postalCode" json:"postalCode"`
}

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Object     string             `bson:"object" json:"object"`
	LocationID primitive.ObjectID `bson:"locationId" json:"locationId"`
	Email      string             `bson:"email" json:"email"`
	Status     string             `bson:"status" json:"status"`       //"ACTIVE" createdat
	Addresses  AddressesEm        `bson:"addresses" json:"addresses"` //[]AddressEm
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type CreateEmployee struct {
	Object     string      `bson:"object" json:"object" validate:"required"`
	LocationID string      `bson:"locationId" json:"locationId" validate:"required"`
	Email      string      `bson:"email" json:"email" validate:"email"`
	Status     string      `bson:"status" json:"status" validate:"required"` //"ACTIVE" createdat
	Addresses  AddressesEm `bson:"addresses" json:"addresses"`
}

type EmployeeUpdate struct {
	Status    string      `bson:"status" json:"status"` //"ACTIVE" createdat
	Email     string      `bson:"email" json:"email" validate:"email"`
	Addresses AddressesEm `bson:"addresses" json:"addresses"` //[]AddressesEm
}

var Employeeupdate map[string]interface{} = map[string]interface{}{
	"status":            "example",
	"email":             "example@example.com",
	"addresses.country": "example",
	"updatedat":         time.Now(),
}

type EmployeeSearch struct {
	Email  string `bson:"email" json:"email"`
	Status string `bson:"status" json:"status"` //"ACTIVE" createdat
}

var Employeesearch map[string]interface{} = map[string]interface{}{
	"email":  "example@example.com",
	"status": "admin",
}

type Location struct {
	LocationID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	MerchantID  string             `bson:"merchantId" json:"merchantId" validate:"required"` //["merchant_ID"]
	Email       string             `bson:"email" json:"email" validate:"email"`
	Addresses   string             `bson:"addresses" json:"addresses"`                   //AddressLo
	Currency    string             `bson:"currency" json:"currency" validate:"required"` //"USD",
	Description string             `bson:"description" json:"description" validate:"required"`
	Website     string             `bson:"website" json:"website" validate:"required"`
	Twitter     string             `bson:"twitter" json:"twitter" validate:"required"`
	Instagram   string             `bson:"instagram" json:"instagram" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
	//Coordinates CoordinatesLo      `bson:"coordinates" json:"coordinates"`
}
