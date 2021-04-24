package usecase

import (
	"context"

	"github.com/pkg/errors"
)

type removeUserInteractor struct {
	repository UserRepository
}

// NewRemoveUser is a fatory for usecase RemoveUser
func NewRemoveUser(repository UserRepository) (RemoveUser, error) {

	if repository == nil {
		return nil, errors.New("UserRepository is a requested param to create RemoveUserInteractor")
	}

	return &removeUserInteractor{repository}, nil
}

// Remove remove user on application
func (interactor *removeUserInteractor) Remove(ctx context.Context, request RemoveUserRequest) error {

	if request == nil {
		return errors.New("RemoveUserRequest is a requested param to remove a user")
	}

	err := interactor.repository.Delete(ctx, request.Name())

	if err != nil {
		return errors.Wrap(err, "Error to try remove a user in RemoveUserInteractor")
	}

	return nil
}
