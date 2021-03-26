package usecase

import "user-api/entity/user"

// UserRepository is a interface that represents the data persistent
type UserRepository interface {
	FindByName(name string) user.User
	Save(user user.User) (user.User, error)
}
