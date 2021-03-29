package usecase

import (
	"user-api/entity/user"
)

// UserRepository is an interface that represents the data persistent for user entities
type UserRepository interface {
	FindByName(name string) user.User
	Save(user user.User) (user.User, error)
}
