package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyMenuID = "_id"
var KeyMenuItems = "menuItems"
var KeyMenuRestaurantID = "restaurantId"

type Menu struct {
	ID           primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	RestaurantID primitive.ObjectID   `json:"restaurantId,omitempty" bson:"restaurantId,omitempty"`
	MenuItems    []primitive.ObjectID `json:"menuItems,omitempty" bson:"menuItems,omitempty"`
	CreatedOn    time.Time            `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	UpdatedOn    time.Time            `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}

func GetMenuCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "menu")
}
