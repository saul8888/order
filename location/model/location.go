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
type AddressLo struct {
	Tag              string `bson:"tag" json:"tag"`
	FirstStreetLine  string `bson:"firtStreetLine" json:"firstStreetLine"`
	SecondStreetLine string `bson:"secondStreetLine" json:"secondStreetLine"`
	City             string `bson:"city" json:"city"`
	PostalCode       string `bson:"postalCode" json:"postalCode"`
	Notes            string `bson:"notes" json:"notes"`
}

type CoordinatesLo struct {
	Latitude  string `bson:"latitude" json:"latitude" validate:"required"`
	Longitude string `bson:"longitude" json:"longitude" validate:"required"`
}

type Location struct {
	LocationID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	MerchantID  string             `bson:"merchantId" json:"merchantId" validate:"required"` //["merchant_ID"]
	Email       string             `bson:"email" json:"email" validate:"email"`
	Addresses   string             `bson:"addresses" json:"addresses" validate:"required"` //[]AddressesLo
	Country     string             `bson:"country" json:"country" validate:"required"`     //"US",
	Currency    string             `bson:"currency" json:"currency" validate:"required"`   //"USD",
	Description string             `bson:"description" json:"description" validate:"required"`
	Website     string             `bson:"website" json:"website" validate:"required"`
	Twitter     string             `bson:"twitter" json:"twitter" validate:"required"`
	Instagram   string             `bson:"instagram" json:"instagram" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Coordinates CoordinatesLo
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//---------------------------------------------------------------//
type CreateLocation struct {
	LocationID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	MerchantID  string             `bson:"merchantId" json:"merchantId" validate:"required"` //["merchant_ID"]
	Email       string             `bson:"email" json:"email" validate:"email"`
	Addresses   string             `bson:"addresses" json:"addresses" validate:"required"` //[]AddressesLo
	Country     string             `bson:"country" json:"country" validate:"required"`     //"US",
	Currency    string             `bson:"currency" json:"currency" validate:"required"`   //"USD",
	Description string             `bson:"description" json:"description" validate:"required"`
	Website     string             `bson:"website" json:"website" validate:"required"`
	Twitter     string             `bson:"twitter" json:"twitter" validate:"required"`
	Instagram   string             `bson:"instagram" json:"instagram" validate:"required"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Coordinates CoordinatesLo
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LocationUpdate struct {
	Name        string `bson:"Name" json:"Name"`
	Currency    string `bson:"currency" json:"currency" validate:"required"` //"USD",
	Email       string `bson:"email" json:"email" validate:"email"`
	Description string `bson:"description" json:"description" validate:"required"`
	Website     string `bson:"website" json:"website" validate:"required"`
	Twitter     string `bson:"twitter" json:"twitter" validate:"required"`
	Instagram   string `bson:"instagram" json:"instagram" validate:"required"`
	Status      string `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
}

var Locationupdate map[string]interface{} = map[string]interface{}{
	"name":        "example",
	"currency":    "example",
	"email":       "example@example.com",
	"phoneNumber": "0000-00000",
	"instagram":   "example",
	"updatedat":   time.Now(),
}

type LocationSearch struct {
	Name        string `bson:"lastName" json:"lastName"`
	Email       string `bson:"email" json:"email" validate:"email"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
}

var Locationsearch map[string]interface{} = map[string]interface{}{
	"name":        "example",
	"email":       "example@example.com",
	"phoneNumber": "0000-00000",
}
