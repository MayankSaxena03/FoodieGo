package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyDeliveryPartnerID = "_id"
var KeyDeliveryPartnerPhone = "phone"
var KeyDeliveryPartnerLastLoginOn = "lastLoginOn"
var KeyDeliveryPartnerName = "name"
var KeyDeliveryPartnerEmail = "email"
var KeyDeliveryPartnerCity = "city"
var KeyDeliveryPartnerBirthday = "birthday"
var KeyDeliveryPartnerProfileImage = "profileImage"

type DeliveryPartner struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	CreatedOn   time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	LastLoginOn time.Time          `json:"lastLoginOn,omitempty" bson:"lastLoginOn,omitempty"`
	DeliveryPartnerProfile
}

type DeliveryPartnerProfile struct {
	Name         string    `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" bson:"email,omitempty"`
	City         string    `json:"city,omitempty" bson:"city,omitempty" validate:"required"`
	Birthday     time.Time `json:"birthday,omitempty" bson:"birthday,omitempty"`
	ProfileImage string    `json:"profileImage,omitempty" bson:"profileImage,omitempty"`
}

func GetDeliveryPartnerCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "deliveryPartners")
}
