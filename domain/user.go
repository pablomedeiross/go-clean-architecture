package domain

// User entity
type User struct {
	Name    string
	Email   string
	Age     int8
	Address []Address
}

// AddAddress include new address in User
func (user *User) AddAddress(address Address) {

	// TODO: Validations
	user.Address = append(user.Address, address)
}
