package double

import (
	"context"
	"user-api/entity/user"
	"user-api/usecase"
)

type userRepositoryDouble struct {
	SaveFunc       func(ctx context.Context, user user.User) (user.User, error)
	FindByNameFunc func(ctx context.Context, name string) (user.User, error)
}

func NewUserRepositoryDouble(

	save func(ctx context.Context, user user.User) (user.User, error),
	findByName func(ctx context.Context, name string) (user.User, error),

) *usecase.UserRepository {

	var use usecase.UserRepository
	use = &userRepositoryDouble{save, findByName}

	return &use
}

func (repository *userRepositoryDouble) FindByName(ctx context.Context, name string) (user.User, error) {
	return repository.FindByNameFunc(ctx, name)
}

func (repository *userRepositoryDouble) Save(ctx context.Context, user user.User) (user.User, error) {
	return repository.SaveFunc(ctx, user)
}
