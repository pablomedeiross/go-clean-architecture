package memory

import (
	"context"
	"time"

	"github.com/benweissmann/memongo"
	"github.com/benweissmann/memongo/memongolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InMemoryMongoDB is a database in memory representation
type InMemoryMongoDB struct {
	server *memongo.Server
	nameDB string
}

// NewInMemoryMongoDB create a new instance of InMemoryMongoDB
func NewInMemoryMongoDB() *InMemoryMongoDB {

	return &InMemoryMongoDB{
		nameDB: memongo.RandomDatabase(),
	}
}

// Start start a database execution
func (helper *InMemoryMongoDB) Start() error {

	server, err := createInMemoryMongoDBServer()

	if err != nil {
		return err
	}

	helper.server = server
	return nil
}

// Stop stop a database execution
func (helper *InMemoryMongoDB) Stop() {
	helper.server.Stop()
}

// URI return the uri of database mongoDB
func (helper *InMemoryMongoDB) URI() string {
	return helper.server.URI()
}

// URI return the name of database mongoDB
func (helper *InMemoryMongoDB) Name() string {
	return helper.nameDB
}

// ConnectInCollection return a new collection connect for searchs, deletions, updates or inserts
// For to use this method is necessary that InMemoryMongoDB was had started
func (helper *InMemoryMongoDB) ConnectInCollection(nameCollection string) (*mongo.Collection, error) {

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

func createInMemoryMongoDBServer() (*memongo.Server, error) {

	server, err := memongo.StartWithOptions(
		&memongo.Options{
			MongoVersion:   "4.0.5",
			StartupTimeout: 1 * time.Second,
			LogLevel:       memongolog.LogLevelDebug,
		},
	)

	return server, err
}
