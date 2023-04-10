package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MayankSaxena03/FoodieGo/core"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	redisModels "github.com/MayankSaxena03/FoodieGo/models/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RestaurantSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginOTP redisModels.LoginSignupBody
	err := json.NewDecoder(r.Body).Decode(&loginOTP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	email := loginOTP.Email
	isValid := helpers.ValidateEmail(email)
	if !isValid {
		json.NewEncoder(w).Encode("Invalid email")
		return
	}

	otp, err := core.GetLoginOTPForEmailFromRedis(email)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	if otp != loginOTP.OTP {
		json.NewEncoder(w).Encode("Invalid OTP")
		return
	}

	exists, err := core.CheckIfRestaurantExists(r.Context(), email)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}
	if exists {
		json.NewEncoder(w).Encode("Restaurant already exists")
		return
	}

	var restaurant models.Restaurant
	restaurant.Email = email
	restaurant.Password, err = helpers.HashPassword(loginOTP.Password)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	result, err := core.CreateRestaurant(r.Context(), restaurant)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	token, refreshToken, err := helpers.GenerateAllTokensForRestaurant(result.InsertedID.(primitive.ObjectID), loginOTP.Email)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	json.NewEncoder(w).Encode(models.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
}

func RestaurantLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginOTP redisModels.LoginSignupBody
	err := json.NewDecoder(r.Body).Decode(&loginOTP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	email := loginOTP.Email
	password := loginOTP.Password

	isValid := helpers.ValidateEmail(email)
	if !isValid {
		json.NewEncoder(w).Encode("Invalid email")
		return
	}

	restaurant, err := core.GetRestaurantByEmail(r.Context(), email)
	if err != nil && err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}
	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode("Restaurant does not exist")
		return
	}

	if helpers.CheckPasswordHash(password, restaurant.Password) {
		token, refreshToken, err := helpers.GenerateAllTokensForRestaurant(restaurant.ID, restaurant.Email)
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
			return
		}

		json.NewEncoder(w).Encode(models.LoginResponse{
			AccessToken:  token,
			RefreshToken: refreshToken,
		})
	} else {
		json.NewEncoder(w).Encode("Invalid password")
	}
}
