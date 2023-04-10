package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MayankSaxena03/FoodieGo/core"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateDeliveryPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	deliveryPartner, err := core.GetDeliveryPartnerByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	validate := validator.New()
	err = validate.Struct(deliveryPartner)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Profile Incomplete")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Profile is Complete")
}

func GetDeliveryPartnerInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	deliveryPartner, err := core.GetDeliveryPartnerByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deliveryPartner)
}

func UpdateDeliveryPartnerInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var deliveryPartner models.DeliveryPartnerProfile
	err = json.NewDecoder(r.Body).Decode(&deliveryPartner)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	err = core.UpdateDeliveryPartnerInfo(r.Context(), deliveryPartnerId, deliveryPartner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Profile Updated")
}

func GetIncomingOrdersForDeliveryPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)

	orders, err := core.GetIncomingOrdersForDelivery(r.Context(), deliveryPartnerId, skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func AcceptOrderForDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderID, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	err = core.AddDeliveryPartnerToOrder(r.Context(), orderID, deliveryPartnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order Accepted")
}

func GetOngoingOrdersForDeliveryPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}
	skip, limit := helpers.GetSkipAndLimitFromRequest(r)

	orders, err := core.GetOngoingOrdersForDelivery(r.Context(), deliveryPartnerId, skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func CompleteOrderForDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderID, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	err = core.CompleteOrderForDelivery(r.Context(), orderID, deliveryPartnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order Completed")
}

func PickupOrderForDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderID, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	err = core.PickupOrderForDelivery(r.Context(), orderID, deliveryPartnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order Picked Up")
}

func GetOrderDetailsForDeliveryPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deliveryPartnerId, _, err := helpers.GetDeliveryPartnerDetailsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	orderID, err := primitive.ObjectIDFromHex(mux.Vars(r)["orderId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	order, err := core.GetOrderDetailsForDeliveryPartner(r.Context(), orderID, deliveryPartnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
