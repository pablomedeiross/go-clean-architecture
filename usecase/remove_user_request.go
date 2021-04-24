package usecase

import "errors"

type removeUserRequest struct {
	name string
}

// NewRemoveUserRequest is a factory for RemoveUserRequest
func NewRemoveUserRequest(name string) (RemoveUserRequest, error) {

	if len(name) <= 0 {
		return nil, errors.New("name is a requested param to create RemoveUserRequest")
	}

	return removeUserRequest{name}, nil
}

// Return name of user that want remove on application
func (request removeUserRequest) Name() string {
	return request.name
}
