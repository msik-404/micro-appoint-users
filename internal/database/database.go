package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBName = os.Getenv("DB_NAME")

const (
	CustomersCollName  = "customers"
	OwnersCollName = "owners"
)

func getURI() string {
	return fmt.Sprintf("mongodb://%s:%s@mongodb:27017", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
}

func ConnectDB() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(getURI()).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a new client and connect to the server
	return mongo.Connect(ctx, opts)
}

func CreateDBIndexes(client *mongo.Client) ([][]string, []error) {
	db := client.Database(DBName)

	usersColl := db.Collection(CustomersCollName)
	usersIndex := []mongo.IndexModel{
		{
			Keys:    bson.M{"mail": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	ownersColl := db.Collection(OwnersCollName)
	ownersIndex := []mongo.IndexModel{
		{
			Keys:    bson.M{"mail": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	var results [][]string
	var errors []error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := usersColl.Indexes().CreateMany(ctx, usersIndex)
	results = append(results, result)
	errors = append(errors, err)

	result, err = ownersColl.Indexes().CreateMany(ctx, ownersIndex)
	results = append(results, result)
	errors = append(errors, err)

	return results, errors
}
