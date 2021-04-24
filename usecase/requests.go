package usecase

// PagedUserRequest represent a request to user's search
type PagedUserRequest interface {
	Page() int
	PageSize() int
	NumberOfPages() int
	TotalRegisters() int
}

// CreateUserRequest represents a request for user creation
type CreateUserRequest interface {
	Name() string
	Email() string
	Age() int
}

//RemoveUserRequest represents a request for user deletion
type RemoveUserRequest interface {
	Name() string
}
