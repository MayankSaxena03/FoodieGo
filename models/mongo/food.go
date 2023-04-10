package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyFoodID = "_id"
var KeyFoodRating = "rating"

type Food struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Rating      float64            `json:"rating,omitempty" bson:"rating,omitempty"`
	TotalRating int                `json:"totalRating,omitempty" bson:"totalRating,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty" validate:"required"`
	IsVeg       bool               `json:"isVeg" bson:"isVeg" validate:"required"`
	IsAvailable bool               `json:"isAvailable" bson:"isAvailable" default:"true"`
	CreatedOn   time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	UpdatedOn   time.Time          `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}

type FoodWithQuantity struct {
	Food
	Quantity int `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

func GetFoodCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "foods")
}
