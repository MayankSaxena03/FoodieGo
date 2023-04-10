package helpers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MayankSaxena03/FoodieGo/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserDetailsFromRequest(r *http.Request) (primitive.ObjectID, string, error) {
	userId, err := primitive.ObjectIDFromHex(r.Header.Get("userId"))
	if err != nil {
		return primitive.NilObjectID, "", err
	}

	userPhone := r.Header.Get("phone")
	if userPhone == "" {
		return primitive.NilObjectID, "", errors.New("user phone not found")
	}

	return userId, userPhone, nil
}

func GetRestaurantDetailsFromRequest(r *http.Request) (primitive.ObjectID, string, error) {
	restaurantId, err := primitive.ObjectIDFromHex(r.Header.Get("restaurantId"))
	if err != nil {
		return primitive.NilObjectID, "", err
	}

	restaurantEmail := r.Header.Get("email")
	if restaurantEmail == "" {
		return primitive.NilObjectID, "", errors.New("restaurant Email not found")
	}

	return restaurantId, restaurantEmail, nil
}

func GetDeliveryPartnerDetailsFromRequest(r *http.Request) (primitive.ObjectID, string, error) {
	partnerId, err := primitive.ObjectIDFromHex(r.Header.Get("deliveryPartnerId"))
	if err != nil {
		return primitive.NilObjectID, "", err
	}

	partnerPhone := r.Header.Get("phone")
	if partnerPhone == "" {
		return primitive.NilObjectID, "", errors.New("delivery Partner Phone not found")
	}

	return partnerId, partnerPhone, nil
}

func GetSkipAndLimitFromRequest(r *http.Request) (int64, int64) {
	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	if skip == "" {
		skip = strconv.Itoa(constants.DefaultSkip)
	}

	if limit == "" {
		limit = strconv.Itoa(constants.DefaultLimit)
	}

	skipInt, _ := strconv.ParseInt(skip, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)

	return skipInt, limitInt
}
