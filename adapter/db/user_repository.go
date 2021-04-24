package db

import (
	"context"
	"user-api/entity/user"
	"user-api/usecase"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbgateway_return_error_to_save        = "DBGateway returned error when tried save user"
	error_map_entity_to_db                = "Error when try to map entity to db"
	nil_dbgateway_requested               = "DbGateway requested for create new UserRepository is nil"
	dbgateway_return_error_to_find_user   = "Error to find a user by name in repository"
	dbgateway_return_error_to_delete_user = "Error to delete user in repository"
)

type userRepository struct {
	db NoSQLDB
}

func NewUserRepository(db *NoSQLDB) (usecase.UserRepository, error) {

	if db == nil {
		return nil, errors.New(nil_dbgateway_requested)
	}

	return &userRepository{*db}, nil
}

func (repo *userRepository) FindByName(ctx context.Context, name string) (user.User, error) {

	dbUser, err := repo.
		db.
		FindUserByName(ctx, name)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, usecase.NewUserDontExistError(name)
	}

	if err != nil {
		return nil, errors.Wrap(err, dbgateway_return_error_to_find_user)
	}

	return dBToEntity(dbUser, dbUser.Id)
}

func (repo *userRepository) Save(ctx context.Context, usr user.User) (user.User, error) {

	userDb, err := entityToDB(usr)

	if err != nil {
		return nil, errors.Wrap(err, error_map_entity_to_db)
	}

	idPersistedUser, err := repo.db.SaveUser(ctx, userDb)

	if err != nil {
		return nil, errors.Wrap(err, dbgateway_return_error_to_save)
	}

	return dBToEntity(userDb, idPersistedUser)
}

func (repo *userRepository) Delete(ctx context.Context, name string) error {

	err := repo.db.DeleteUser(ctx, name)

	if err != nil {
		return errors.Wrap(err, dbgateway_return_error_to_delete_user)
	}

	return nil
}

func dBToEntity(db User, id primitive.ObjectID) (user.User, error) {

	return user.
		NewBuilder().
		Id(id.Hex()).
		Name(db.Name).
		Email(db.Email).
		Age(db.Age).
		AddressesIds(db.AddressesIds).
		Build()
}

func entityToDB(entity user.User) (User, error) {

	var id primitive.ObjectID = primitive.NilObjectID
	var err error

	if len(entity.Id()) < 0 {
		id, err = primitive.ObjectIDFromHex(entity.Id())
	}

	if err != nil {
		return User{}, errors.Wrap(err, error_map_entity_to_db)
	}

	return User{
			id,
			entity.Name(),
			entity.Email(),
			entity.Age(),
			entity.AddressesIds(),
		},
		err
}
