package core

import (
	"context"
	"time"

	redisModels "github.com/MayankSaxena03/FoodieGo/models/redis"

	"github.com/MayankSaxena03/FoodieGo/database"
	"github.com/MayankSaxena03/FoodieGo/helpers"
)

func SetLoginOTPForPhoneInRedis(phone string, otp string) error {
	key := helpers.ParseRedisKey(redisModels.KeyPhoneLoginOTP, []string{phone})
	err := database.RedisClient.Set(context.Background(), key, otp, 5*time.Minute).Err()
	return err
}

func GetLoginOTPForPhoneFromRedis(phone string) (string, error) {
	key := helpers.ParseRedisKey(redisModels.KeyPhoneLoginOTP, []string{phone})
	otp, err := database.RedisClient.Get(context.Background(), key).Result()
	return otp, err
}

func DeleteLoginOTPForPhoneInRedis(phone string) error {
	key := helpers.ParseRedisKey(redisModels.KeyPhoneLoginOTP, []string{phone})
	err := database.RedisClient.Del(context.Background(), key).Err()
	return err
}

func SetLoginOTPForEmailInRedis(email string, otp string) error {
	key := helpers.ParseRedisKey(redisModels.KeyEmailLoginOTP, []string{email})
	err := database.RedisClient.Set(context.Background(), key, otp, 5*time.Minute).Err()
	return err
}

func GetLoginOTPForEmailFromRedis(email string) (string, error) {
	key := helpers.ParseRedisKey(redisModels.KeyEmailLoginOTP, []string{email})
	otp, err := database.RedisClient.Get(context.Background(), key).Result()
	return otp, err
}

func DeleteLoginOTPForEmailInRedis(email string) error {
	key := helpers.ParseRedisKey(redisModels.KeyEmailLoginOTP, []string{email})
	err := database.RedisClient.Del(context.Background(), key).Err()
	return err
}
