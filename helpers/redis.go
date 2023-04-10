package helpers

import "strings"

func ParseRedisKey(key string, values []string) string {
	for _, value := range values {
		key = strings.Replace(key, "?", value, 1)
	}
	return key
}
