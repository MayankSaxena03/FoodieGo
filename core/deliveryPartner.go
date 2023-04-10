package core

import (
	"context"
	"time"

	"github.com/MayankSaxena03/FoodieGo/constants"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDeliveryPartner(ctx context.Context, deliveryPartner models.DeliveryPartner) (models.DeliveryPartner, error) {
	deliveryPartnerCollection := models.GetDeliveryPartnerCollection()
	deliveryPartner.CreatedOn = time.Now()
	deliveryPartner.LastLoginOn = time.Now()
	result, err := deliveryPartnerCollection.InsertOne(ctx, deliveryPartner)
	if err != nil {
		return models.DeliveryPartner{}, err
	}

	deliveryPartner.ID = result.InsertedID.(primitive.ObjectID)
	return deliveryPartner, nil
}

func GetDeliveryPartnerByID(ctx context.Context, id primitive.ObjectID) (models.DeliveryPartner, error) {
	var deliveryPartner models.DeliveryPartner
	deliveryPartnerCollection := models.GetDeliveryPartnerCollection()
	query := primitive.M{
		models.KeyDeliveryPartnerID: id,
	}
	err := deliveryPartnerCollection.FindOne(ctx, query).Decode(&deliveryPartner)
	if err != nil {
		return models.DeliveryPartner{}, err
	}

	return deliveryPartner, nil
}

func UpdateDeliveryPartnerInfo(ctx context.Context, id primitive.ObjectID, deliveryPartner models.DeliveryPartnerProfile) error {
	deliveryPartnerCollection := models.GetDeliveryPartnerCollection()
	query := primitive.M{
		models.KeyDeliveryPartnerID: id,
	}
	update := primitive.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyDeliveryPartnerName:         deliveryPartner.Name,
			models.KeyDeliveryPartnerEmail:        deliveryPartner.Email,
			models.KeyDeliveryPartnerCity:         deliveryPartner.City,
			models.KeyDeliveryPartnerBirthday:     deliveryPartner.Birthday,
			models.KeyDeliveryPartnerProfileImage: deliveryPartner.ProfileImage,
		},
	}
	var oldDoc models.DeliveryPartner
	err1 := deliveryPartnerCollection.FindOne(ctx, query).Decode(&oldDoc)

	result := deliveryPartnerCollection.FindOneAndUpdate(ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return result.Err()
	}
	if err1 == nil {
		var newDoc models.DeliveryPartner
		result.Decode(&newDoc)
		_ = CreateChangeLog(context.Background(), oldDoc, newDoc, oldDoc.ID, constants.EntityTypeDeliveryPartner)
	}
	return nil
}
