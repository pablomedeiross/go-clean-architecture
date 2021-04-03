package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Age          int                `bson:"age"`
	AddressesIds []string           `bson:"addressesIds,omitempty"`
}
