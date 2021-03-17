package address

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
