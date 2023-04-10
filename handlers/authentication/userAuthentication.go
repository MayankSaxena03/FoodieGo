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

func UserLogin(w http.ResponseWriter, r *http.Request) {
	//This function will handle both login and signup
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
	var user models.User
	userCollection := models.GetUserCollection()
	query := bson.M{
		models.KeyUserPhone: otpDetails.Phone,
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyUserLastLoginOn: time.Now(),
		},
	}
	err = userCollection.FindOneAndUpdate(r.Context(), query, update).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err == mongo.ErrNoDocuments {
		//User not found, create a new user
		user := models.User{
			Phone: otpDetails.Phone,
		}
		result, err := core.CreateUser(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.ID = result.ID
		user.Phone = otpDetails.Phone
	}

	token, refreshToken, err := helpers.GenerateAllTokensForUser(user.ID, user.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
}
