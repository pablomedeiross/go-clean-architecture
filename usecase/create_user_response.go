package usecase

import "errors"

type CreateUserResponse interface {
	ID() int
}

type createUserResponse struct {
	id int
}

// NewCreateUserReponse is a factory of CreateUserResponse 
func NewCreateUserReponse(id int) (CreateUserResponse, error) {

	var response CreateUserResponse
	var err error

	if id == 0 {
		err = errors.New("id null to the create CreateUserResponse")
		
	} else {
		response = createUserResponse{id: id,}
	}

	return response, err
}

// ID return CreateUserResponse id
func (response createUserResponse) ID() int {
	return response.id
}
