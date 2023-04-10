package core

import (
	"context"

	"github.com/MayankSaxena03/FoodieGo/constants"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFoodByIDs(ctx context.Context, ids []primitive.ObjectID, skip int64, limit int64) ([]models.Food, error) {
	foodCollection := models.GetFoodCollection()
	query := bson.M{
		models.KeyFoodID: bson.M{
			constants.MongoKeywordIn: ids,
		},
	}
	cursor, err := foodCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			models.KeyFoodRating: -1,
		},
	})
	if err != nil {
		return nil, err
	}
	var foods []models.Food
	err = cursor.All(ctx, &foods)
	if err != nil {
		return nil, err
	}
	return foods, nil
}

func GetAllFoodByRestaurantID(ctx context.Context, restaurantID primitive.ObjectID, skip int64, limit int64) ([]models.Food, error) {
	menuCollection := models.GetMenuCollection()
	query := bson.M{
		models.KeyMenuRestaurantID: restaurantID,
	}
	result := menuCollection.FindOne(ctx, query)
	var menu models.Menu
	err := result.Decode(&menu)
	if err != nil {
		return nil, err
	}
	return GetFoodByIDs(ctx, menu.MenuItems, skip, limit)
}

func GetMenuByFoodID(ctx context.Context, foodID primitive.ObjectID) (models.Menu, error) {
	menuCollection := models.GetMenuCollection()
	query := bson.M{
		models.KeyMenuItems: foodID,
	}
	result := menuCollection.FindOne(ctx, query)
	var menu models.Menu
	err := result.Decode(&menu)
	return menu, err
}

func CreateFood(ctx context.Context, food models.Food) (primitive.ObjectID, error) {
	foodCollection := models.GetFoodCollection()
	result, err := foodCollection.InsertOne(ctx, food)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func DeleteFood(ctx context.Context, foodID primitive.ObjectID) error {
	foodCollection := models.GetFoodCollection()
	_, err := foodCollection.DeleteOne(ctx, bson.M{
		models.KeyFoodID: foodID,
	})
	return err
}

func UpdateFood(ctx context.Context, foodID primitive.ObjectID, food models.Food) error {
	foodCollection := models.GetFoodCollection()
	_, err := foodCollection.UpdateOne(ctx, bson.M{
		models.KeyFoodID: food.ID,
	}, bson.M{
		constants.MongoKeywordSet: food,
	})
	return err
}

func GetFoodByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]models.Food, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID: orderID,
	}
	result := orderCollection.FindOne(ctx, query)
	var order models.Order
	err := result.Decode(&order)
	if err != nil {
		return nil, err
	}

	var foodIds []primitive.ObjectID
	for _, item := range order.Items {
		foodIds = append(foodIds, item.ID)
	}

	return GetFoodByIDs(ctx, foodIds, 0, 0)
}
