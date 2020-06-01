package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

// Address data model
type AddressMe struct {
	Country    string `bson:"country" json:"country" validate:"required"` //"US",
	State      string `bson:"state" json:"state"`                         //"NY",
	City       string `bson:"city" json:"city"`
	Street     string `bson:"street" json:"street"`
	PostalCode string `bson:"postalCode" json:"postalCode"`
}

type Merchant struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required"`             //"ACTIVE"
	LanguageCode string             `bson:"languageCode" json:"languageCode" validate:"required"` //"en-US",
	Currency     string             `bson:"currency" json:"currency" validate:"required"`         //"USD",
	Addresses    AddressMe          `bson:"addresses" json:"addresses"`                           //[]AddressesMe
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type MerchantUpdate struct {
	Name      string    `bson:"name" json:"name"`
	Status    string    `bson:"status" json:"status" validate:"required"`
	Currency  string    `bson:"currency" json:"currency" validate:"required"`
	Addresses AddressMe `bson:"addresses" json:"addresses"`
}

var Merchantupdate = map[string]interface{}{
	"name":              "",
	"status":            "",
	"currency":          "",
	"addresses.country": "",
	"updatedat":         time.Now(),
}

type MerchantSearch struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email" validate:"email"`
	Password string `bson:"password" json:"password"`
}

var Merchantsearch = map[string]interface{}{
	"Name":     "example",
	"email":    "example@example.com",
	"password": "admin",
}
