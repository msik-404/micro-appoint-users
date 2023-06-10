package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/msik-404/micro-appoint-users/internal/database"
)

type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Mail      *string            `bson:"mail,omitempty" binding:"max=30"`
	HashedPwd *string            `bson:"pwd,omitempty"`
	Name      *string            `bson:"name,omitempty" binding:"max=30"`
	Surname   *string            `bson:"surname,omitempty" binding:"max=30"`
}

func FindOneCustomerCredentials(
	ctx context.Context,
	db *mongo.Database,
	mail string,
) *mongo.SingleResult {
	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{Key: "_id", Value: 1},
		{Key: "pwd", Value: 1},
	})

	coll := db.Collection(database.CustomersCollName)
	filter := bson.M{"mail": mail}
	return coll.FindOne(ctx, filter, opts)
}

func (customer *Customer) InsertOne(
	ctx context.Context,
	db *mongo.Database,
) (*mongo.InsertOneResult, error) {
	coll := db.Collection(database.CustomersCollName)
	return coll.InsertOne(ctx, customer)
}

func DeleteOneCustomer(
	ctx context.Context,
	db *mongo.Database,
	customerId primitive.ObjectID,
) (*mongo.DeleteResult, error) {
	coll := db.Collection(database.CustomersCollName)
	filter := bson.M{"_id": customerId}
	return coll.DeleteOne(ctx, filter)
}

func (customerUpdate *Customer) UpdateOne(
	ctx context.Context,
	db *mongo.Database,
	customerID primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	coll := db.Collection(database.CustomersCollName)
	update := bson.M{"$set": customerUpdate}
	return coll.UpdateByID(ctx, customerID, update)
}

type Owner struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Mail      *string              `bson:"mail,omitempty" binding:"max=30"`
	HashedPwd *string              `bson:"pwd,omitempty"`
	Name      *string              `bson:"name,omitempty" binding:"max=30"`
	Surname   *string              `bson:"surname,omitempty" binding:"max=30"`
	Companies []primitive.ObjectID `bson:"companies,omitempty"`
}

func FindOneOwnerCredentials(
	ctx context.Context,
	db *mongo.Database,
	mail string,
) *mongo.SingleResult {
	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{Key: "_id", Value: 1},
		{Key: "pwd", Value: 1},
	})

	coll := db.Collection(database.OwnersCollName)
	filter := bson.M{"mail": mail}
	return coll.FindOne(ctx, filter, opts)
}

func FindOneOwnerCompanies(
	ctx context.Context,
	db *mongo.Database,
	ownerID primitive.ObjectID,
) *mongo.SingleResult {
	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{Key: "companies", Value: 1},
	})

	coll := db.Collection(database.OwnersCollName)
	filter := bson.M{"_id": ownerID}
	return coll.FindOne(ctx, filter, opts)
}

func (owner *Owner) InsertOne(
	ctx context.Context,
	db *mongo.Database,
) (*mongo.InsertOneResult, error) {
	coll := db.Collection(database.OwnersCollName)
	return coll.InsertOne(ctx, owner)
}

func InsertOneOwnerCompany(
	ctx context.Context,
	db *mongo.Database,
	ownerID primitive.ObjectID,
	companyID primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	coll := db.Collection(database.OwnersCollName)
	update := bson.M{"$push": bson.M{"companies": companyID}}
	return coll.UpdateByID(ctx, ownerID, update)
}

func DeleteOneOwnerCompany(
	ctx context.Context,
	db *mongo.Database,
	ownerID primitive.ObjectID,
	companyID primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	coll := db.Collection(database.OwnersCollName)
	update := bson.M{"$pull": bson.M{"companies": companyID}}
	return coll.UpdateByID(ctx, ownerID, update)
}

func DeleteOneOwner(
	ctx context.Context,
	db *mongo.Database,
	ownerID primitive.ObjectID,
) (*mongo.DeleteResult, error) {
	coll := db.Collection(database.OwnersCollName)
	filter := bson.M{"_id": ownerID}
	return coll.DeleteOne(ctx, filter)
}

func (ownerUpdate *Owner) UpdateOne(
	ctx context.Context,
	db *mongo.Database,
	ownerId primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	coll := db.Collection(database.OwnersCollName)
	update := bson.M{"$set": ownerUpdate}
	return coll.UpdateByID(ctx, ownerId, update)
}
