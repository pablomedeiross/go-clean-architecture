package domain

import (
	"errors"
	"strconv"
)

// User is a Aggregate root entity
type User struct {
	name      string
	email     string
	age       int
	addresses []Address
}

// NewUser is a factory of User
func NewUser(name string, email string, age int, addresses []Address) (*User, error) {

	var us *User
	var err error

	if len(name) <= 0 || len(email) <= 0 || age <= 0 {
		err = errors.New("Error creating new User with arguments : " + name + ", " + email + ", " + strconv.Itoa(age))

	} else {
		us = &User{
			name:      name,
			email:     email,
			age:       age,
			addresses: addresses,
		}
	}

	return us, err
}

// Name return User name
func (us User) Name() string {
	return us.name
}

// Email return User email
func (us User) Email() string {
	return us.email
}

// Age return User age
func (us User) Age() int {
	return us.age
}

// Addresses return User addresses
func (us User) Addresses() []Address {
	return us.addresses
}

// AddAddress include new address in User
func (us *User) AddAddress(address Address) error {

	var err error

	if exists := addressExists(address, us.addresses); exists {
		err = errors.New("Address already exists in User: " + us.Name())
	}

	us.addresses = append(us.addresses, address)

	return err
}

func addressExists(add Address, addresses []Address) bool {

	exist := false

	for _, actual := range addresses {

		if actual == add {
			exist = true
		}
	}

	return exist
}
