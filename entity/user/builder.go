package user

import (
	"errors"
	"strconv"
)

// Builder of User entity
type Builder interface {
	Id(id string) Builder
	Name(name string) Builder
	Email(email string) Builder
	Age(age int) Builder
	AddressesIds(ids []string) Builder
	Build() (User, error)
}

type builder struct {
	id           string
	name         string
	email        string
	age          int
	addressesIds []string
}

// NewBuilder return a new Builder for User entity
func NewBuilder() Builder {
	return &builder{}
}

// Id set a id for the entity on building
func (b *builder) Id(id string) Builder {
	b.id = id
	return b
}

// Name set a id for the entity on building
func (b *builder) Name(name string) Builder {
	b.name = name
	return b
}

// Email set a id for the entity on building
func (b *builder) Email(email string) Builder {
	b.email = email
	return b
}

// Age set a id for the entity on building
func (b *builder) Age(age int) Builder {
	b.age = age
	return b
}

// AddressesIds set a id for the entity on building
func (b *builder) AddressesIds(ids []string) Builder {
	b.addressesIds = ids
	return b
}

// Build return a new User or a error in case of invalid arguments
func (b *builder) Build() (User, error) {

	err := b.validateRequestParams()

	if err != nil {
		return nil, err
	}

	return &user{
		b.id,
		b.name,
		b.email,
		b.age,
		b.addressesIds,
	}, err
}

func (b builder) validateRequestParams() error {

	var err error

	if b.age <= 0 || len(b.name) <= 0 || len(b.email) <= 0 {
		err = errors.New(
			"Error creating new User with arguments : " +
				b.name + ", " +
				b.email + ", " +
				strconv.Itoa(b.age))
	}

	return err
}
