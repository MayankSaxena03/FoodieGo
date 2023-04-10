package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MayankSaxena03/FoodieGo/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	routes.UserAuthRoutes(router)
	routes.UserRoutes(router)
	routes.RestaurantAuthRoutes(router)
	routes.RestaurantRoutes(router)
	routes.DeliveryPartnerAuthRoutes(router)
	routes.DeliveryPartnerRoutes(router)

	fmt.Println("Server is running on port " + port)
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, router))
}
