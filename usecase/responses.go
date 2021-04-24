package usecase

import "os/user"

// CreateUserResponse is a response for user creation
type CreateUserResponse interface {
	Id() string
}

// PagedUserResponse represents a users paginated result from repository
type PagedUserResponse interface {
	Content() []user.User
	Page() int
	PageSize() int
	NumberOfPages() int
	TotalRegisters() int
}
