package double

import (
	"user-api/adapter/db"
)

type dbGatewayDouble struct {
	SaveUserFunc       func(user db.User) (db.User, error)
	FindUserByNameFunc func(name string) db.User
}

func NewDBGateway(

	saveUser func(user db.User) (db.User, error),
	findUserByName func(name string) db.User,

) db.DBGateway {

	return &dbGatewayDouble{
		SaveUserFunc:       saveUser,
		FindUserByNameFunc: findUserByName,
	}
}

func (gateway *dbGatewayDouble) SaveUser(user db.User) (db.User, error) {
	return gateway.SaveUserFunc(user)
}

func (gateway *dbGatewayDouble) FindUserByName(name string) db.User {
	return gateway.FindUserByNameFunc(name)
}
