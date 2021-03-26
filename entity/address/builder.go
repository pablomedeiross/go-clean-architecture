package address

import (
	"errors"
	"strconv"
)

// Builder of Address entity
type Builder interface {
	Id(id string) Builder
	Street(street string) Builder
	Number(number int) Builder
	Neighborhood(neighborhood string) Builder
	Zipcode(zipcode int) Builder
	Build() (Address, error)
}

type builder struct {
	id           string
	street       string
	number       int
	neighborhood string
	zipcode      int
}

// NewBuilder return a new builder of User entity
func NewBuilder() Builder {
	return &builder{}
}

// Id set a id for the entity on building
func (b *builder) Id(id string) Builder {
	b.id = id
	return b
}

// Street set a street for the entity on building
func (b *builder) Street(street string) Builder {
	b.street = street
	return b
}

// Number set a number for the entity on building
func (b *builder) Number(number int) Builder {
	b.number = number
	return b
}

// Neighborhood set a neighborhood for the entity on building
func (b *builder) Neighborhood(neighborhood string) Builder {
	b.neighborhood = neighborhood
	return b
}

// Zipcode set a zipcode for the entity on building
func (b *builder) Zipcode(zipcode int) Builder {
	b.zipcode = zipcode
	return b
}

// Build return a new Address or a error in case of invalid arguments
func (b *builder) Build() (Address, error) {

	var addr Address
	var err error = b.validateNewRequiredParams()

	if err == nil {
		addr = address{
			b.id,
			b.street,
			b.number,
			b.neighborhood,
			b.zipcode,
		}
	}

	return addr, err
}

func (b builder) validateNewRequiredParams() error {

	var err error

	if len(b.street) <= 0 || b.number <= 0 || len(b.neighborhood) <= 0 || b.zipcode <= 0 {

		err = errors.New(
			"Error creating new Address with arguments : " +
				b.street + ", " +
				strconv.Itoa(b.number) + ", " +
				b.neighborhood + ", " +
				strconv.Itoa(b.zipcode))
	}

	return err
}
