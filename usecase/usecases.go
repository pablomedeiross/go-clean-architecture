package usecase

import (
	"context"
)

// SearchUser is a usecase for users research in application
type SearchUser interface {
	Search(ctx context.Context, request PagedUserRequest) (PagedUserResponse, error)
}

// CreateUser is a usecase for creation of new user in application
type CreateUser interface {
	Create(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error)
}

// RemoveUser is a usecase for remove of user in application
type RemoveUser interface {
	Remove(ctx context.Context, request RemoveUserRequest) error
}
