package core

import (
	"context"
	"errors"

	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddRating(ctx context.Context, userID primitive.ObjectID, orderID primitive.ObjectID, rating int) error {
	ratingCollection := models.GetRatingCollection()
	query := bson.M{
		models.KeyRatingOrderID: orderID,
		models.KeyRatingUserID:  userID,
	}
	result := ratingCollection.FindOne(ctx, query)
	if result.Err() == nil && result.Err() != mongo.ErrNoDocuments {
		return result.Err()
	} else if result.Err() == mongo.ErrNoDocuments {
		_, err := ratingCollection.InsertOne(ctx, models.Rating{
			UserID:  userID,
			OrderID: orderID,
			Rating:  rating,
		})
		if err != nil {
			return err
		}
	} else {
		return errors.New("already rated")
	}

	restaurant, err := GetRestaurantByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	restaurant.Rating = restaurant.Rating*float64(restaurant.TotalRating) + float64(rating)
	restaurant.TotalRating++
	restaurant.Rating /= float64(restaurant.TotalRating)

	err = UpdateRestaurant(ctx, restaurant.ID, restaurant)
	if err != nil {
		return err
	}

	allFood, err := GetFoodByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	for _, food := range allFood {
		food.Rating = food.Rating*float64(food.TotalRating) + float64(rating)
		food.TotalRating++
		food.Rating /= float64(food.TotalRating)

		err = UpdateFood(ctx, food.ID, food)
		if err != nil {
			return err
		}
	}

	return nil
}
