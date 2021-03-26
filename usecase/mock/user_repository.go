package mock

import (
	"user-api/entity/user"
	"user-api/usecase"
)

type userRepositoryMock struct {
	SaveFunc       func(user user.User) (user.User, error)
	FindByNameFunc func(name string) user.User
}

func NewUserRepositoryMock(
	save func(user user.User) (user.User, error),
	findByName func(name string) user.User) *usecase.UserRepository {

	var use usecase.UserRepository
	use = &userRepositoryMock{save, findByName}

	return &use
}

func (repository *userRepositoryMock) FindByName(name string) user.User {
	return repository.FindByNameFunc(name)
}

func (repository *userRepositoryMock) Save(user user.User) (user.User, error) {
	return repository.SaveFunc(user)
}
