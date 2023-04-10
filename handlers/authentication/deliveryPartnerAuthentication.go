package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MayankSaxena03/FoodieGo/constants"
	"github.com/MayankSaxena03/FoodieGo/core"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	redisModels "github.com/MayankSaxena03/FoodieGo/models/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeliveryPartnerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var otpDetails redisModels.LoginSignupBody
	err := json.NewDecoder(r.Body).Decode(&otpDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	otp, err := core.GetLoginOTPForPhoneFromRedis(otpDetails.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	if otpDetails.OTP != otp {
		json.NewEncoder(w).Encode("Invalid or expired OTP")
		return
	}

	_ = core.DeleteLoginOTPForPhoneInRedis(otpDetails.Phone)
	var deliveryPartner models.DeliveryPartner
	deliveryPartnerCollection := models.GetDeliveryPartnerCollection()
	query := bson.M{
		models.KeyDeliveryPartnerPhone: otpDetails.Phone,
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyDeliveryPartnerLastLoginOn: time.Now(),
		},
	}
	err = deliveryPartnerCollection.FindOneAndUpdate(r.Context(), query, update).Decode(&deliveryPartner)
	if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err == mongo.ErrNoDocuments {
		//User not found, create a new user
		deliveryPartner := models.DeliveryPartner{
			Phone: otpDetails.Phone,
		}
		result, err := core.CreateDeliveryPartner(r.Context(), deliveryPartner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		deliveryPartner.ID = result.ID
		deliveryPartner.Phone = otpDetails.Phone

	}

	token, refreshToken, err := helpers.GenerateAllTokensForDeliveryPartner(deliveryPartner.ID, deliveryPartner.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
}
