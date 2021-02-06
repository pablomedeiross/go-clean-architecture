package domain

import (
	"errors"
	"strconv"
)

// Address is a entity
type Address struct {
	street       string
	number       int
	neighborhood string
	zipCode      int
}

// NewAddress is a factory for Address
func NewAddress(street string, number int, neighborhood string, zipCode int) (*Address, error) {

	var addr *Address
	var err error

	if len(street) <= 0 || number <= 0 || len(neighborhood) <= 0 || zipCode <= 0 {

		err = errors.New(
			"Error creating new Address with arguments : " +
				street + ", " +
				strconv.Itoa(number) + ", " +
				neighborhood + ", " +
				strconv.Itoa(zipCode))

	} else {
		addr = &Address{
			street:       street,
			number:       number,
			neighborhood: neighborhood,
			zipCode:      zipCode,
		}
	}

	return addr, err
}

// Street return Address Street
func (ad Address) Street() string {
	return ad.street
}

// Number return Address Number
func (ad Address) Number() int {
	return ad.number
}

// Neighborhood return Address Neighborhood
func (ad Address) Neighborhood() string {
	return ad.neighborhood
}

// ZipCode return Address ZipCodeilk
func (ad Address) ZipCode() int {
	return ad.zipCode
}
