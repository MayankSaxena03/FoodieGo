package models

import (
	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyRatingOrderID = "orderId"
var KeyRatingUserID = "userId"

type Rating struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	OrderID primitive.ObjectID `json:"orderId,omitempty" bson:"orderId,omitempty"`
	Rating  int                `json:"rating,omitempty" bson:"rating,omitempty"`
}

func GetRatingCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "ratings")
}
