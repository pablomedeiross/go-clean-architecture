package usecase

import "context"

// CreateUser is a usecase for creation of new user in application
type CreateUser interface {
	Create(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error)
}
