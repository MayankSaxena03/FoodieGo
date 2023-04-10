package core

import (
	"context"
	"errors"
	"time"

	"github.com/MayankSaxena03/FoodieGo/constants"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOffer(ctx context.Context, restaurantID primitive.ObjectID, offer models.Offer) error {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferRestaurantID: restaurantID,
		models.KeyOfferCoupon:       offer.Coupon,
	}
	var existingOffer models.Offer
	err := offerCollection.FindOne(ctx, query).Decode(&existingOffer)
	if err != nil && err != mongo.ErrNoDocuments {
		return errors.New("offer already exists")
	}

	offer.CreatedOn = time.Now()
	offer.UpdatedOn = time.Now()
	_, err = offerCollection.InsertOne(ctx, offer)
	return err
}

func GetAllOffersByRestaurantID(ctx context.Context, restaurantID primitive.ObjectID, skip int64, limit int64) ([]models.Offer, error) {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferRestaurantID: restaurantID,
	}

	cursor, err := offerCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})

	if err != nil {
		return nil, err
	}

	var offers []models.Offer
	err = cursor.All(ctx, &offers)
	if err != nil {
		return nil, err
	}

	return offers, nil
}

func GetValidOffersByRestaurantID(ctx context.Context, restaurantID primitive.ObjectID, skip int64, limit int64) ([]models.Offer, error) {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferRestaurantID: restaurantID,
		models.KeyOfferValidFrom: bson.M{
			"$lte": time.Now(),
		},
		models.KeyOfferValidTill: bson.M{
			"$gte": time.Now(),
		},
	}

	cursor, err := offerCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})

	if err != nil {
		return nil, err
	}

	var offers []models.Offer
	err = cursor.All(ctx, &offers)
	if err != nil {
		return nil, err
	}

	return offers, nil
}

func GetOfferByID(ctx context.Context, offerId primitive.ObjectID) (models.Offer, error) {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferID: offerId,
	}
	var offer models.Offer
	err := offerCollection.FindOne(ctx, query).Decode(&offer)
	return offer, err
}

func UpdateOfferByID(ctx context.Context, offerId primitive.ObjectID, offer models.Offer) error {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferID: offerId,
	}
	update := bson.M{
		constants.MongoKeywordSet: offer,
	}
	_, err := offerCollection.UpdateOne(ctx, query, update)
	return err
}

func DeleteOfferByID(ctx context.Context, offerId primitive.ObjectID) error {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferID: offerId,
	}
	_, err := offerCollection.DeleteOne(ctx, query)
	return err
}

func GetValidOfferByCoupon(ctx context.Context, restaurantID primitive.ObjectID, coupon string) (models.Offer, error) {
	offerCollection := models.GetOfferCollection()
	query := bson.M{
		models.KeyOfferCoupon:       coupon,
		models.KeyOfferRestaurantID: restaurantID,
		models.KeyOfferValidFrom: bson.M{
			"$lte": time.Now(),
		},
		models.KeyOfferValidTill: bson.M{
			"$gte": time.Now(),
		},
	}
	var offer models.Offer
	err := offerCollection.FindOne(ctx, query).Decode(&offer)
	return offer, err
}
