package db

import (
	"user-api/entity/user"
	"user-api/usecase"

	"github.com/pkg/errors"
)

const (
	dbgateway_return_error_to_save = "DBGateway returned error when tried save user"
	user_creation_with_error       = "Error when to try create new user from return DbGateway"
	nil_dbgateway_requested        = "DbGateway requested for create new UserRepository is nil"
)

type userRepository struct {
	dbGateway DBGateway
}

func NewUserRepository(dbGateway DBGateway) (usecase.UserRepository, error) {

	if dbGateway == nil {
		return nil, errors.New(nil_dbgateway_requested)
	}

	return &userRepository{dbGateway}, nil
}

func (repo *userRepository) FindByName(name string) user.User {
	return nil
}

func (repo *userRepository) Save(usr user.User) (user.User, error) {

	userDb := User{
		usr.Id(),
		usr.Name(),
		usr.Email(),
		usr.Age(),
		usr.AddressesIds(),
	}

	userDbReturn, err :=
		repo.dbGateway.SaveUser(userDb)

	if err != nil {
		return nil, errors.Wrap(err, dbgateway_return_error_to_save)
	}

	userReturn, err := user.
		NewBuilder().
		Id(userDbReturn.Id).
		Name(userDbReturn.Name).
		Email(userDbReturn.Email).
		Age(userDbReturn.Age).
		AddressesIds(userDbReturn.AddressesIds).
		Build()

	if err != nil {
		return nil, errors.Wrap(err, user_creation_with_error)
	}

	return userReturn, nil
}
