package routes

import (
	"github.com/MayankSaxena03/FoodieGo/handlers"
	authHandlers "github.com/MayankSaxena03/FoodieGo/handlers/authentication"
	commonHandlers "github.com/MayankSaxena03/FoodieGo/handlers/common"
	"github.com/MayankSaxena03/FoodieGo/middleware"
	"github.com/gorilla/mux"
)

func DeliveryPartnerAuthRoutes(router *mux.Router) {
	authDeliveryPartnerGroup := router.PathPrefix("/api/delivery-partner/auth").Subrouter()
	authDeliveryPartnerGroup.HandleFunc("/login", authHandlers.DeliveryPartnerLogin).Methods("POST")
	authDeliveryPartnerGroup.HandleFunc("/generate-otp", commonHandlers.GenerateOTPForEmail).Methods("POST")
}

func DeliveryPartnerRoutes(router *mux.Router) {
	deliveryPartnerGroup := router.PathPrefix("/api/delivery-partner").Subrouter()
	deliveryPartnerGroup.Use(middleware.AuthenticateDeliveryPartner)
	deliveryPartnerGroup.HandleFunc("/validate", handlers.ValidateDeliveryPartner).Methods("GET")
	deliveryPartnerGroup.HandleFunc("/me", handlers.GetDeliveryPartnerInfo).Methods("GET")
	deliveryPartnerGroup.HandleFunc("/update", handlers.UpdateDeliveryPartnerInfo).Methods("PUT")

	deliveryPartnerGroup.HandleFunc("/incoming-orders", handlers.GetIncomingOrdersForDeliveryPartner).Methods("GET")
	deliveryPartnerGroup.HandleFunc("/accept-order/{orderID}", handlers.AcceptOrderForDelivery).Methods("POST")
	deliveryPartnerGroup.HandleFunc("/ongoing-orders", handlers.GetOngoingOrdersForDeliveryPartner).Methods("GET")
	deliveryPartnerGroup.HandleFunc("/pickup-order/{orderID}", handlers.PickupOrderForDelivery).Methods("POST")
	deliveryPartnerGroup.HandleFunc("/complete-order/{orderID}", handlers.CompleteOrderForDelivery).Methods("POST")
	deliveryPartnerGroup.HandleFunc("/order/{orderID}", handlers.GetOrderDetailsForDeliveryPartner).Methods("GET")
}
