package core

import (
	"context"
	"errors"
	"reflect"
	"time"

	models "github.com/MayankSaxena03/FoodieGo/models/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func findDifferentFields(s1 interface{}, s2 interface{}) ([]string, error) {
	differentFields := []string{}

	s1Value := reflect.ValueOf(s1)
	s2Value := reflect.ValueOf(s2)

	if s1Value.Kind() != reflect.Struct || s2Value.Kind() != reflect.Struct {
		return nil, errors.New("v1 and v2 must be structs")
	}

	if s1Value.Type() != s2Value.Type() {
		return nil, errors.New("v1 and v2 must be of the same type")
	}

	for i := 0; i < s1Value.NumField(); i++ {
		s1Field := s1Value.Field(i)
		s2Field := s2Value.Field(i)

		if !reflect.DeepEqual(s1Field.Interface(), s2Field.Interface()) {
			differentFields = append(differentFields, s1Value.Type().Field(i).Name)
		}
	}

	return differentFields, nil
}

func CreateChangeLog(ctx context.Context, v1 interface{}, v2 interface{}, entityID primitive.ObjectID, entityType string) error {
	logCollection := models.GetLogCollection()

	diff, err := findDifferentFields(v1, v2)
	if err != nil {
		return err
	}
	if len(diff) == 0 {
		return nil
	}

	changes := make([]models.Changes, len(diff))
	for i, fieldName := range diff {
		changes[i] = models.Changes{
			FieldName: fieldName,
			OldValue:  reflect.ValueOf(v1).FieldByName(fieldName).Interface(),
			NewValue:  reflect.ValueOf(v2).FieldByName(fieldName).Interface(),
		}
	}

	log := models.Log{
		EntityID:   entityID,
		EntityType: entityType,
		CreatedOn:  time.Now(),
		Changes:    changes,
	}

	_, err = logCollection.InsertOne(ctx, log)

	return err
}
