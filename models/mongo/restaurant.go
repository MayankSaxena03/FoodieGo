package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyRestaurantID = "_id"
var KeyRestaurantEmail = "email"
var KeyRestaurantPhone = "phone"
var KeyRestaurantPassword = "password"
var KeyRestaurantMenuID = "menuId"
var KeyRestaurantName = "name"
var KeyRestaurantAddress = "address"
var KeyRestaurantRating = "rating"

type Restaurant struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Rating      float64            `json:"rating,omitempty" bson:"rating,omitempty"`
	TotalRating int                `json:"totalRating,omitempty" bson:"totalRating,omitempty"`
	Address     Address            `json:"address,omitempty" bson:"address,omitempty"`
	IsVeg       bool               `json:"isVeg" bson:"isVeg"`
	MenuID      primitive.ObjectID `json:"menuId,omitempty" bson:"menuId,omitempty"`
	OpenTime    time.Time          `json:"openTime,omitempty" bson:"openTime,omitempty" validate:"required"`
	CloseTime   time.Time          `json:"closeTime,omitempty" bson:"closeTime,omitempty" validate:"required"`
	IsClosed    bool               `json:"isClosed" bson:"isClosed"`
	CreatedOn   time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	UpdatedOn   time.Time          `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}

func GetRestaurantCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "restaurants")
}
