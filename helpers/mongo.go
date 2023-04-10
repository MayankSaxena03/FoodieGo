package helpers

import (
	"strings"

	"github.com/MayankSaxena03/FoodieGo/constants"
)

func MongoJoinFields(fields ...string) string {
	return strings.Join(fields, constants.MongoKeySeparator)
}
