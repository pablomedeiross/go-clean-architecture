package db

import (
	"context"
	"time"
	"user-api/adapter/db"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoDB_error_to_connect     = "Error to connect in database: "
	mongoDB_error_to_ping        = "Error to ping database after create connection :"
	mongoDB_error_save_user      = "Error to save user in database"
	mongoDB_nill_creat_dbgateway = "Database is a requested param to create DBGateway"
	user_collection_name         = "User"
)

type mongoDBRepository struct {
	db mongo.Database
}

func NewNoSQLDB(urlDatabase string, nameDatabase string) (db.NoSQLDB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(urlDatabase + "/" + nameDatabase))
	err = client.Connect(ctx)

	if err != nil {
		return nil, errors.Wrap(err, mongoDB_error_to_connect)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, errors.Wrap(err, mongoDB_error_to_ping)
	}

	db := client.Database(nameDatabase)

	return &mongoDBRepository{*db}, nil
}

func (repository *mongoDBRepository) SaveUser(ctx context.Context, user db.User) (primitive.ObjectID, error) {

	result, err := repository.
		db.
		Collection(user_collection_name).
		InsertOne(ctx, user, options.InsertOne())

	if err != nil {
		return primitive.NilObjectID, errors.Wrap(err, mongoDB_error_save_user)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (repository *mongoDBRepository) FindUserByName(ctx context.Context, name string) (db.User, error) {

	var usr *db.User = &db.User{}

	err := repository.
		db.
		Collection(user_collection_name).
		FindOne(ctx, bson.M{"name": name}, options.FindOne()).
		Decode(usr)

	return *usr, err
}
