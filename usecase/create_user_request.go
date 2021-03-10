package usecase

import "errors"

type CreateUserRequest interface {

	Name() string
	Email() string
	Age() int
}

type createUserRequest struct {
	name string
	email string
	age int
}

func NewCreateUserRequest(name string, email string, age int) (CreateUserRequest, error) {

	var request CreateUserRequest
	var err error

	if len(name) == 0 || len(email) == 0 || age <= 0 {
		err = errors.New("NewCreateUserRequest called without requested parameter")
	
	} else {
		request = createUserRequest {
			name: name,
			email: email,
			age: age,
		} 
	}

	return request, err
}

// Return CreateUserRequest name
func (request createUserRequest) Name() string {
	return request.name
}

// Return CreateUserRequest email
func (request createUserRequest) Email() string {
	return request.email
}

// Return CreateUserRequest age
func (request createUserRequest) Age() int {
	return request.age
}