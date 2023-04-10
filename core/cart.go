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

func GetCartByUserID(ctx context.Context, userID primitive.ObjectID) (models.Cart, error) {
	var cart models.Cart
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userID,
	}
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	return cart, err
}

func AddToUserCart(ctx context.Context, userId primitive.ObjectID, food models.Food) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	menu, err := GetMenuByFoodID(ctx, food.ID)

	if err == mongo.ErrNoDocuments {
		cart = models.Cart{
			UserID:       userId,
			RestaurantID: menu.RestaurantID,
			Items:        []models.FoodWithQuantity{},
		}
	} else {
		if cart.RestaurantID != menu.RestaurantID && cart.RestaurantID != primitive.NilObjectID {
			return errors.New("cannot add food from different restaurants to cart! Clear Cart First")
		}
	}

	isIncremented := false
	for i, item := range cart.Items {
		if item.ID == food.ID {
			cart.Items[i].Quantity += 1
			cart.TotalPrice += food.Price
			isIncremented = true
			break
		}
	}

	if !isIncremented {
		cart.Items = append(cart.Items, models.FoodWithQuantity{
			Food:     food,
			Quantity: 1,
		})
		cart.TotalPrice += food.Price
	}

	if cart.Coupon != "" {
		offer, err := GetValidOfferByCoupon(ctx, cart.RestaurantID, cart.Coupon)
		if err != nil {
			return err
		}
		cart.PriceAfterDiscount = cart.TotalPrice - (cart.TotalPrice * float64(offer.Discount/100))
		if cart.PriceAfterDiscount < cart.TotalPrice-offer.MaxDiscount && offer.MaxDiscount != 0 {
			cart.PriceAfterDiscount = cart.TotalPrice - offer.MaxDiscount
		}
	}

	cart.UpdatedOn = time.Now()
	query = bson.M{
		models.KeyCartUserID: userId,
	}
	update := bson.M{
		constants.MongoKeywordSet: cart,
	}
	options := options.FindOneAndUpdate().SetUpsert(true)

	result := cartCollection.FindOneAndUpdate(ctx, query, update, options)
	return result.Err()
}

func RemoveFromUserCart(ctx context.Context, userId primitive.ObjectID, foodId primitive.ObjectID) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	if err != nil {
		return err
	}

	for i, item := range cart.Items {
		if item.ID == foodId {
			if item.Quantity > 1 {
				cart.Items[i].Quantity -= 1
			} else {
				if i == len(cart.Items)-1 {
					cart.Items = cart.Items[:i]
				} else {
					cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
				}
			}
			cart.TotalPrice -= item.Price
			break
		}
	}

	if len(cart.Items) == 0 {
		return ClearUserCart(ctx, userId)
	}

	if cart.Coupon != "" {
		offer, err := GetValidOfferByCoupon(ctx, cart.RestaurantID, cart.Coupon)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		if err != mongo.ErrNoDocuments {
			cart.PriceAfterDiscount = cart.TotalPrice - (cart.TotalPrice * float64(offer.Discount/100))
			if cart.PriceAfterDiscount < cart.TotalPrice-offer.MaxDiscount && offer.MaxDiscount != 0 {
				cart.PriceAfterDiscount = cart.TotalPrice - offer.MaxDiscount
			}
		}
	}

	cart.UpdatedOn = time.Now()
	query = bson.M{
		models.KeyCartUserID: userId,
	}
	update := bson.M{
		constants.MongoKeywordSet: cart,
	}

	err = cartCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}

func ClearUserCart(ctx context.Context, userId primitive.ObjectID) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	err := cartCollection.FindOneAndDelete(ctx, query).Err()
	return err
}

func AddNotesToUserCart(ctx context.Context, userId primitive.ObjectID, notes string) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	if err != nil {
		return err
	}

	cart.Notes = notes
	cart.UpdatedOn = time.Now()
	query = bson.M{
		models.KeyCartUserID: userId,
	}
	update := bson.M{
		constants.MongoKeywordSet: cart,
	}

	err = cartCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}

func ApplyCouponToUserCart(ctx context.Context, userId primitive.ObjectID, coupon string) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	if err != nil {
		return err
	}

	offer, err := GetValidOfferByCoupon(ctx, cart.RestaurantID, coupon)
	if err != nil {
		return err
	}

	cart.Coupon = coupon
	cart.PriceAfterDiscount = cart.TotalPrice - (cart.TotalPrice * float64(offer.Discount/100))
	if cart.PriceAfterDiscount < cart.TotalPrice-offer.MaxDiscount && offer.MaxDiscount != 0 {
		cart.PriceAfterDiscount = cart.TotalPrice - offer.MaxDiscount
	}

	cart.UpdatedOn = time.Now()
	query = bson.M{
		models.KeyCartUserID: userId,
	}
	update := bson.M{
		constants.MongoKeywordSet: cart,
	}

	err = cartCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}

func RemoveCouponFromUserCart(ctx context.Context, userId primitive.ObjectID) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	if err != nil {
		return err
	}

	cart.Coupon = ""
	cart.PriceAfterDiscount = cart.TotalPrice

	cart.UpdatedOn = time.Now()
	query = bson.M{
		models.KeyCartUserID: userId,
	}
	update := bson.M{
		constants.MongoKeywordSet: cart,
	}

	err = cartCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}
