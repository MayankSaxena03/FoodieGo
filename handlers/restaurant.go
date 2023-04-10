package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MayankSaxena03/FoodieGo/core"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Profile Incomplete")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Profile is Complete")
}

func GetRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurant)
}

func UpdateRestaurantInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	email := restaurant.Email
	rating := restaurant.Rating
	totalRating := restaurant.TotalRating
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}
	if email != restaurant.Email {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email cannot be changed" + err.Error())
		return
	}
	if rating != restaurant.Rating || totalRating != restaurant.TotalRating {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Rating cannot be changed")
		return
	}

	validate := validator.New()
	err = validate.Struct(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	err = core.UpdateRestaurant(r.Context(), restaurantID, restaurant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Profile Updated")
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	if restaurant.MenuID != primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Menu Already Exists")
		return
	}

	menu := models.Menu{}
	err = json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	menu.RestaurantID = restaurantID

	validate := validator.New()
	err = validate.Struct(menu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	err = core.CreateMenu(r.Context(), menu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Menu Created")
}

func CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	if restaurant.MenuID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Menu Does Not Exist. Create Menu First")
		return
	}

	var food models.Food
	err = json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	err = core.CreateMenuItem(r.Context(), restaurant.MenuID, food)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Food Item Created")
}

func DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	foodID, err := primitive.ObjectIDFromHex(vars["menuItemId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Food ID")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	if restaurant.MenuID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Menu Does Not Exist. Create Menu First")
		return
	}

	var food models.Food
	err = json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	err = core.DeleteMenuItem(r.Context(), restaurant.MenuID, foodID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Food Item Deleted")
}

func UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	foodID, err := primitive.ObjectIDFromHex(vars["menuItemId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Food ID")
		return
	}

	restaurant, err := core.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	if restaurant.MenuID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Menu Does Not Exist. Create Menu First")
		return
	}

	var food models.Food
	err = json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(food)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	err = core.UpdateFood(r.Context(), foodID, food)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Menu Item Updated")
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	menu, err := core.GetMenuByRestaurantID(r.Context(), restaurantID)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menu)
}

func GetMenuItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	items, err := core.GetAllFoodByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var offer models.Offer
	err = json.NewDecoder(r.Body).Decode(&offer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(offer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}
	offer.RestaurantID = restaurantID
	offer.CreatedOn = time.Now()

	err = core.CreateOffer(r.Context(), restaurantID, offer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Offer Created")
}

func GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	offers, err := core.GetAllOffersByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offers)
}

func GetOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	offerId, err := primitive.ObjectIDFromHex(mux.Vars(r)["offerId"])
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Offer Id")
		return
	}

	offer, err := core.GetOfferByID(r.Context(), offerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offer)
}

func UpdateOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantId, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	offerId, err := primitive.ObjectIDFromHex(mux.Vars(r)["offerId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Offer Id")
		return
	}

	var offer models.Offer
	err = json.NewDecoder(r.Body).Decode(&offer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(offer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request: " + err.Error())
		return
	}

	if offer.RestaurantID != restaurantId {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	offer.UpdatedOn = time.Now()

	err = core.UpdateOfferByID(r.Context(), offerId, offer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Offer Updated")
}

func DeleteOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantId, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	offerId, err := primitive.ObjectIDFromHex(mux.Vars(r)["offerId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Offer Id")
		return
	}

	offer, err := core.GetOfferByID(r.Context(), offerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	if offer.RestaurantID != restaurantId {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	err = core.DeleteOfferByID(r.Context(), offerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Offer Deleted")
}

func GetPendingOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orders, err := core.GetAllPendingOrdersByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func GetAcceptedOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orders, err := core.GetAllAcceptedOrdersByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func GetAllRestaurantOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orders, err := core.GetAllOrdersByRestaurantID(r.Context(), restaurantID, skip, limit)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func GetRestaurantOrderDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderId, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Order Id")
		return
	}

	order, err := core.GetOrderDetailsByRestaurantID(r.Context(), orderId, restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func AcceptOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderId, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Order Id")
		return
	}

	order, err := core.GetOrderDetailsByRestaurantID(r.Context(), orderId, restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	if order.OrderStatus != "Pending" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Order is not in pending state")
		return
	}

	err = core.AcceptOrderByRestaurantID(r.Context(), orderId, restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order Accepted")
}

func RejectOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderId, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Order Id")
		return
	}

	order, err := core.GetOrderDetailsByRestaurantID(r.Context(), orderId, restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	if order.OrderStatus != "Pending" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Order is not in pending state")
		return
	}

	err = core.RejectOrderByRestaurantID(r.Context(), orderId, restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order Rejected")
}

func GetRestaurantIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurantID, _, err := helpers.GetRestaurantDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	income, err := core.GetRestaurantIncome(r.Context(), restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong" + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(income)
}
