package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var KeyAddressID = "_id"
var KeyAddressCity = "city"

type Address struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	AddressLine1 string             `json:"addressLine1,omitempty" bson:"addressLine1,omitempty" validate:"required"`
	AddressLine2 string             `json:"addressLine2,omitempty" bson:"addressLine2,omitempty"`
	City         string             `json:"city,omitempty" bson:"city,omitempty" validate:"required"`
	Zip          string             `json:"zip,omitempty" bson:"zip,omitempty" validate:"required"`
	Landmark     string             `json:"landmark,omitempty" bson:"landmark,omitempty"`
	CreatedOn    time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	UpdatedOn    time.Time          `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}
