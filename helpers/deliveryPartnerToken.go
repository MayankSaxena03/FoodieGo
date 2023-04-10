package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeliveryPartnerSignedDetails struct {
	PartnerID    primitive.ObjectID
	PartnerPhone string
	jwt.StandardClaims
}

func GenerateAllTokensForDeliveryPartner(partnerID primitive.ObjectID, phone string) (string, string, error) {
	var SECRETKEY = os.Getenv("SECRETKEY")
	claims := &DeliveryPartnerSignedDetails{
		PartnerID:    partnerID,
		PartnerPhone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &DeliveryPartnerSignedDetails{
		PartnerID:    partnerID,
		PartnerPhone: phone,
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

func ValidateDeliveryPartnerToken(signedToken string) (*DeliveryPartnerSignedDetails, error) {
	var SECRETKEY = os.Getenv("SECRETKEY")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&DeliveryPartnerSignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRETKEY), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*DeliveryPartnerSignedDetails)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
