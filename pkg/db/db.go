package db

import (
	"context"
	"log"
	"time"

	"github.com/kernelsafe/weather-api-go/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DBName holds the database name
var DBName string

const (
	dbHost = "DB_HOST"
	dbName = "DB_NAME"
	// TimeOut of the connection
	TimeOut = 5 * time.Second
)

func check(e error) {
	if e != nil {
		// panic(e)
		log.Fatal(e)
	}
}

// GetClient returns a Client object
func GetClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(util.GetEnv(dbHost, "mongodb://localhost:27017")))
	check(err)
	DBName = util.GetEnv(dbName, "weathier-api")
	err = client.Ping(ctx, readpref.Primary())
	check(err)
	return client
}

// AddDocument adds a document to the database
func AddDocument(client *mongo.Client, collectionName string, doc interface{}) (interface{}, error) {
	collection := client.Database(DBName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}
