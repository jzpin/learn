package entities

import "fmt"

// User has Name and Email properties
type User struct {
	Name  string
	Email string
}

// implement for *user type
// if a function is defined for pointer, it means the func intends to modify item
// as a result, it can only be used with &item
func (u *User) notify() {
	fmt.Printf("Send email to %s <%s>\n", u.Name, u.Email)
}
