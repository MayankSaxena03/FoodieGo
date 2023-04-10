package core

import (
	"context"

	"github.com/MayankSaxena03/FoodieGo/constants"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMenuByRestaurantID(ctx context.Context, restaurantID primitive.ObjectID) (models.Menu, error) {
	menuCollection := models.GetMenuCollection()
	query := bson.M{
		models.KeyMenuRestaurantID: restaurantID,
	}
	result := menuCollection.FindOne(ctx, query)
	var menu models.Menu
	err := result.Decode(&menu)
	return menu, err
}

func CreateMenu(ctx context.Context, menu models.Menu) error {
	menuCollection := models.GetMenuCollection()
	_, err := menuCollection.InsertOne(ctx, menu)
	if err != nil {
		return err
	}

	restaurantCollection := models.GetRestaurantCollection()
	query := bson.M{
		models.KeyRestaurantID: menu.RestaurantID,
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyRestaurantMenuID: menu.ID,
		},
	}
	_, err = restaurantCollection.UpdateOne(ctx, query, update)
	return err
}

func CreateMenuItem(ctx context.Context, menuID primitive.ObjectID, food models.Food) error {
	foodID, err := CreateFood(ctx, food)
	if err != nil {
		return err
	}

	menuCollection := models.GetMenuCollection()
	query := bson.M{
		models.KeyMenuID: menuID,
	}
	update := bson.M{
		constants.MongoKeywordPush: bson.M{
			models.KeyMenuItems: foodID,
		},
	}
	err = menuCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}

func DeleteMenuItem(ctx context.Context, menuID primitive.ObjectID, foodID primitive.ObjectID) error {
	menuCollection := models.GetMenuCollection()
	query := bson.M{
		models.KeyMenuID: menuID,
	}
	update := bson.M{
		constants.MongoKeywordPull: bson.M{
			models.KeyMenuItems: foodID,
		},
	}
	err := menuCollection.FindOneAndUpdate(ctx, query, update).Err()
	if err != nil {
		return err
	}

	err = DeleteFood(ctx, foodID)
	return err
}
