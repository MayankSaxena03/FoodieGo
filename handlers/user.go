package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MayankSaxena03/FoodieGo/core"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	user, err := core.GetUserByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Profile Incomplete")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Profile is Complete")
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	user, err := core.GetUserByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var profile models.UserProfile
	err = json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	validate := validator.New()
	err = validate.Struct(profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	update := bson.M{
		models.KeyUserName:     profile.Name,
		models.KeyUserEmail:    profile.Email,
		models.KeyUserBirthday: profile.Birthday,
	}

	err = core.UpdateUserByID(r.Context(), id, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Profile Updated")
}

func AddAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var address models.Address
	err = json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	validate := validator.New()
	err = validate.Struct(address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in Address : " + err.Error())
		return
	}

	err = core.AddAddressInUser(r.Context(), userId, address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Address Added")
}

func SetPrimaryAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	addressID, err := primitive.ObjectIDFromHex(vars["addressId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Address ID")
		return
	}

	err = core.SetPrimaryAddressInUser(r.Context(), userId, addressID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Primary Address Set")
}

func UpdateAddressByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	addressID, err := primitive.ObjectIDFromHex(vars["addressId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Address ID")
		return
	}

	var newAddress models.Address
	err = json.NewDecoder(r.Body).Decode(&newAddress)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	validate := validator.New()
	err = validate.Struct(newAddress)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in Address : " + err.Error())
		return
	}

	err = core.UpdateAddressInUser(r.Context(), userId, addressID, newAddress)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Address Updated")
}

func DeleteAddressByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	addressID, err := primitive.ObjectIDFromHex(vars["addressId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Address ID")
		return
	}

	err = core.DeleteAddressInUser(r.Context(), userId, addressID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Address Deleted")
}

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	addresses, err := core.GetAllAddressesOfUser(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addresses)
}

func GetAddressByID(w http.ResponseWriter, r *http.Request) {
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	addressID, err := primitive.ObjectIDFromHex(vars["addressId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Address ID")
		return
	}

	address, err := core.GetAddressByIDInUser(r.Context(), userId, addressID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	user, err := core.GetUserByID(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	var city string
	if len(user.Address) == 0 {
		city, err = helpers.GetCurrentCity()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Something went wrong")
			return
		}
	} else {
		for _, address := range user.Address {
			if address.ID == user.PrimaryAddressID {
				city = address.City
				break
			}
		}
	}

	restaurants, err := core.GetRestaurantsByCity(r.Context(), city, skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurants)
}

func GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	restaurantID, err := primitive.ObjectIDFromHex(vars["restaurantId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Restaurant ID")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurant)
}

func GetRestaurantMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	vars := mux.Vars(r)
	restaurantID, err := primitive.ObjectIDFromHex(vars["restaurantId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Restaurant ID")
		return
	}

	food, err := core.GetAllFoodByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(food)
}

func GetRestaurantOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	vars := mux.Vars(r)
	restaurantID, err := primitive.ObjectIDFromHex(vars["restaurantId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Restaurant ID")
		return
	}

	offers, err := core.GetValidOffersByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offers)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	cart, err := core.GetCartByUserID(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var food models.Food
	err = json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Food")
		return
	}

	err = core.AddToUserCart(r.Context(), userId, food)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Added to cart")
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	foodID, err := primitive.ObjectIDFromHex(vars["foodId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Food ID")
		return
	}

	err = core.RemoveFromUserCart(r.Context(), userId, foodID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Removed from cart")
}

func ClearCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	err = core.ClearUserCart(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Cart cleared")
}

func AddNotesToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var notes string
	err = json.NewDecoder(r.Body).Decode(&notes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Notes")
		return
	}

	err = core.AddNotesToUserCart(r.Context(), userId, notes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Notes added")
}

func ApplyCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	couponCode := vars["couponCode"]

	err = core.ApplyCouponToUserCart(r.Context(), userId, couponCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Coupon applied")
}

func RemoveCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	err = core.RemoveCouponFromUserCart(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Coupon removed")
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	paymentMethod := r.URL.Query().Get("paymentMethod")

	err = core.CreateOrderForUser(r.Context(), userId, paymentMethod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order created")
	//TODO: Notify Restaurant and Delivery Partners
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	orderID, err := primitive.ObjectIDFromHex(vars["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Order ID")
		return
	}

	err = core.CancelOrderForUser(r.Context(), userId, orderID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order cancelled")
}

func GetAllUserOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orders, err := core.GetAllOrdersForUser(r.Context(), userId, skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func GetUserOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	orderID, err := primitive.ObjectIDFromHex(vars["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Order ID")
		return
	}

	order, err := core.GetOrderByIDForUser(r.Context(), userId, orderID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func GetUserActiveOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orders, err := core.GetActiveOrdersForUser(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func AddRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, _, err := helpers.GetUserDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	orderID, err := primitive.ObjectIDFromHex(vars["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Order ID")
		return
	}

	var rating models.Rating
	err = json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Rating")
		return
	}

	if rating.Rating < 1 || rating.Rating > 5 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Rating should be between 1 and 5")
		return
	}

	err = core.AddRating(r.Context(), userId, orderID, rating.Rating)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Rating added")
}
