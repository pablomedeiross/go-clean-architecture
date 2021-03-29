package usecase

import (
	"user-api/entity/address"
	"user-api/entity/user"
)

// UserRepository is an interface that represents the data persistent for user entities
type UserRepository interface {
	FindByName(name string) user.User
	Save(user user.User) (user.User, error)
}

// AddressRepository is an interface that represents the data persistent for address entities
type AddressRepository interface {
	Save(address address.Address) (address.Address, error)
}
