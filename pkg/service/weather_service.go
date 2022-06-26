package service

import (
	"context"

	"github.com/kernelsafe/weather-api-go/pkg/db"
	"github.com/kernelsafe/weather-api-go/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func SetCollection(newCollection *mongo.Collection) {
	collection = newCollection
}
// GetOne returns one item
func GetOne(id primitive.ObjectID) (*model.WeatherResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.TimeOut)
	defer cancel()
	res := &model.WeatherResponse{}
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteOne deletes one item
func DeleteOne(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.TimeOut)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// InsertOne inserts one item
func InsertOne(doc model.WeatherRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.TimeOut)
	defer cancel()
	_, err := collection.InsertOne(ctx, doc)
	return err
}

// GetAll returns all items
func GetAll() ([]*model.WeatherResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.TimeOut)
	defer cancel()
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"date": -1})
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var temperatures []*model.WeatherResponse

	for cursor.Next(context.TODO()) {
		var value model.WeatherResponse
		err := cursor.Decode(&value)
		if err != nil {
			return nil, err
		}
		temperatures = append(temperatures, &value)
	}

	return temperatures, nil
}
