package domain

import (
	"errors"
	"strconv"
)

// User Entity interface
type User interface {
	GetName() string
	AddAddress(address Address) error
}

// User entity implementation
type user struct {
	name    string
	email   string
	age     int
	address []Address
}

// NewUser is a factory of User
func NewUser(name string, email string, age int, address []Address) (User, error) {

	var us User
	var err error

	if len(name) <= 0 || len(email) <= 0 || age <= 0 {
		err = errors.New("Error creating new User with arguments : " + name + ", " + email + ", " + strconv.Itoa(age))

	} else {
		us = &user{
			name:    name,
			email:   email,
			age:     age,
			address: address,
		}
	}

	return us, err
}

// GetName return user name
func (us user) GetName() string {
	return us.name
}

// AddAddress include new address in User
func (us user) AddAddress(address Address) error {

	var err error

	if addressExists(address, us.address) {
		err = errors.New("Address already exists in user: " + us.GetName())
	}

	us.address = append(us.address, address)

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
