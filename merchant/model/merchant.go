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
	Tag              string `bson:"tag" json:"tag"`
	FirstStreetLine  string `bson:"firtStreetLine" json:"firstStreetLine"`
	SecondStreetLine string `bson:"secondStreetLine" json:"secondStreetLine"`
	City             string `bson:"city" json:"city"`
	PostalCode       string `bson:"postalCode" json:"postalCode"`
	Notes            string `bson:"notes" json:"notes"`
}

type Merchant struct {
	MerchantID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required"`             //"ACTIVE"
	Country      string             `bson:"country" json:"country" validate:"required"`           //"US",
	LanguageCode string             `bson:"languageCode" json:"languageCode" validate:"required"` //"en-US",
	Currency     string             `bson:"currency" json:"currency" validate:"required"`         //"USD",
	Addresses    string             `bson:"addresses" json:"addresses"`                           //[]AddressesMe
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateMerchant struct {
	MerchantID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required"`             //"ACTIVE"
	Country      string             `bson:"country" json:"country" validate:"required"`           //"US",
	LanguageCode string             `bson:"languageCode" json:"languageCode" validate:"required"` //"en-US",
	Currency     string             `bson:"currency" json:"currency" validate:"required"`         //"USD",
	Addresses    string             `bson:"addresses" json:"addresses"`                           //[]AddressesMe
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MerchantUpdate struct {
	Name      string `bson:"name" json:"name"`
	Status    string `bson:"status" json:"status" validate:"required"`
	Currency  string `bson:"currency" json:"currency" validate:"required"`
	Addresses string `bson:"addresses" json:"addresses"` //[]AddressesEm
}

var Merchantupdate map[string]interface{} = map[string]interface{}{
	"name":      "example",
	"status":    "example",
	"currency":  "example",
	"addresses": "example",
	"updatedat": time.Now(),
}

type MerchantSearch struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email" validate:"email"`
	Password string `bson:"password" json:"password"`
}

var Merchantsearch map[string]interface{} = map[string]interface{}{
	"Name":     "example",
	"email":    "example@example.com",
	"password": "admin",
}
