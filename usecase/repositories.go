package usecase

import (
	"context"
	"errors"
	"user-api/entity/user"
)

// UserRepository is an interface that represents the data persistent for user entities
type UserRepository interface {
	FindByName(ctx context.Context, name string) (user.User, error)
	Save(ctx context.Context, user user.User) (user.User, error)
}

type UserDontExist error

// NewUserDontExistError create a new error UserDontExist with name of user
func NewUserDontExistError(nameUser string) UserDontExist {
	return errors.New("User of the name: " + nameUser + " don't exists in repository")
}
