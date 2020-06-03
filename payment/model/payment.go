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
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Object        string             `bson:"object" json:"object"`
	Name          string             `bson:"name" json:"name" validate:"required"`
	Status        string             `bson:"status" json:"status" validate:"required"`             //"ACTIVE"
	Country       string             `bson:"country" json:"country" validate:"required"`           //"US",
	LanguageCode  string             `bson:"languageCode" json:"languageCode" validate:"required"` //"en-US",
	Currency      string             `bson:"currency" json:"currency" validate:"required"`         //"USD",
	Amount        AmountMoney        `bson:"amount" json:"amount" validate:"required"`             //AmountMoney
	LocationID    primitive.ObjectID `bson:"locationID" json:"locationID" validate:"required"`
	OrderID       string             `bson:"orderID" json:"orderID" validate:"required"` //["ORDER_ID"],
	CustomerID    primitive.ObjectID `bson:"customerID" json:"customerID" validate:"required"`
	ReceiptNumber string             `bson:"receiptNumber" json:"receiptNumber" validate:"required"` //"GQTF",
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type PaymentUpdate struct {
	Name     string      `bson:"name" json:"name"`
	Status   string      `bson:"status" json:"status" validate:"required"`
	Currency string      `bson:"currency" json:"currency" validate:"required"`
	Amount   AmountMoney `bson:"amount" json:"amount" validate:"required"` //AmountMoney
}

var Paymentupdate = map[string]interface{}{
	"name":               "",
	"status":             "",
	"currency":           "",
	"amountMoney.amount": "",
}

type PaymentSearch struct {
	Name          string `bson:"name" json:"name"`
	Status        string `bson:"status" json:"status" validate:"required"`               //"ACTIVE"
	ReceiptNumber string `bson:"receiptNumber" json:"receiptNumber" validate:"required"` //"GQTF",
}

var Paymentsearch = map[string]interface{}{
	"name":          "",
	"status":        "",
	"receiptNumber": "",
}
