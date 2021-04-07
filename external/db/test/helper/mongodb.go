package helper

import (
	"context"
	"user-api/adapter/db"
	"user-api/external/db/memory"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const user_collection_name = "User"

// Insert a user in MongoDB
func InsertUser(helper *memory.InMemoryMongoDB, usr db.User) error {

	collection, err := helper.ConnectInCollection(user_collection_name)
	_, err = collection.InsertOne(context.Background(), &usr, options.InsertOne())

	return err
}

// Search user by name in MongoDB, if don't find return a error
func FindUserById(helper *memory.InMemoryMongoDB, id primitive.ObjectID) (db.User, error) {

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
