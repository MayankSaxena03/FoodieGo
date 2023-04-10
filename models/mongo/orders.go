package models

import (
	"time"

	"github.com/MayankSaxena03/FoodieGo/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KeyOrderID = "_id"
var KeyOrderUserID = "userId"
var KeyOrderStatus = "orderStatus"
var KeyOrderRestaurantID = "restaurantId"
var KeyOrderPriceAfterOffer = "priceAfterOffer"
var KeyOrderDeliveryAddress = "deliveryAddress"
var KeyOrderDeliveryPartnerID = "deliveryPartnerId"

type Order struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID            primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	RestaurantID      primitive.ObjectID `json:"restaurantId,omitempty" bson:"restaurantId,omitempty"`
	DeliveryAddress   Address            `json:"deliveryAddress,omitempty" bson:"deliveryAddress,omitempty"`
	Items             []FoodWithQuantity `json:"items,omitempty" bson:"items,omitempty"`
	Coupon            string             `json:"coupon,omitempty" bson:"coupon,omitempty"`
	TotalPrice        float64            `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	PriceAfterOffer   float64            `json:"priceAfterOffer,omitempty" bson:"priceAfterOffer,omitempty"`
	PaymentMethod     string             `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty"` // COD, Online
	Notes             string             `json:"notes,omitempty" bson:"notes,omitempty"`
	OrderStatus       string             `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	DeliveryPartnerID primitive.ObjectID `json:"deliveryPartner,omitempty" bson:"deliveryPartner,omitempty"`
	OrderTime         time.Time          `json:"orderTime,omitempty" bson:"orderTime,omitempty"`
}

type OrderDetailsForUser struct {
	ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RestaurantName       string             `json:"restaurantName,omitempty" bson:"restaurantName,omitempty"`
	RestaurantAddress    Address            `json:"restaurantAddress,omitempty" bson:"restaurantAddress,omitempty"`
	RestaurantPhone      string             `json:"restaurantPhone,omitempty" bson:"restaurantPhone,omitempty"` // TODO: Remove this
	DeliveryAddress      Address            `json:"deliveryAddress,omitempty" bson:"deliveryAddress,omitempty"`
	Items                []FoodWithQuantity `json:"items,omitempty" bson:"items,omitempty"`
	Price                float64            `json:"priceAfterOffer,omitempty" bson:"priceAfterOffer,omitempty"`
	PaymentMethod        string             `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty"` // COD, Online
	OrderStatus          string             `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	DeliveryPartnerName  string             `json:"deliveryPartnerName,omitempty" bson:"deliveryPartnerName,omitempty"`
	DeliveryPartnerPhone string             `json:"deliveryPartnerPhone,omitempty" bson:"deliveryPartnerPhone,omitempty"`
	OrderTime            time.Time          `json:"orderTime,omitempty" bson:"orderTime,omitempty"`
}

type OrderDetailsForRestaurant struct {
	ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID               primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	DeliveryPartnerName  string             `json:"deliveryPartnerName,omitempty" bson:"deliveryPartnerName,omitempty"`
	DeliveryPartnerPhone string             `json:"deliveryPartnerPhone,omitempty" bson:"deliveryPartnerPhone,omitempty"`
	Items                []FoodWithQuantity `json:"items,omitempty" bson:"items,omitempty"`
	Notes                string             `json:"notes,omitempty" bson:"notes,omitempty"`
	Price                float64            `json:"price,omitempty" bson:"price,omitempty"`
	OrderStatus          string             `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
}

type OrderDetailsForDeliveryPartner struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username          string             `json:"username,omitempty" bson:"username,omitempty"`
	UserPhone         string             `json:"userPhone,omitempty" bson:"userPhone,omitempty"`
	DeliveryAddress   Address            `json:"deliveryAddress,omitempty" bson:"deliveryAddress,omitempty"`
	RestaurantName    string             `json:"restaurantName,omitempty" bson:"restaurantName,omitempty"`
	RestaurantAddress Address            `json:"restaurantAddress,omitempty" bson:"restaurantAddress,omitempty"`
	RestaurantPhone   string             `json:"restaurantPhone,omitempty" bson:"restaurantPhone,omitempty"`
	Items             []FoodWithQuantity `json:"items,omitempty" bson:"items,omitempty"`
	PaymentMethod     string             `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty"` // COD, Online
	Price             float64            `json:"price,omitempty" bson:"price,omitempty"`
	OrderStatus       string             `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
}

func GetOrderCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "orders")
}
