# **FoodieGo Backend**

## Description
This is a food delivery  backend built using Golang. The api's allows users to browse restaurants, view menus, place orders, and track deliveries. The api's also supports delivery partners who can accept and deliver orders and restaurants who can manage their menus and orders.

## Directory Structure

```
├── /constants
│   ├── common.go
│   ├── entities.go
│   ├── mongoConstant.go
│   └── order.go
├── /core
│   ├── address.go
│   ├── cart.go
│   ├── deliveryPartner.go
│   ├── food.go
│   ├── logs.go
│   ├── menu.go
│   ├── offers.go
│   ├── orders.go
│   ├── otp.go
│   ├── ratings.go
│   ├── restaurant.go
│   └── user.go
├── /database
│   ├── mongoConnection.go
│   └── redisConnection.go
├── /handlers
│   ├── /authentication
│   │   ├── deliveryPartnerAuthentication.go
│   │   ├── restaurantAuthentication.go
│   │   └── userAuthentication.go
│   └── /common
│       └── otp.go
├── /helpers
│   ├── deliveryPartner.go
│   ├── location.go
│   ├── mongo.go
│   ├── otp.go
│   ├── password.go
│   ├── redis.go
│   ├── requestDetails.go
│   ├── restaurantToken.go
│   └── userToken.go
├── /middleware
│   └── authenticate.go
├── /models
│   ├── /mongo
│   │   ├── address.go
│   │   ├── cart.go
│   │   ├── deliveryPartner.go
│   │   ├── food.go
│   │   ├── logs.go
│   │   ├── menu.go
│   │   ├── offers.go
│   │   ├── orders.go
│   │   ├── ratings.go
│   │   ├── restaurant.go
│   │   └── user.go
│   └── /redis
│       └── otp.go
├── /routes
│   ├── deliveryPartner.go
│   ├── restaurant.go
│   └── user.go
├── .env
├── go.mod
├── go.sum
└── server.go
```

In this directory structure, we have separate directories for constants, core functionality, database connections, handlers, helpers, middleware, models, and routes. We also have files for managing dependencies and environment variables.

The constants directory contains various constants that are used throughout the project, such as error codes and API endpoints. The core directory contains the core functionality of the application, including logic for handling orders, menus, restaurants, and users. The database directory contains code for connecting to MongoDB and Redis databases. The handlers directory contains handlers for the different API routes, separated into authentication and common handlers. The helpers directory contains utility functions used across different parts of the code. The middleware directory contains middleware functions for authenticating users and performing other tasks before handling API requests. The models directory contains data models for MongoDB and Redis. The routes directory contains the API routes for delivery partners, restaurants, and users. Finally, we have files for managing dependencies and environment variables.

## Dependencies

This project requires the following dependencies:

- github.com/dgrijalva/jwt-go v3.2.0+incompatible
- github.com/go-playground/validator/v10 v10.12.0
- github.com/gorilla/handlers v1.5.1
- github.com/gorilla/mux v1.8.0
- github.com/joho/godotenv v1.5.1
- github.com/redis/go-redis/v9 v9.0.3
- go.mongodb.org/mongo-driver v1.11.4
- golang.org/x/crypto v0.7.0

Make sure to install these dependencies before running the project. You can install them using the `go get` command or by adding them to your `go.mod` file and running `go mod download`.

## Installation and Usage

To install and use this project, follow these steps:

1. Install Go on your system if it is not already installed. You can download Go from the official website: https://golang.org/dl/

2. Clone the project repository using the command:

    ```bash
    git clone https://github.com/MayankSaxena03/FoodieGo.git
    ```

3. Navigate to the project directory:

    ```bash
    cd FoodieGo
    ```

4. Install the required dependencies by running the following command:

    ```go
    go mod download
    ```

5. Create a file named .env in the project root directory and add the required environment variables:

    ```makefile
    MONGODB_URL=mongodb://localhost:27017/food_delivery
    REDIS_URL=localhost:6379
    SECRETKEY=my_secret_key
    ```

    Replace my_secret_key with your own secret key for JWT token encryption.

6. Start the server by running the following command:

    ```go
    go run server.go
    ```

    This will start the server on the default port (8080).

7. You can now access the API endpoints using any HTTP client, such as Postman or cURL.

Note: Make sure that MongoDB and Redis are running on your system before starting the server.


# API Endpoints for Users

## Authentication Routes
| Route                        | Method | Description                          |
|------------------------------|--------|--------------------------------------|
| /api/user/auth/login         | POST   | Log in a user                        |
| /api/user/auth/generate-otp  | POST   | Generate OTP for email verification |

## Profile Routes
| Route                                            | Method | Description                            |
|--------------------------------------------------|--------|----------------------------------------|
| /api/user/validate                               | GET    | Validate user profile                  |
| /api/user/me                                     | GET    | Get user info                          |
| /api/user/update                                 | PUT    | Update user info                       |
| /api/user/address                                | POST   | Add user address                       |
| /api/user/primary-address/{addressId}            | PUT    | Set primary user address               |
| /api/user/address/{addressId}                     | PUT    | Update user address by ID              |
| /api/user/address/{addressId}                     | DELETE | Delete user address by ID              |
| /api/user/address                                | GET    | Get all user addresses                 |
| /api/user/address/{addressId}                     | GET    | Get user address by ID                 |

## RESTAURANT ROUTES
| Route                                            | Method | Description                            |
|--------------------------------------------------|--------|----------------------------------------|
| /api/user/restaurants                            | GET    | Get all restaurants                    |
| /api/user/restaurants/{restaurantId}             | GET    | Get restaurant by ID                   |
| /api/user/restaurants/{restaurantId}/menu        | GET    | Get restaurant menu                    |
| /api/user/restaurants/offers/{restaurantId}      | GET    | Get restaurant offers                  |

## Cart Routes
| Route                                            | Method | Description                            |
|--------------------------------------------------|--------|----------------------------------------|
| /api/user/cart                                   | GET    | Get user cart                          |
| /api/user/cart                                   | POST   | Add item to user cart                  |
| /api/user/cart/{foodId}                          | DELETE | Remove item from user cart             |
| /api/user/clear-cart                             | DELETE | Clear user cart                        |
| /api/user/cart/notes                             | POST   | Add notes to user cart                 |
| /api/user/apply-coupon/{couponCode}              | POST   | Apply coupon code to user cart         |
| /api/user/remove-coupon                          | DELETE | Remove coupon from user cart           |

## Order Routes
| Route                                            | Method | Description                            |
|--------------------------------------------------|--------|----------------------------------------|
| /api/user/create-order                           | POST   | Create new order for user              |
| /api/user/cancel-order/{orderId}                 | POST   | Cancel order for user                  |
| /api/user/orders                                 | GET    | Get all orders for user                |
| /api/user/orders/{orderId}                       | GET    | Get user order by ID                   |
| /api/user/active-orders                          | GET    | Get active orders for user             |
| /api/user/rating/{orderID}                       | POST   | Add rating for a specific order for user|


# API Endpoints for Restaurants

## Authentication Routes

| Route                                | Method | Description                            |
|--------------------------------------|--------|----------------------------------------|
| /api/restaurant/auth/signup          | POST   | Sign up a new restaurant               |
| /api/restaurant/auth/login           | POST   | Log in a restaurant                    |
| /api/restaurant/auth/generate-otp    | POST   | Generate OTP for email verification    |

## Profile Routes

| Route                            | Method | Description                    |
|----------------------------------|--------|--------------------------------|
| /api/restaurant/validate        | GET    | Validate restaurant profile    |
| /api/restaurant/info            | GET    | Get restaurant info             |
| /api/restaurant/update          | PUT    | Update restaurant info          |

## Menu Routes

| Route                                       | Method | Description                 |
|---------------------------------------------|--------|-----------------------------|
| /api/restaurant/create-menu                 | POST   | Create a new menu            |
| /api/restaurant/get-menu                    | GET    | Get all menus for a restaurant |
| /api/restaurant/get-menu-items              | GET    | Get all menu items for a restaurant |
| /api/restaurant/add-menu-item               | POST   | Add a new menu item          |
| /api/restaurant/delete-menu-item/{menuItemId}| DELETE | Delete a menu item           |
| /api/restaurant/update-menu-item/{menuItemId}| PUT    | Update a menu item           |

## Offer Routes

| Route                                     | Method | Description                        |
|-------------------------------------------|--------|------------------------------------|
| /api/restaurant/create-offer              | POST   | Create a new offer                  |
| /api/restaurant/get-offers                | GET    | Get all offers for a restaurant     |
| /api/restaurant/get-offer/{offerId}       | GET    | Get a specific offer for a restaurant |
| /api/restaurant/update-offer/{offerId}    | PUT    | Update a specific offer for a restaurant |
| /api/restaurant/delete-offer/{offerId}    | DELETE | Delete a specific offer for a restaurant |

## Order Routes

| Route                                     | Method | Description                            |
|-------------------------------------------|--------|----------------------------------------|
| /api/restaurant/pending-orders            | GET    | Get all pending orders for a restaurant |
| /api/restaurant/accepted-orders           | GET    | Get all accepted orders for a restaurant |
| /api/restaurant/orders                    | GET    | Get all orders for a restaurant         |
| /api/restaurant/order/{orderId}           | GET    | Get a specific order for a restaurant   |
| /api/restaurant/accept-order/{orderId}    | PUT    | Accept a specific order for a restaurant |
| /api/restaurant/reject-order/{orderId}    | PUT    | Reject a specific order for a restaurant |

## Income Routes

| Route                          | Method | Description                   |
|--------------------------------|--------|-------------------------------|
| /api/restaurant/income        | GET    | Get the income of a restaurant |


# API Endpoints for Delivery Partners

## Authentication Routes

| Endpoint                      | Method | Description                               |
| -----------------------------|--------| ----------------------------------------- |
| /api/delivery-partner/auth/login | POST   | Delivery partner login                    |
| /api/delivery-partner/auth/generate-otp | POST   | Generate OTP for delivery partner email   |

## Delivery Partner Routes

| Endpoint                                     | Method | Description                                          |
| -------------------------------------------- |--------| ---------------------------------------------------- |
| /api/delivery-partner/validate               | GET    | Validate delivery partner                            |
| /api/delivery-partner/me                      | GET    | Get delivery partner info                            |
| /api/delivery-partner/update                  | PUT    | Update delivery partner info                         |
| /api/delivery-partner/incoming-orders         | GET    | Get incoming orders for delivery partner             |
| /api/delivery-partner/accept-order/{orderID}  | POST   | Accept order for delivery                            |
| /api/delivery-partner/ongoing-orders          | GET    | Get ongoing orders for delivery partner              |
| /api/delivery-partner/pickup-order/{orderID}  | POST   | Pickup order for delivery                            |
| /api/delivery-partner/complete-order/{orderID}| POST   | Complete order for delivery                          |
| /api/delivery-partner/order/{orderID}         | GET    | Get order details for delivery partner               |
