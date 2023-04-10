package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyCartUserID = "userId"

type Cart struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID             primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	RestaurantID       primitive.ObjectID `json:"restaurantId,omitempty" bson:"restaurantId,omitempty"`
	Items              []FoodWithQuantity `json:"items,omitempty" bson:"items,omitempty"`
	Coupon             string             `json:"coupon,omitempty" bson:"coupon,omitempty"`
	Notes              string             `json:"notes,omitempty" bson:"notes,omitempty"`
	TotalPrice         float64            `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	PriceAfterDiscount float64            `json:"priceAfterDiscount,omitempty" bson:"priceAfterDiscount,omitempty"`
	UpdatedOn          time.Time          `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}

func GetCartCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "carts")
}
