package domain

import (
	"errors"
	"strconv"
)

// Address interface entity
type Address interface {
	GetStreet() string
	GetNumber() int
	GetNeighborhood() string
	GetZipCode() int
}

// address is a implementation for Address interface entity
type address struct {
	street       string
	number       int
	neighborhood string
	zipCode      int
}

// NewAddress is a factory for Address
func NewAddress(street string, number int, neighborhood string, zipCode int) (Address, error) {

	var addr Address
	var err error

	if len(street) <= 0 || number <= 0 || len(neighborhood) <= 0 || zipCode <= 0 {

		err = errors.New(
			"Error creating new Address with arguments : " +
				street + ", " +
				strconv.Itoa(number) + ", " +
				neighborhood + ", " +
				strconv.Itoa(zipCode))

	} else {
		addr = &address{
			street:       street,
			number:       number,
			neighborhood: neighborhood,
			zipCode:      zipCode,
		}
	}

	return addr, err
}

func (ad address) GetStreet() string {
	return ad.street
}

func (ad address) GetNumber() int {
	return ad.number
}

func (ad address) GetNeighborhood() string {
	return ad.neighborhood
}

func (ad address) GetZipCode() int {
	return ad.zipCode
}
