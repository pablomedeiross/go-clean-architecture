package db

type DBGateway interface {
	SaveUser(user User) (User, error)
	FindUserByName(name string) User
}
