package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EntityID   primitive.ObjectID `json:"entityId,omitempty" bson:"entityId,omitempty"`
	EntityType string             `json:"entityType,omitempty" bson:"entityType,omitempty"`
	CreatedOn  time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	Changes    []Changes          `json:"changes,omitempty" bson:"changes,omitempty"`
}

type Changes struct {
	FieldName string      `json:"fieldName,omitempty" bson:"fieldName,omitempty"`
	OldValue  interface{} `json:"oldValue,omitempty" bson:"oldValue,omitempty"`
	NewValue  interface{} `json:"newValue,omitempty" bson:"newValue,omitempty"`
}

func GetLogCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "logs")
}
