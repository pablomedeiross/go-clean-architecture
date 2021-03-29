package db

type User struct {
	Id           string   `bson:"_id,omitempty"`
	Name         string   `bson:"name"`
	Email        string   `bson:"email"`
	Age          int      `bson:"age"`
	AddressesIds []string `bson:"addressesIds,omitempty"`
}
