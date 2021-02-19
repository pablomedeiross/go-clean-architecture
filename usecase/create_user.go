package usecase

type CreateUser interface {
	Create(request CreateUserRequest) CreateUserResponse
}