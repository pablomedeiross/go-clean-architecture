package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoSQLDB interface {
	SaveUser(ctx context.Context, user User) (primitive.ObjectID, error)
	FindUserByName(ctx context.Context, name string) (User, error)
	DeleteUser(ctx context.Context, nam string) error
}
