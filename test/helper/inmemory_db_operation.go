package helper

import (
	"context"
	"user-api/adapter/db"
	"user-api/db/inmemory"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const user_collection_name = "User"

// Insert a user in MongoDB
func InsertUser(helper *inmemory.InMemoryMongoDB, usr db.User) error {

	collection, _ := helper.ConnectInCollection(user_collection_name)
	_, err := collection.InsertOne(context.Background(), &usr, options.InsertOne())

	return err
}

// Search user by name in MongoDB, if don't find return a error
func FindUserById(helper *inmemory.InMemoryMongoDB, id primitive.ObjectID) (db.User, error) {

	var userReturned *db.User = &db.User{}
	collection, err := helper.ConnectInCollection(user_collection_name)

	if err != nil {
		return *userReturned, nil
	}

	err = collection.
		FindOne(context.Background(), &bson.M{"_id": id}, options.FindOne()).
		Decode(userReturned)

	return *userReturned, err
}

func FindUserByName(helper *inmemory.InMemoryMongoDB, name string) (db.User, error) {

	var userReturned *db.User = &db.User{}
	collection, err := helper.ConnectInCollection(user_collection_name)

	if err != nil {
		return *userReturned, nil
	}

	err = collection.
		FindOne(context.Background(), &bson.M{"name": name}, options.FindOne()).
		Decode(userReturned)

	return *userReturned, err
}
