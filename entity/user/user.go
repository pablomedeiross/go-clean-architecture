package user

import (
	"errors"
	"strconv"
)

// User entity implemantation
type user struct {
	id           int
	name         string
	email        string
	age          int
	addressesIds []string
}

// User entity
type User interface {
	ID() int
	Name() string
	Email() string
	Age() int
	AddressesIDs() []string
	AddAddressID(AddAddressID string) error
}

// New is a factory of User
func New(name string, email string, age int) (User, error) {
	return createUser(name, email, age)
}


// NewPersisted is a factory of User
func NewPersisted(id int, name string, email string, age int, addressesIds []string) (User, error) {

	var us User
	var err error

	if len(strconv.Itoa(id)) < 11 || addressesIds == nil || len(addressesIds) < 0 {
		err = errors.New("Error creating new User with arguments : " + name + ", " + email + ", " + strconv.Itoa(age))

	} else {
		us, err = createUser(name, email, age)

		if err == nil {

			us = user {
				id: id,
				name: us.Name(),
				email: us.Email(),
				age: us.Age(),
				addressesIds: addressesIds,
			}
		}
	}

	return us, err
}

func createUser(name string, email string, age int) (User, error) {

	var us User
	var err error

	if len(name) <= 0 || len(email) <= 0 || age <= 0 {
		err = errors.New("Error creating new User with arguments : " + name + ", " + email + ", " + strconv.Itoa(age))

	} else {
		us = user{
			name:  name,
			email: email,
			age:   age,
		}
	}

	return us, err
}

// Id return User id
func (us user) ID() int {
	return us.id
}

// Name return User name
func (us user) Name() string {
	return us.name
}

// Email return User email
func (us user) Email() string {
	return us.email
}

// Age return User age
func (us user) Age() int {
	return us.age
}

// AddressesIDs return User addresses
func (us user) AddressesIDs() []string {
	return us.addressesIds
}

// AddAddressID include new address in User
func (us user) AddAddressID(addressID string) error {

	var err error

	if exists := addressExists(addressID, us.addressesIds); exists {
		err = errors.New("AddressId already exists in User: " + us.Name())
	}

	us.addressesIds = append(us.addressesIds, addressID)

	return err
}

func addressExists(add string, addresses []string) bool {

	exist := false

	for _, actual := range addresses {

		if actual == add {
			exist = true
		}
	}

	return exist
}
