package double

import (
	"context"
	"user-api/adapter/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoSQLDBDouble struct {
	SaveUserFunc       func(ctx context.Context, user db.User) (primitive.ObjectID, error)
	FindUserByNameFunc func(ctx context.Context, name string) (db.User, error)
	DeleteUserFunc     func(ctx context.Context, nam string) error
}

func NewNoSQLDB(

	saveUser func(ctx context.Context, user db.User) (primitive.ObjectID, error),
	findUserByName func(ctx context.Context, name string) (db.User, error),
	deleteUser func(ctx context.Context, nam string) error,

) db.NoSQLDB {

	return &NoSQLDBDouble{
		SaveUserFunc:       saveUser,
		FindUserByNameFunc: findUserByName,
		DeleteUserFunc:     deleteUser,
	}
}

func (gateway *NoSQLDBDouble) SaveUser(ctx context.Context, user db.User) (primitive.ObjectID, error) {
	return gateway.SaveUserFunc(ctx, user)
}

func (gateway *NoSQLDBDouble) FindUserByName(ctx context.Context, name string) (db.User, error) {
	return gateway.FindUserByNameFunc(ctx, name)
}

func (gateway *NoSQLDBDouble) DeleteUser(ctx context.Context, name string) error {
	return gateway.DeleteUserFunc(ctx, name)
}
