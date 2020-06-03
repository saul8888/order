package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type Catalog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Category    string             `bson:"category" json:"category" validate:"required"`
	Description string             `bson:"description" json:"description" validate:"required"`
	LocationID  primitive.ObjectID `bson:"locationId" json:"locationId" validate:"required"`
	Options1    string             `bson:"options1" json:"options1" validate:"required"` //information additional
	Marcas      []Variations       `bson:"marcas" json:"marcas" validate:"required"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Variations struct {
	Name     string      `bson:"name" json:"name" validate:"required"`
	SKU      string      `bson:"SKU" json:"SKU"` //inforamtion the product
	Unit     string      `bson:"unit" json:"unit"`
	Price    int         `bson:"price" json:"price"`
	Currency string      `bson:"currency" json:"currency"`
	Stock    OptionStock `bson:"stock" json:"stock"`
	//CreatedAt time.Time   `bson:"createdAt" json:"createdAt"`
	//UpdatedAt time.Time   `bson:"updatedAt" json:"updatedAt"`
}

type OptionStock struct {
	Description string `bson:"description" json:"description"`
	InStock     int    `bson:"inStock" json:"inStock"`
	AlertStock  int    `bson:"alertStock" json:"alertStock"`
}

type CatalogUpdate struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

var Catalogupdate = map[string]interface{}{
	"name":        "",
	"currency":    "",
	"email":       "",
	"phoneNumber": "",
	"instagram":   "",
	"status":      "",
	"updatedat":   time.Now(),
}

type CatalogSearch struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
