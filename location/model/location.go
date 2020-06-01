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
	Country    string `bson:"country" json:"country" validate:"required"` //"US",
	State      string `bson:"state" json:"state" validate:"required"`     //"NY",
	City       string `bson:"city" json:"city" validate:"required"`
	Street     string `bson:"street" json:"street" validate:"required"`
	PostalCode string `bson:"postalCode" json:"postalCode" validate:"required"`
}

type CoordinatesLo struct {
	Latitude  string `bson:"latitude" json:"latitude" validate:"required"`
	Longitude string `bson:"longitude" json:"longitude" validate:"required"`
}

type Location struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	MerchantID  primitive.ObjectID `bson:"merchantId" json:"merchantId"`
	Email       string             `bson:"email" json:"email" validate:"email"`
	Addresses   AddressLo          `bson:"addresses" json:"addresses"`
	Currency    string             `bson:"currency" json:"currency" validate:"required"` //"USD",
	Description string             `bson:"description" json:"description"`
	Website     string             `bson:"website" json:"website" validate:"required"`
	Twitter     string             `bson:"twitter" json:"twitter"`
	Instagram   string             `bson:"instagram" json:"instagram"`
	Status      string             `bson:"status" json:"status" validate:"required"` //"ACTIVE"
	PhoneNumber string             `bson:"phoneNumber" json:"phoneNumber"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
	//Coordinates CoordinatesLo      `bson:"coordinates" json:"coordinates"`
}

type LocationUpdate struct {
	Name        string `bson:"Name" json:"Name"`
	Currency    string `bson:"currency" json:"currency"` //"USD",
	Email       string `bson:"email" json:"email" validate:"email"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	Instagram   string `bson:"instagram" json:"instagram"`
	Status      string `bson:"status" json:"status"`
}

var Locationupdate = map[string]interface{}{
	"name":        "",
	"currency":    "",
	"email":       "",
	"phoneNumber": "",
	"instagram":   "",
	"status":      "",
	"updatedat":   time.Now(),
}

type LocationSearch struct {
	Name        string `bson:"lastName" json:"lastName"`
	Email       string `bson:"email" json:"email" validate:"email"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
}

var Locationsearch = map[string]interface{}{
	"name":        "",
	"email":       "",
	"phoneNumber": "",
}
