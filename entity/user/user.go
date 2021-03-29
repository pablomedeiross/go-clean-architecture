package user

import (
	"errors"
)

// User entity implemantation
type user struct {
	id           string
	name         string
	email        string
	age          int
	addressesIds []string
}

// User entity
type User interface {
	Id() string
	Name() string
	Email() string
	Age() int
	AddressesIds() []string
	AddAddressId(addressId string) error
}

// Id return User id
func (us user) Id() string {
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
func (us user) AddressesIds() []string {
	return us.addressesIds
}

// AddAddressID include new address in User
func (us *user) AddAddressId(addressId string) error {

	var err error

	if exists := addressExists(addressId, us.addressesIds); exists {
		err = errors.New("AddressId already exists in User: " + us.Name())
	}

	us.addressesIds = append(us.addressesIds, addressId)

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
