package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type AmountMoney struct {
	Amount   string `bson:"amount" json:"amount" validate:"required"`     //200,
	Currency string `bson:"currency" json:"currency" validate:"required"` //"USD"
}

type Payment struct {
	PaymentID     primitive.ObjectID
	Name          string `bson:"name" json:"name" validate:"required"`
	Status        string `bson:"status" json:"status" validate:"required"`               //"ACTIVE"
	Country       string `bson:"country" json:"country" validate:"required"`             //"US",
	LanguageCode  string `bson:"languageCode" json:"languageCode" validate:"required"`   //"en-US",
	Currency      string `bson:"currency" json:"currency" validate:"required"`           //"USD",
	Amount        string `bson:"amount" json:"amount" validate:"required"`               //AmountMoney
	LocationID    string `bson:"locationID" json:"locationID" validate:"required"`       //["LOCATION_ID"],
	OrderID       string `bson:"orderID" json:"orderID" validate:"required"`             //["ORDER_ID"],
	CustomerID    string `bson:"customerID" json:"customerID" validate:"required"`       //["CUSTOMER_ID"],
	ReceiptNumber string `bson:"receiptNumber" json:"receiptNumber" validate:"required"` //"GQTF",
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreatePayment struct {
	PaymentID     primitive.ObjectID
	Name          string `bson:"name" json:"name" validate:"required"`
	Status        string `bson:"status" json:"status" validate:"required"`               //"ACTIVE"
	Country       string `bson:"country" json:"country" validate:"required"`             //"US",
	LanguageCode  string `bson:"languageCode" json:"languageCode" validate:"required"`   //"en-US",
	Currency      string `bson:"currency" json:"currency" validate:"required"`           //"USD",
	Amount        string `bson:"amount" json:"amount" validate:"required"`               //AmountMoney
	LocationID    string `bson:"locationID" json:"locationID" validate:"required"`       //["LOCATION_ID"],
	OrderID       string `bson:"orderID" json:"orderID" validate:"required"`             //["ORDER_ID"],
	CustomerID    string `bson:"customerID" json:"customerID" validate:"required"`       //["CUSTOMER_ID"],
	ReceiptNumber string `bson:"receiptNumber" json:"receiptNumber" validate:"required"` //"GQTF",
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PaymentUpdate struct {
	Name     string `bson:"name" json:"name"`
	Status   string `bson:"status" json:"status" validate:"required"`
	Currency string `bson:"currency" json:"currency" validate:"required"`
	Amount   string `bson:"amount" json:"amount" validate:"required"` //AmountMoney
}

var Paymentupdate map[string]interface{} = map[string]interface{}{
	"name":     "example",
	"status":   "example",
	"currency": "example",
	"amount":   "example",
}

type PaymentSearch struct {
	Name          string `bson:"name" json:"name"`
	Status        string `bson:"status" json:"status" validate:"required"`               //"ACTIVE"
	ReceiptNumber string `bson:"receiptNumber" json:"receiptNumber" validate:"required"` //"GQTF",
}

var Paymentsearch map[string]interface{} = map[string]interface{}{
	"name":          "example",
	"status":        "example",
	"receiptNumber": "example",
}
