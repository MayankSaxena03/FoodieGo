package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyUserID = "_id"
var KeyUserLastLoginOn = "lastLoginOn"
var KeyUserPhone = "phone"
var KeyUserName = "name"
var KeyUserEmail = "email"
var KeyUserBirthday = "birthday"
var KeyUserAddress = "address"
var KeyUserPrimaryAddress = "primaryAddress"

type User struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Phone            string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Address          []Address          `json:"address,omitempty" bson:"address,omitempty"`
	PrimaryAddressID primitive.ObjectID `json:"primaryAddress,omitempty" bson:"primaryAddress,omitempty"`
	CreatedOn        time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	LastLoginOn      time.Time          `json:"lastLoginOn,omitempty" bson:"lastLoginOn,omitempty"`
	UserProfile
}

type UserProfile struct {
	Name     string    `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email    string    `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Birthday time.Time `json:"birthday,omitempty" bson:"birthday,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

func GetUserCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "users")
}
