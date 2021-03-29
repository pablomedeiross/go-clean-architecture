package double

import (
	"user-api/entity/user"
	"user-api/usecase"
)

type userRepositoryDouble struct {
	SaveFunc       func(user user.User) (user.User, error)
	FindByNameFunc func(name string) user.User
}

func NewUserRepositoryDouble(

	save func(user user.User) (user.User, error),
	findByName func(name string) user.User,

) *usecase.UserRepository {

	var use usecase.UserRepository
	use = &userRepositoryDouble{save, findByName}

	return &use
}

func (repository *userRepositoryDouble) FindByName(name string) user.User {
	return repository.FindByNameFunc(name)
}

func (repository *userRepositoryDouble) Save(user user.User) (user.User, error) {
	return repository.SaveFunc(user)
}
