package core

import (
	"context"
	"time"

	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	userCollection := models.GetUserCollection()
	user.CreatedOn = time.Now()
	user.LastLoginOn = time.Now()
	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User
	query := bson.M{
		models.KeyUserID: id,
	}
	userCollection := models.GetUserCollection()
	err := userCollection.FindOne(ctx, query).Decode(&user)
	return user, err
}

func UpdateUserByID(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	query := bson.M{
		models.KeyUserID: id,
	}

	userCollection := models.GetUserCollection()
	err := userCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}
