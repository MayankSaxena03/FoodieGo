package core

import (
	"context"
	"errors"
	"time"

	"github.com/MayankSaxena03/FoodieGo/constants"
	"github.com/MayankSaxena03/FoodieGo/helpers"
	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOrderByID(ctx context.Context, id primitive.ObjectID) (models.Order, error) {
	collection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID: id,
	}
	var order models.Order
	err := collection.FindOne(ctx, query).Decode(&order)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func CreateOrderForUser(ctx context.Context, userId primitive.ObjectID, paymentMethod string) error {
	cartCollection := models.GetCartCollection()
	query := bson.M{
		models.KeyCartUserID: userId,
	}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, query).Decode(&cart)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if err == mongo.ErrNoDocuments {
		return errors.New("cart is empty")
	}

	userCollection := models.GetUserCollection()
	var user models.User
	query = bson.M{
		models.KeyUserID: userId,
	}
	err = userCollection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		return err
	}

	if user.Address == nil || len(user.Address) == 0 {
		return errors.New("add a delivery address to place an order")
	}

	var deliveryAddress models.Address
	for _, address := range user.Address {
		if address.ID == user.PrimaryAddressID {
			deliveryAddress = address
			break
		}
	}

	restaurant, err := GetRestaurantByID(ctx, cart.RestaurantID)
	if err != nil {
		return err
	}

	if restaurant.Address.City != deliveryAddress.City {
		return errors.New("delivery for this restaurant is not available in your city")
	}

	if cart.PriceAfterDiscount == 0 {
		cart.PriceAfterDiscount = cart.TotalPrice
	}

	orderCollection := models.GetOrderCollection()
	order := models.Order{
		UserID:          userId,
		RestaurantID:    cart.RestaurantID,
		DeliveryAddress: deliveryAddress,
		Items:           cart.Items,
		Coupon:          cart.Coupon,
		TotalPrice:      cart.TotalPrice,
		PriceAfterOffer: cart.PriceAfterDiscount,
		Notes:           cart.Notes,
		OrderStatus:     "Pending",
		OrderTime:       time.Now(),
		PaymentMethod:   paymentMethod,
	}

	_, err = orderCollection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	err = ClearUserCart(ctx, userId)
	if err != nil {
		return err
	}

	return err
}

func CancelOrderForUser(ctx context.Context, userId primitive.ObjectID, orderId primitive.ObjectID) error {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:     orderId,
		models.KeyOrderUserID: userId,
		models.KeyOrderStatus: "Pending",
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyOrderStatus: constants.KeyOrderCancelled,
		},
	}
	_, err := orderCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil
}

func GetAllOrdersForUser(ctx context.Context, userId primitive.ObjectID, skip int64, limit int64) ([]models.Order, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderUserID: userId,
	}
	cursor, err := orderCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderByIDForUser(ctx context.Context, userId primitive.ObjectID, orderId primitive.ObjectID) (models.OrderDetailsForUser, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:     orderId,
		models.KeyOrderUserID: userId,
	}
	var order models.Order
	err := orderCollection.FindOne(ctx, query).Decode(&order)
	if err != nil {
		return models.OrderDetailsForUser{}, err
	}

	restaurant, err := GetRestaurantByID(ctx, order.RestaurantID)
	if err != nil {
		return models.OrderDetailsForUser{}, err
	}

	orderDetails := models.OrderDetailsForUser{
		ID:                order.ID,
		RestaurantName:    restaurant.Name,
		RestaurantPhone:   restaurant.Phone,
		RestaurantAddress: restaurant.Address,
		Items:             order.Items,
		DeliveryAddress:   order.DeliveryAddress,
	}

	var deliveryPartner models.DeliveryPartner
	if order.DeliveryPartnerID != primitive.NilObjectID {
		deliveryPartner, err = GetDeliveryPartnerByID(ctx, order.DeliveryPartnerID)
		if err != nil {
			return models.OrderDetailsForUser{}, err
		}
		orderDetails.DeliveryPartnerName = deliveryPartner.Name
		orderDetails.DeliveryPartnerPhone = deliveryPartner.Phone
	}

	if order.PriceAfterOffer != 0 {
		orderDetails.Price = order.PriceAfterOffer
	} else {
		orderDetails.Price = order.TotalPrice
	}

	orderDetails.OrderStatus = order.OrderStatus

	return orderDetails, nil
}

func GetActiveOrdersForUser(ctx context.Context, userId primitive.ObjectID) ([]models.Order, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderUserID: userId,
		models.KeyOrderStatus: bson.M{
			constants.MongoKeywordIn: []string{"Pending", constants.KeyOrderAccepted, constants.KeyOrderOutForDelivery},
		},
	}
	cursor, err := orderCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetRestaurantByOrderID(ctx context.Context, orderId primitive.ObjectID) (models.Restaurant, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID: orderId,
	}
	var order models.Order
	err := orderCollection.FindOne(ctx, query).Decode(&order)
	if err != nil {
		return models.Restaurant{}, err
	}

	restaurant, err := GetRestaurantByID(ctx, order.RestaurantID)
	if err != nil {
		return models.Restaurant{}, err
	}

	return restaurant, nil
}

func GetAllPendingOrdersByRestaurantID(ctx context.Context, restaurantId primitive.ObjectID, skip int64, limit int64) ([]models.Order, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderRestaurantID: restaurantId,
		models.KeyOrderStatus:       "Pending",
	}
	cursor, err := orderCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetAllAcceptedOrdersByRestaurantID(ctx context.Context, restaurantId primitive.ObjectID, skip int64, limit int64) ([]models.Order, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderRestaurantID: restaurantId,
		models.KeyOrderStatus:       constants.KeyOrderAccepted,
	}
	cursor, err := orderCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetAllOrdersByRestaurantID(ctx context.Context, restaurantId primitive.ObjectID, skip int64, limit int64) ([]models.Order, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderRestaurantID: restaurantId,
	}
	cursor, err := orderCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			models.KeyOrderID: -1,
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderDetailsByRestaurantID(ctx context.Context, orderId primitive.ObjectID, restaurantId primitive.ObjectID) (models.OrderDetailsForRestaurant, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:           orderId,
		models.KeyOrderRestaurantID: restaurantId,
	}
	var order models.Order
	err := orderCollection.FindOne(ctx, query).Decode(&order)
	if err != nil {
		return models.OrderDetailsForRestaurant{}, err
	}

	var orderForRestaurant models.OrderDetailsForRestaurant
	orderForRestaurant.ID = order.ID
	orderForRestaurant.UserID = order.UserID
	orderForRestaurant.Items = order.Items
	orderForRestaurant.Notes = order.Notes
	if order.PriceAfterOffer != 0 {
		orderForRestaurant.Price = order.PriceAfterOffer
	} else {
		orderForRestaurant.Price = order.TotalPrice
	}

	orderForRestaurant.OrderStatus = order.OrderStatus

	if order.DeliveryPartnerID != primitive.NilObjectID {
		deliveryPartner, err := GetDeliveryPartnerByID(ctx, order.DeliveryPartnerID)
		if err != nil {
			return models.OrderDetailsForRestaurant{}, err
		}
		orderForRestaurant.DeliveryPartnerName = deliveryPartner.Name
		orderForRestaurant.DeliveryPartnerPhone = deliveryPartner.Phone
	}

	return orderForRestaurant, nil
}

func AcceptOrderByRestaurantID(ctx context.Context, orderId primitive.ObjectID, restaurantId primitive.ObjectID) error {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:           orderId,
		models.KeyOrderRestaurantID: restaurantId,
		models.KeyOrderStatus:       "Pending",
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyOrderStatus: constants.KeyOrderAccepted,
		},
	}
	_, err := orderCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil
}

func RejectOrderByRestaurantID(ctx context.Context, orderId primitive.ObjectID, restaurantId primitive.ObjectID) error {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:           orderId,
		models.KeyOrderRestaurantID: restaurantId,
		models.KeyOrderStatus:       "Pending",
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyOrderStatus: constants.KeyOrderRejected,
		},
	}
	_, err := orderCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil
}

func GetIncomingOrdersForDelivery(ctx context.Context, deliveryPartneId primitive.ObjectID, skip int64, limit int64) ([]models.Order, error) {
	deliveryPartnerCollection := models.GetDeliveryPartnerCollection()
	query := bson.M{
		models.KeyDeliveryPartnerID: deliveryPartneId,
	}
	var deliveryPartner models.DeliveryPartner
	err := deliveryPartnerCollection.FindOne(ctx, query).Decode(&deliveryPartner)
	if err != nil {
		return nil, err
	}

	orderCollection := models.GetOrderCollection()
	query = bson.M{
		models.KeyOrderStatus: constants.KeyOrderAccepted,
		helpers.MongoJoinFields(models.KeyOrderDeliveryAddress, models.KeyAddressCity): deliveryPartner.City,
	}

	cursor, err := orderCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			models.KeyOrderID: -1,
		},
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func AddDeliveryPartnerToOrder(ctx context.Context, orderId primitive.ObjectID, deliveryPartnerId primitive.ObjectID) error {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID: orderId,
	}
	order, err := GetOrderByID(ctx, orderId)
	if err != nil {
		return err
	}

	if order.DeliveryPartnerID != primitive.NilObjectID {
		return errors.New("delivery partner already assigned to order")
	}

	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyOrderDeliveryPartnerID: deliveryPartnerId,
		},
	}

	err = orderCollection.FindOneAndUpdate(ctx, query, update).Err()
	return err
}

func GetOngoingOrdersForDelivery(ctx context.Context, deliveryPartnerId primitive.ObjectID, skip int64, limit int64) ([]models.Order, error) {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderStatus: bson.M{
			constants.MongoKeywordIn: []string{constants.KeyOrderAccepted, constants.KeyOrderOutForDelivery},
		},
		models.KeyOrderDeliveryPartnerID: deliveryPartnerId,
	}
	cursor, err := orderCollection.Find(ctx, query, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			models.KeyOrderID: -1,
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err = cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func CompleteOrderForDelivery(ctx context.Context, orderId primitive.ObjectID, deliveryPartnerId primitive.ObjectID) error {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:                orderId,
		models.KeyOrderDeliveryPartnerID: deliveryPartnerId,
		models.KeyOrderStatus:            "OutForDelivery",
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyOrderStatus: constants.KeyOrderDelivered,
		},
	}
	_, err := orderCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil
}

func PickupOrderForDelivery(ctx context.Context, orderId primitive.ObjectID, deliveryPartnerId primitive.ObjectID) error {
	orderCollection := models.GetOrderCollection()
	query := bson.M{
		models.KeyOrderID:                orderId,
		models.KeyOrderDeliveryPartnerID: deliveryPartnerId,
		models.KeyOrderStatus:            constants.KeyOrderAccepted,
	}
	update := bson.M{
		constants.MongoKeywordSet: bson.M{
			models.KeyOrderStatus: constants.KeyOrderOutForDelivery,
		},
	}
	_, err := orderCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	return nil
}

func GetOrderDetailsForDeliveryPartner(ctx context.Context, orderId primitive.ObjectID, deliveryPartnerId primitive.ObjectID) (models.OrderDetailsForDeliveryPartner, error) {
	order, err := GetOrderByID(ctx, orderId)
	if err != nil {
		return models.OrderDetailsForDeliveryPartner{}, err
	}

	if order.DeliveryPartnerID != deliveryPartnerId {
		return models.OrderDetailsForDeliveryPartner{}, errors.New("delivery partner not assigned to order")
	}

	orderDetailsForDeliveryPartner := models.OrderDetailsForDeliveryPartner{}
	orderDetailsForDeliveryPartner.ID = order.ID
	orderDetailsForDeliveryPartner.OrderStatus = order.OrderStatus
	orderDetailsForDeliveryPartner.DeliveryAddress = order.DeliveryAddress
	orderDetailsForDeliveryPartner.Items = order.Items

	user, err := GetUserByID(ctx, order.UserID)
	if err != nil {
		return models.OrderDetailsForDeliveryPartner{}, err
	}

	orderDetailsForDeliveryPartner.Username = user.Name
	orderDetailsForDeliveryPartner.UserPhone = user.Phone

	if order.PaymentMethod == "COD" {
		orderDetailsForDeliveryPartner.PaymentMethod = "Cash On Delivery"
		if order.PriceAfterOffer > 0 {
			orderDetailsForDeliveryPartner.Price = order.PriceAfterOffer
		} else {
			orderDetailsForDeliveryPartner.Price = order.TotalPrice
		}
	} else {
		orderDetailsForDeliveryPartner.PaymentMethod = "Online Payment"
	}

	restaurant, err := GetRestaurantByID(ctx, order.RestaurantID)
	if err != nil {
		return models.OrderDetailsForDeliveryPartner{}, err
	}

	orderDetailsForDeliveryPartner.RestaurantName = restaurant.Name
	orderDetailsForDeliveryPartner.RestaurantPhone = restaurant.Phone
	orderDetailsForDeliveryPartner.RestaurantAddress = restaurant.Address

	return orderDetailsForDeliveryPartner, nil
}
