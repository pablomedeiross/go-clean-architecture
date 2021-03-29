package usecase

// CreateUserRequest represents a request for user creation
type CreateUserRequest interface {
	Name() string
	Email() string
	Age() int
}

// AddAddressRequest represents a request for address inclusion to a user
type AddAddressRequest interface {
	Street() string
}
