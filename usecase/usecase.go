package usecase

// CreateUser is a usecase for creation of new user in application
type CreateUser interface {
	Create(request CreateUserRequest) (CreateUserResponse, error)
}

// AddAddress include a new address for a user
type AddAddress interface {
	AddAddress(request AddAddressRequest) error
}
