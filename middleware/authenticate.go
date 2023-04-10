package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/MayankSaxena03/FoodieGo/helpers"
)

func AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized: Invalid Token")
			return
		}
		// Validate the token
		claims, err := helpers.ValidateUserToken(token)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized: Invalid Token")
			return
		}

		r.Header.Set("userId", claims.UserID.Hex())
		r.Header.Set("phone", claims.UserPhone)
		next.ServeHTTP(w, r)
	})
}

func AuthenticateRestaurant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized: Invalid Token")
			return
		}
		// Validate the token
		claims, err := helpers.ValidateRestaurantToken(token)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized: Invalid Token")
			return
		}

		r.Header.Set("restaurantId", claims.RestaurantID.Hex())
		r.Header.Set("email", claims.RestaurantEmail)
		next.ServeHTTP(w, r)
	})
}

func AuthenticateDeliveryPartner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized: Invalid Token")
			return
		}
		// Validate the token
		claims, err := helpers.ValidateDeliveryPartnerToken(token)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized: Invalid Token")
			return
		}

		r.Header.Set("deliveryPartnerId", claims.PartnerID.Hex())
		r.Header.Set("phone", claims.PartnerPhone)
		next.ServeHTTP(w, r)
	})
}
