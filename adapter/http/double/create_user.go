package double

import "user-api/usecase"

type createUserDouble struct {
	createFunc func(request usecase.CreateUserRequest) (usecase.CreateUserResponse, error)
}

func NewCreateUser(create func(request usecase.CreateUserRequest) (usecase.CreateUserResponse, error)) usecase.CreateUser {
	return &createUserDouble{createFunc: create}
}

func (create *createUserDouble) Create(request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
	return create.createFunc(request)
}

type createUserResponse struct {
	idFunc func() string
}

func NewCreateUserResponse(idFunc func() string) usecase.CreateUserResponse {
	return &createUserResponse{idFunc: idFunc}
}

func (response *createUserResponse) Id() string {
	return response.idFunc()
}
