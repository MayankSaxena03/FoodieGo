package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantSignedDetails struct {
	RestaurantID    primitive.ObjectID
	RestaurantEmail string
	jwt.StandardClaims
}

func GenerateAllTokensForRestaurant(restaurantID primitive.ObjectID, email string) (string, string, error) {
	var SECRETKEY = os.Getenv("SECRETKEY")
	claims := &RestaurantSignedDetails{
		RestaurantID:    restaurantID,
		RestaurantEmail: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &RestaurantSignedDetails{
		RestaurantID:    restaurantID,
		RestaurantEmail: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateRestaurantToken(signedToken string) (*RestaurantSignedDetails, error) {
	var SECRETKEY = os.Getenv("SECRETKEY")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&RestaurantSignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRETKEY), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RestaurantSignedDetails)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
