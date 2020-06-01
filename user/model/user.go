package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName     string             `bson:"userName" json:"userName" validate:"required"`         //"saul8",
	UserCreate   string             `bson:"userCreate" json:"userCreate" validate:"required"`     //"true",
	UserSuper    string             `bson:"userSuper" json:"userSuper" validate:"required"`       //"true",
	UserCatalogs string             `bson:"userCatalogs" json:"userCatalogs" validate:"required"` //"false", ->User can update system catalogs
	Password     string             `bson:"password" json:"password" validate:"required"`         //"121s31rg1rs3",
	Valuntil     string             `bson:"valuntil" json:"valuntil" validate:"required"`         //"infinity", ->abstime password expiry time
	UserLimit    string             `bson:"userLimit" json:"userLimit" validate:"required"`       //2, ->The number of connections that the user can open
	Status       string             `bson:"status" json:"status" validate:"required"`             //"ACTIVE"
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type UserUpdate struct {
	UserName     string    `bson:"userName" json:"userName" validate:"required"`         //"saul8",
	UserCreate   string    `bson:"userCreate" json:"userCreate" validate:"required"`     //"true",
	UserSuper    string    `bson:"userSuper" json:"userSuper" validate:"required"`       //"true",
	UserCatalogs string    `bson:"userCatalogs" json:"userCatalogs" validate:"required"` //"false", ->User can update system catalogs
	Valuntil     string    `bson:"valuntil" json:"valuntil" validate:"required"`         //"infinity", ->abstime password expiry time
	UserLimit    string    `bson:"userLimit" json:"userLimit" validate:"required"`       //2, ->The number of connections that the user can open
	Status       string    `bson:"status" json:"status" validate:"required"`             //"ACTIVE"
	UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
}

var Userupdate = map[string]interface{}{
	"userName":     "",
	"userCreate":   "",
	"userSuper":    "",
	"userCatalogs": "",
	"valuntil":     "",
	"userLimit":    "",
	"status":       "",
	"updatedat":    time.Now(),
}

type UserSearch struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email" validate:"email"`
	Password string `bson:"password" json:"password"`
}

var Usersearch = map[string]interface{}{
	"Name":     "example",
	"email":    "example@example.com",
	"password": "admin",
}

/*
{
    "user": [
        {
            "id": "6AH6E4EXAMPLE",
            "usename": "saul8",
            "usecreate": "true",
            "usesuper": "true",
            "usecatalogs": "false", ->User can update system catalogs
            "passwd": "121s31rg1rs3",
            "valuntil": "infinity", ->abstime password expiry time
            "useconnlimit": 2, ->The number of connections that the user can open
            "status": "ACTIVE"
            "created_at": "2019-02-20T01:28:49Z",
            "updated_at": "2019-02-20T01:28:49Z"
        }
    ]
}

The number of connections that the user can open.
*/
