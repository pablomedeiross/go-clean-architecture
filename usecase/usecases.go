package usecase

import "context"

// CreateUser is a usecase for creation of new user in application
type CreateUser interface {
	Create(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error)
}

// RemoveUser is a usecase for remove of user in application
type RemoveUser interface {
	Remove(ctx context.Context, request RemoveUserRequest) error
}
