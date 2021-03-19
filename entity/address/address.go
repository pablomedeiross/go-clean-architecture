package address

// Address entity interface
type Address interface {
	Id() string
	Street() string
	Number() int
	Neighborhood() string
	Zipcode() int
}

// address is a entity implementation
type address struct {
	id           string
	street       string
	number       int
	neighborhood string
	zipcode      int
}

// GetID return Address identifier
func (ad address) Id() string {
	return ad.id
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
func (ad address) Zipcode() int {
	return ad.zipcode
}
