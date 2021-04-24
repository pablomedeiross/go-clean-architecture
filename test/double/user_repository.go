package double

import (
	"context"
	"user-api/entity/user"
	"user-api/usecase"
)

type userRepositoryDouble struct {
	SaveFunc       func(ctx context.Context, user user.User) (user.User, error)
	FindByNameFunc func(ctx context.Context, name string) (user.User, error)
	DeleteFunc     func(ctx context.Context, name string) error
}

func NewUserRepositoryDouble(

	save func(ctx context.Context, user user.User) (user.User, error),
	findByName func(ctx context.Context, name string) (user.User, error),
	delete func(ctx context.Context, name string) error,

) *usecase.UserRepository {

	repo := usecase.UserRepository(&userRepositoryDouble{save, findByName, delete})
	return &repo
}

func (repository *userRepositoryDouble) FindByName(ctx context.Context, name string) (user.User, error) {
	return repository.FindByNameFunc(ctx, name)
}

func (repository *userRepositoryDouble) Save(ctx context.Context, user user.User) (user.User, error) {
	return repository.SaveFunc(ctx, user)
}

func (repository *userRepositoryDouble) Delete(ctx context.Context, name string) error {
	return repository.DeleteFunc(ctx, name)
}
