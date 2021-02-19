package usecase

type CreateUserRequest interface {

	Name() string
	Email() string
	Age() int
}
