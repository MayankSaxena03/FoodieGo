package routes

import (
	"github.com/MayankSaxena03/FoodieGo/handlers"
	authHandlers "github.com/MayankSaxena03/FoodieGo/handlers/authentication"
	commonHandlers "github.com/MayankSaxena03/FoodieGo/handlers/common"
	"github.com/MayankSaxena03/FoodieGo/middleware"
	"github.com/gorilla/mux"
)

func UserAuthRoutes(router *mux.Router) {
	authUserGroup := router.PathPrefix("/api/user/auth").Subrouter()
	authUserGroup.HandleFunc("/login", authHandlers.UserLogin).Methods("POST")
	authUserGroup.HandleFunc("/generate-otp", commonHandlers.GenerateOTPForEmail).Methods("POST")
}

func UserRoutes(router *mux.Router) {
	userGroup := router.PathPrefix("/api/user").Subrouter()
	userGroup.Use(middleware.AuthenticateUser)
	userGroup.HandleFunc("/validate", handlers.ValidateUser).Methods("GET")
	userGroup.HandleFunc("/me", handlers.GetUserInfo).Methods("GET")
	userGroup.HandleFunc("/update", handlers.UpdateUserInfo).Methods("PUT")
	userGroup.HandleFunc("/address", handlers.AddAddress).Methods("POST")
	userGroup.HandleFunc("/primary-address/{addressId}", handlers.SetPrimaryAddress).Methods("PUT")
	userGroup.HandleFunc("/address/{addressId}", handlers.UpdateAddressByID).Methods("PUT")
	userGroup.HandleFunc("/address/{addressId}", handlers.DeleteAddressByID).Methods("DELETE")
	userGroup.HandleFunc("/address", handlers.GetAddresses).Methods("GET")
	userGroup.HandleFunc("/address/{addressId}", handlers.GetAddressByID).Methods("GET")

	userGroup.HandleFunc("/restaurants", handlers.GetRestaurants).Methods("GET")
	userGroup.HandleFunc("/restaurants/{restaurantId}", handlers.GetRestaurantByID).Methods("GET")
	userGroup.HandleFunc("/restaurants/{restaurantId}/menu", handlers.GetRestaurantMenu).Methods("GET")
	userGroup.HandleFunc("/restaurants/offers/{restaurantId}", handlers.GetRestaurantOffers).Methods("GET")

	userGroup.HandleFunc("/cart", handlers.GetCart).Methods("GET")
	userGroup.HandleFunc("/cart", handlers.AddToCart).Methods("POST")
	userGroup.HandleFunc("/cart/{foodId}", handlers.RemoveFromCart).Methods("DELETE")
	userGroup.HandleFunc("/clear-cart", handlers.ClearCart).Methods("DELETE")
	userGroup.HandleFunc("/cart/notes", handlers.AddNotesToCart).Methods("POST")
	userGroup.HandleFunc("/apply-coupon/{couponCode}", handlers.ApplyCoupon).Methods("POST")
	userGroup.HandleFunc("/remove-coupon", handlers.RemoveCoupon).Methods("DELETE")

	userGroup.HandleFunc("/create-order", handlers.CreateOrder).Methods("POST")
	userGroup.HandleFunc("/cancel-order/{orderId}", handlers.CancelOrder).Methods("POST")
	userGroup.HandleFunc("/orders", handlers.GetAllUserOrders).Methods("GET")
	userGroup.HandleFunc("/orders/{orderId}", handlers.GetUserOrder).Methods("GET")
	userGroup.HandleFunc("/active-orders", handlers.GetUserActiveOrders).Methods("GET")

	userGroup.HandleFunc("/rating/{orderID}", handlers.AddRating).Methods("POST")
}
