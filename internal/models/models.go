package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Customer struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Mail      string               `bson:"mail" binding:"max=30"`
	HashedPwd []byte               `bson:"pwd"`
	Name      string               `bson:"name,omitempty" binding:"max=30"`
	Surname   string               `bson:"surname,omitempty" binding:"max=30"`
	Orders    []primitive.ObjectID `bson:"orders,omitempty"`
}

func FindOneCustomer(
	db *mongo.Database,
	mail string,
) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{Key: "_id", Value: 1},
		{Key: "pwd", Value: 1},
	})

	coll := db.Collection("customers")
	filter := bson.M{"mail": mail}
	return coll.FindOne(ctx, filter, opts)
}

func (customer *Customer) InsertOne(
	db *mongo.Database,
) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("customers")
	return coll.InsertOne(ctx, customer)
}

func DeleteOneCustomer(
	db *mongo.Database,
	customerId primitive.ObjectID,
) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("customers")
	filter := bson.M{"_id": customerId}
	return coll.DeleteOne(ctx, filter)
}

type CustomerUpdate struct {
	Mail      string               `bson:"mail" binding:"max=30"`
	HashedPwd []byte               `bson:"pwd"`
	Name      string               `bson:"name,omitempty" binding:"max=30"`
	Surname   string               `bson:"surname,omitempty" binding:"max=30"`
	Orders    []primitive.ObjectID `bson:"orders,omitempty"`
}

func (customerUpdate *CustomerUpdate) UpdateOne(
	db *mongo.Database,
	customerID primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("customers")
	update := bson.M{"$set": customerUpdate}
	return coll.UpdateByID(ctx, customerID, update)
}

type Owner struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Mail        string               `bson:"mail" binding:"max=30"`
	HashedPwd   []byte               `bson:"pwd"`
	Name        string               `bson:"name,omitempty" binding:"max=30"`
	Surname     string               `bson:"surname,omitempty" binding:"max=30"`
	Possessions []primitive.ObjectID `bson:"possessions,omitempty"`
}

func FindOneOwner(
	db *mongo.Database,
	mail string,
) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{Key: "_id", Value: 1},
		{Key: "pwd", Value: 1},
	})

	coll := db.Collection("owners")
	filter := bson.M{"mail": mail}
	return coll.FindOne(ctx, filter, opts)
}

func (owner *Owner) InsertOne(
	db *mongo.Database,
) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("owners")
	return coll.InsertOne(ctx, owner)
}

func DeleteOneOwner(
	db *mongo.Database,
	ownerID primitive.ObjectID,
) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("owners")
	filter := bson.M{"_id": ownerID}
	return coll.DeleteOne(ctx, filter)
}

type OwnerUpdate struct {
	Mail        string               `bson:"mail" binding:"max=30"`
	HashedPwd   []byte               `bson:"pwd"`
	Name        string               `bson:"name,omitempty" binding:"max=30"`
	Surname     string               `bson:"surname,omitempty" binding:"max=30"`
	Possessions []primitive.ObjectID `bson:"possessions,omitempty"`
}

func (ownerUpdate *OwnerUpdate) UpdateOne(
	db *mongo.Database,
	ownerId primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("owners")
	update := bson.M{"$set": ownerUpdate}
	return coll.UpdateByID(ctx, ownerId, update)
}
