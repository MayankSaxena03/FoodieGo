package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyOfferID = "_id"
var KeyOfferRestaurantID = "restaurantId"
var KeyOfferCoupon = "coupon"
var KeyOfferValidFrom = "validFrom"
var KeyOfferValidTill = "validTill"

type Offer struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Coupon       string             `json:"coupon,omitempty" bson:"coupon,omitempty" validate:"required"`
	RestaurantID primitive.ObjectID `json:"restaurantId,omitempty" bson:"restaurantId,omitempty"`
	Discount     int                `json:"discount,omitempty" bson:"discount,omitempty" validate:"required,gte=0,lte=100"`
	MaxDiscount  float64            `json:"maxDiscount,omitempty" bson:"maxDiscount,omitempty"`
	ValidFrom    time.Time          `json:"validFrom,omitempty" bson:"validFrom,omitempty" validate:"required"`
	ValidTill    time.Time          `json:"validTill,omitempty" bson:"validTill,omitempty" validate:"required"`
	CreatedOn    time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	UpdatedOn    time.Time          `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}

func GetOfferCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "offers")
}
