package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address data model
type AddressCu struct {
	Tag              string `bson:"tag" json:"tag"`
	FirstStreetLine  string `bson:"firtStreetLine" json:"firstStreetLine"`
	SecondStreetLine string `bson:"secondStreetLine" json:"secondStreetLine"`
	City             string `bson:"city" json:"city"`
	PostalCode       string `bson:"postalCode" json:"postalCode"`
	Notes            string `bson:"notes" json:"notes"`
}

type Customer struct {
	CustomerID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName   string             `bson:"firstName" json:"firstName" validate:"required"`
	LastName    string             `bson:"lastName" json:"lastName" validate:"required"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	Password    string             `bson:"password" json:"password" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Addresses   string             `bson:"addresses" json:"addresses" validate:"required"` //[]AddressCu
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type CreateCustomer struct {
	CustomerID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName   string             `bson:"firstName" json:"firstName" validate:"required"`
	LastName    string             `bson:"lastName" json:"lastName" validate:"required"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	Password    string             `bson:"password" json:"password" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Addresses   string             `bson:"addresses" json:"addresses" validate:"required"` //[]AddressCu
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CustomerUpdate struct {
	FirstName   string `bson:"firstName" json:"firstName"`
	LastName    string `bson:"lastName" json:"lastName"`
	Email       string `bson:"email" json:"email" validate:"email"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	Addresses   string `bson:"addresses" json:"addresses"` //[]AddressesEm
}

var Customerupdate map[string]interface{} = map[string]interface{}{
	"firstName":   "example",
	"lastName":    "example",
	"email":       "example@example.com",
	"phoneNumber": "0000-00000",
	"addresses":   "example",
	"updatedat":   time.Now(),
}

type CustomerSearch struct {
	FirstName   string `bson:"firstName" json:"firstName"`
	LastName    string `bson:"lastName" json:"lastName"`
	Email       string `bson:"email" json:"email" validate:"email"`
	Password    string `bson:"password" json:"password"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
}

var Customersearch map[string]interface{} = map[string]interface{}{
	"firstName":   "example",
	"lastName":    "example",
	"email":       "example@example.com",
	"password":    "admin",
	"phoneNumber": "0000-00000",
}
