package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address data model
type AddressCu struct {
	Country    string `bson:"country" json:"country" validate:"required"` //"US",
	State      string `bson:"state" json:"state" validate:"required"`     //"NY",
	City       string `bson:"city" json:"city" validate:"required"`
	Street     string `bson:"street" json:"street" validate:"required"`
	PostalCode string `bson:"postalCode" json:"postalCode" validate:"required"`
}

type Customer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Object      string             `bson:"object" json:"object"`
	FirstName   string             `bson:"firstName" json:"firstName" validate:"required"`
	LastName    string             `bson:"lastName" json:"lastName" validate:"required"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	Password    string             `bson:"password" json:"password" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Addresses   AddressCu          `bson:"addresses" json:"addresses" validate:"required"` //[]AddressCu
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
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
