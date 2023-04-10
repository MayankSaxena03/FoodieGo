package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MayankSaxena03/FoodieGo/constants"
	"github.com/MayankSaxena03/FoodieGo/core"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	redisModels "github.com/MayankSaxena03/FoodieGo/models/redis"
)

func GenerateOTPForPhone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginOTP redisModels.LoginSignupBody
	err := json.NewDecoder(r.Body).Decode(&loginOTP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}
	phone := loginOTP.Phone
	isValid := helpers.ValidatePhoneNumber(phone)
	if !isValid {
		json.NewEncoder(w).Encode("Invalid phone number")
		return
	}

	otp := helpers.RandomOTP(constants.DefaultOTPLength)

	err = core.SetLoginOTPForPhoneInRedis(phone, otp)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	json.NewEncoder(w).Encode("OTP generated successfully")
}

func GenerateOTPForEmail(w http.ResponseWriter, r *http.Request) {
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

	otp := helpers.RandomOTP(constants.DefaultOTPLength)

	err = core.SetLoginOTPForEmailInRedis(email, otp)
	if err != nil {
		json.NewEncoder(w).Encode("Something went wrong : " + err.Error())
		return
	}

	json.NewEncoder(w).Encode("OTP generated successfully")
}
