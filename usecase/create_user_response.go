package usecase

import "errors"

// CreateUserResponse is a response for user creation
type CreateUserResponse interface {
	Id() string
}

type createUserResponse struct {
	id string
}

func newCreateUserReponse(id string) (CreateUserResponse, error) {

	if len(id) <= 0 {
		return nil, errors.New("id nil to the create CreateUserResponse")
	}

	return createUserResponse{id}, nil
}

// Id return id of created user
func (response createUserResponse) Id() string {
	return response.id
}
