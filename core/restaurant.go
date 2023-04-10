package core

import (
	"context"

	"github.com/MayankSaxena03/FoodieGo/constants"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckIfRestaurantExists(ctx context.Context, email string) (bool, error) {
	restaurantCollection := models.GetRestaurantCollection()
	query := bson.M{
		models.KeyRestaurantEmail: email,
	}
	count, err := restaurantCollection.CountDocuments(ctx, query)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}
	return count > 0, nil
}

func CreateRestaurant(ctx context.Context, restaurant models.Restaurant) (*mongo.InsertOneResult, error) {
	restaurantCollection := models.GetRestaurantCollection()
	result, err := restaurantCollection.InsertOne(ctx, restaurant)
	return result, err
}

func GetRestaurantByID(ctx context.Context, id primitive.ObjectID) (models.Restaurant, error) {
	restaurantCollection := models.GetRestaurantCollection()
	query := bson.M{
		models.KeyRestaurantID: id,
	}
	var project = bson.M{
		models.KeyRestaurantPassword: 0,
	}
	var restaurant models.Restaurant
	err := restaurantCollection.FindOne(ctx, query, options.FindOne().SetProjection(project)).Decode(&restaurant)
	return restaurant, err
}

func GetRestaurantByEmail(ctx context.Context, email string) (models.Restaurant, error) {
	restaurantCollection := models.GetRestaurantCollection()
	query := bson.M{
		models.KeyRestaurantEmail: email,
	}
	var project = bson.M{
		models.KeyRestaurantPassword: 0,
	}
	var restaurant models.Restaurant
	err := restaurantCollection.FindOne(ctx, query, options.FindOne().SetProjection(project)).Decode(&restaurant)
	return restaurant, err
}

func GetRestaurantsByCity(ctx context.Context, city string, skip int64, limit int64) ([]models.Restaurant, error) {
	restaurantCollection := models.GetRestaurantCollection()
	query := bson.M{
		helpers.MongoJoinFields(models.KeyRestaurantAddress, models.KeyAddressCity): city,
	}
	var project = bson.M{
		models.KeyRestaurantPassword: 0,
	}
	var restaurants []models.Restaurant
	cursor, err := restaurantCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			models.KeyRestaurantRating: -1,
		},
		Projection: project,
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &restaurants)
	return restaurants, err
}

func UpdateRestaurant(ctx context.Context, id primitive.ObjectID, restaurant models.Restaurant) error {
	restaurantCollection := models.GetRestaurantCollection()
	query := bson.M{
		models.KeyRestaurantID: id,
	}
	update := bson.M{
		constants.MongoKeywordSet: restaurant,
	}

	_, err := restaurantCollection.UpdateOne(ctx, query, update)
	return err
}

func GetRestaurantIncome(ctx context.Context, restaurantId primitive.ObjectID) (float64, error) {
	ordersCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderRestaurantID: restaurantId,
		models.KeyOrderStatus: bson.M{
			constants.MongoKeywordIn: []string{constants.KeyOrderDelivered, constants.KeyOrderAccepted, constants.KeyOrderOutForDelivery},
		},
	}
	project := bson.M{
		models.KeyOrderPriceAfterOffer: 1,
	}
	cursor, err := ordersCollection.Find(ctx, query, &options.FindOptions{
		Projection: project,
	})
	if err != nil {
		return 0, err
	}
	var orders []models.Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return 0, err
	}
	var income float64
	for _, order := range orders {
		income += order.PriceAfterOffer
	}
	return income, nil
}
