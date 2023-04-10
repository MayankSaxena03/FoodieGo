package core

import (
	"context"
	"errors"
	"time"

	"github.com/MayankSaxena03/FoodieGo/constants"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckIfAddressIsPrimary(ctx context.Context, userID primitive.ObjectID, addressID primitive.ObjectID) (bool, error) {
	userCollection := models.GetUserCollection()
	query := bson.M{
		models.KeyUserID:             userID,
		models.KeyUserPrimaryAddress: addressID,
		helpers.MongoJoinFields(models.KeyUserAddress, models.KeyAddressID): addressID,
	}
	result := userCollection.FindOne(ctx, query)
	return result.Err() == nil, result.Err()
}

func AddAddressInUser(ctx context.Context, userID primitive.ObjectID, address models.Address) error {
	userCollection := models.GetUserCollection()
	address.CreatedOn = time.Now()
	address.UpdatedOn = time.Now()
	address.ID = primitive.NewObjectID()
	query := bson.M{
		models.KeyUserID: userID,
	}
	update := bson.M{
		constants.MongoKeywordPush: bson.M{
			models.KeyUserAddress:        address,
			models.KeyUserPrimaryAddress: address.ID,
		},
	}
	result := userCollection.FindOneAndUpdate(ctx, query, update)
	return result.Err()
}

func UpdateAddressInUser(ctx context.Context, userID primitive.ObjectID, addressID primitive.ObjectID, address models.Address) error {
	userCollection := models.GetUserCollection()
	address.UpdatedOn = time.Now()

	query := bson.M{
		models.KeyUserID: userID,
		helpers.MongoJoinFields(models.KeyUserAddress, models.KeyAddressID): addressID,
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			helpers.MongoJoinFields(models.KeyUserAddress, constants.MongoKeyPositionalOperator): address,
			models.KeyUserPrimaryAddress: address.ID,
		},
	}

	err := userCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}

func DeleteAddressInUser(ctx context.Context, userID primitive.ObjectID, addressID primitive.ObjectID) error {
	userCollection := models.GetUserCollection()
	isPrimary, err := CheckIfAddressIsPrimary(ctx, userID, addressID)
	if err != nil {
		return err
	}
	if isPrimary {
		return errors.New("address is primary address")
	}
	query := bson.M{
		models.KeyUserID: userID,
		helpers.MongoJoinFields(models.KeyUserAddress, models.KeyAddressID): addressID,
	}
	update := bson.M{
		constants.MongoKeywordPull: bson.M{
			models.KeyUserAddress: bson.M{
				models.KeyAddressID: addressID,
			},
		},
	}
	result := userCollection.FindOneAndUpdate(ctx, query, update)
	return result.Err()
}

func GetAllAddressesOfUser(ctx context.Context, userID primitive.ObjectID) ([]models.Address, error) {
	userCollection := models.GetUserCollection()
	var user models.User
	query := bson.M{
		models.KeyUserID: userID,
	}
	result := userCollection.FindOne(ctx, query)
	err := result.Decode(&user)
	return user.Address, err
}

func GetAddressByIDInUser(ctx context.Context, userID primitive.ObjectID, addressID primitive.ObjectID) (models.Address, error) {
	userCollection := models.GetUserCollection()
	var user models.User
	query := bson.M{
		models.KeyUserID: userID,
		helpers.MongoJoinFields(models.KeyUserAddress, models.KeyAddressID): addressID,
	}
	result := userCollection.FindOne(ctx, query)
	err := result.Decode(&user)
	if err != nil {
		return models.Address{}, err
	}
	for _, address := range user.Address {
		if address.ID == addressID {
			return address, err
		}
	}
	return models.Address{}, errors.New("address not found")
}

func SetPrimaryAddressInUser(ctx context.Context, userID primitive.ObjectID, addressID primitive.ObjectID) error {
	userCollection := models.GetUserCollection()
	query := bson.M{
		models.KeyUserID: userID,
		helpers.MongoJoinFields(models.KeyUserAddress, models.KeyAddressID): addressID,
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyUserPrimaryAddress: addressID,
		},
	}

	err := userCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}
