package routes

import (
	"github.com/MayankSaxena03/FoodieGo/handlers"
	authHandlers "github.com/MayankSaxena03/FoodieGo/handlers/authentication"
	commonHandlers "github.com/MayankSaxena03/FoodieGo/handlers/common"
	"github.com/MayankSaxena03/FoodieGo/middleware"
	"github.com/gorilla/mux"
)

func RestaurantAuthRoutes(router *mux.Router) {
	authRestaurantGroup := router.PathPrefix("/api/restaurant/auth").Subrouter()
	authRestaurantGroup.HandleFunc("/signup", authHandlers.RestaurantSignUp).Methods("POST")
	authRestaurantGroup.HandleFunc("/login", authHandlers.RestaurantLogin).Methods("POST")
	authRestaurantGroup.HandleFunc("/generate-otp", commonHandlers.GenerateOTPForEmail).Methods("POST")
}

func RestaurantRoutes(router *mux.Router) {
	restaurantGroup := router.PathPrefix("/api/restaurant").Subrouter()
	restaurantGroup.Use(middleware.AuthenticateRestaurant)
	restaurantGroup.HandleFunc("/validate", handlers.ValidateRestaurant).Methods("GET")
	restaurantGroup.HandleFunc("/info", handlers.GetRestaurantInfo).Methods("GET")
	restaurantGroup.HandleFunc("/update", handlers.UpdateRestaurantInfo).Methods("PUT")

	restaurantGroup.HandleFunc("/create-menu", handlers.CreateMenu).Methods("POST")
	restaurantGroup.HandleFunc("/get-menu", handlers.GetMenu).Methods("GET")
	restaurantGroup.HandleFunc("/get-menu-items", handlers.GetMenuItems).Methods("GET")
	restaurantGroup.HandleFunc("/add-menu-item", handlers.CreateMenuItem).Methods("POST")
	restaurantGroup.HandleFunc("/delete-menu-item/{menuItemId}", handlers.DeleteMenuItem).Methods("DELETE")
	restaurantGroup.HandleFunc("/update-menu-item/{menuItemId}", handlers.UpdateMenuItem).Methods("PUT")

	restaurantGroup.HandleFunc("/create-offer", handlers.CreateOffer).Methods("POST")
	restaurantGroup.HandleFunc("/get-offers", handlers.GetOffers).Methods("GET")
	restaurantGroup.HandleFunc("/get-offer/{offerId}", handlers.GetOffer).Methods("GET")
	restaurantGroup.HandleFunc("/update-offer/{offerId}", handlers.UpdateOffer).Methods("PUT")
	restaurantGroup.HandleFunc("/delete-offer/{offerId}", handlers.DeleteOffer).Methods("DELETE")

	restaurantGroup.HandleFunc("/pending-orders", handlers.GetPendingOrders).Methods("GET")
	restaurantGroup.HandleFunc("/accepted-orders", handlers.GetAcceptedOrders).Methods("GET")
	restaurantGroup.HandleFunc("/orders", handlers.GetAllRestaurantOrders).Methods("GET")
	restaurantGroup.HandleFunc("/order/{orderId}", handlers.GetRestaurantOrderDetails).Methods("GET")
	restaurantGroup.HandleFunc("/accept-order/{orderId}", handlers.AcceptOrder).Methods("PUT")
	restaurantGroup.HandleFunc("/reject-order/{orderId}", handlers.RejectOrder).Methods("PUT")

	restaurantGroup.HandleFunc("/income", handlers.GetRestaurantIncome).Methods("GET")
}
