package address

import (
	"errors"
	"strconv"
)

// Address entity interface
type Address interface {
	GetID() string
	Street() string
	Number() int
	Neighborhood() string
	ZipCode() int
}

// address is a entity implementation
type address struct {
	ID           string
	street       string
	number       int
	neighborhood string
	zipCode      int
}

// New is a factory for Address
func New(street string, number int, neighborhood string, zipCode int) (Address, error) {
	return build(street, number, neighborhood, zipCode)
}

// NewPersisted is a factory method for Address persisted
func NewPersisted(ID string, street string, number int, neighborhood string, zipCode int) (Address, error) {

	var addr Address
	var err error

	if len(ID) == 0 {
		err = errors.New("Error creating new Address, ID is null")

	} else if err == nil {

		addr, err = build(street, number, neighborhood, zipCode)

		if err == nil {
			
			addr = address {
				ID : ID,
				street: addr.Street(),
				number: addr.Number(),
				neighborhood: addr.Neighborhood(),
				zipCode: addr.ZipCode(),
			}
		}
	}

	return addr, err
}

func build(street string, number int, neighborhood string, zipCode int) (Address, error) {

	var addr Address
	err := validateNewRequiredParams(street, number, neighborhood, zipCode)

	if err == nil {
		addr = address{
			street:       street,
			number:       number,
			neighborhood: neighborhood,
			zipCode:      zipCode,
		}
	}

	return addr, err
}

func validateNewRequiredParams(street string, number int, neighborhood string, zipCode int) error {

	var err error

	if len(street) <= 0 || number <= 0 || len(neighborhood) <= 0 || zipCode <= 0 {

		err = errors.New(
			"Error creating new Address with arguments : " +
				street + ", " +
				strconv.Itoa(number) + ", " +
				neighborhood + ", " +
				strconv.Itoa(zipCode))
	}

	return err
}

// GetID return Address identifier
func (ad address) GetID() string {
	return ad.ID
}

// Street return Address Street
func (ad address) Street() string {
	return ad.street
}

// Number return Address Number
func (ad address) Number() int {
	return ad.number
}

// Neighborhood return Address Neighborhood
func (ad address) Neighborhood() string {
	return ad.neighborhood
}

// ZipCode return Address ZipCodeilk
func (ad address) ZipCode() int {
	return ad.zipCode
}
