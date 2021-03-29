package usecase

// CreateUserRequest represents a request for user creation
type CreateUserRequest interface {
	Name() string
	Email() string
	Age() int
}
