package address

import (
	"errors"
	"strconv"
)

type Builder interface {
	ID(id string) Builder
	Street(street string) Builder
	Number(number int) Builder
	Neighborhood(neighborhood string) Builder
	ZipCode(zipcode int) Builder
	Build() (Address, error)
}

type builder struct {
	id           string
	street       string
	number       int
	neighborhood string
	zipcode      int
}

func NewBuilder() Builder {
	return &builder{}
}

func (b *builder) ID(id string) Builder {
	b.id = id
	return b
}

func (b *builder) Street(street string) Builder {
	b.street = street
	return b
}

func (b *builder) Number(number int) Builder {
	b.number = number
	return b
}

func (b *builder) Neighborhood(neighborhood string) Builder {
	b.neighborhood = neighborhood
	return b
}

func (b *builder) ZipCode(zipcode int) Builder {
	b.zipcode = zipcode
	return b
}

func (b *builder) Build() (Address, error) {

	var err error

	err = b.validateNewRequiredParams()

	if err != nil {
		return nil, err
	}

	return address{
		b.id,
		b.street,
		b.number,
		b.neighborhood,
		b.zipcode,
	}, nil
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
