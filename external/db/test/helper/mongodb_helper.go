package helper

import (
	"context"
	"time"
	"user-api/adapter/db"

	"github.com/benweissmann/memongo"
	"github.com/benweissmann/memongo/memongolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const user_collection_name = "User"

type DBHelper struct {
	server *memongo.Server
	nameDB string
}

func NewMongoDBHelper() *DBHelper {

	return &DBHelper{
		nameDB: memongo.RandomDatabase(),
	}
}

// init in memory MongoDB server
func (helper *DBHelper) StartMongoDB() error {

	server, err := CreateInMemoryMongoDBServer()

	if err != nil {
		return err
	}

	helper.server = server
	return nil
}

func (helper *DBHelper) StopMongoDB() {
	helper.server.Stop()
}

func (helper *DBHelper) DatabaseURI() string {
	return helper.server.URI()
}

func (helper *DBHelper) DatabaseName() string {
	return helper.nameDB
}

func CreateInMemoryMongoDBServer() (*memongo.Server, error) {

	server, err := memongo.StartWithOptions(
		&memongo.Options{
			MongoVersion:   "4.0.5",
			StartupTimeout: 1 * time.Second,
			LogLevel:       memongolog.LogLevelDebug,
		},
	)

	return server, err
}

// Insert a user in MongoDB
func (helper *DBHelper) InsertUser(usr db.User) error {

	collection, err := helper.connectInCollection(user_collection_name)
	_, err = collection.InsertOne(context.Background(), &usr, options.InsertOne())

	return err
}

// Search user by name in MongoDB, if don't find return a error
func (helper *DBHelper) FindUserById(id primitive.ObjectID) (db.User, error) {

	var userReturned *db.User = &db.User{}
	collection, err := helper.connectInCollection(user_collection_name)

	if err != nil {
		return *userReturned, nil
	}

	err = collection.
		FindOne(context.Background(), &bson.M{"_id": id}, options.FindOne()).
		Decode(userReturned)

	return *userReturned, err
}

func (helper *DBHelper) connectInCollection(nameCollection string) (*mongo.Collection, error) {

	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(helper.server.URI() + "/" + helper.nameDB))

	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return client.
			Database(helper.nameDB).
			Collection(nameCollection),
		nil

}
