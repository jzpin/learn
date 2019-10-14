// this package create user struct and notifier interface
// This listing demonstrate the concept of method set
// Receiver (t T) -> T and *T
// Receiver (t *T) -> *T
package main

import "fmt"

type user struct {
	name  string
	email string
}

type notifier interface {
	notify()
}

// implement for *user type
// if a function is defined for pointer, it means the func intends to modify item
// as a result, it can only be used with &item
func (u *user) notify() {
	fmt.Printf("Send email to %s <%s>\n", u.name, u.email)
}

func sendNotification(u notifier) {
	u.notify()
}

// type Duration int

func main() {
	u := user{"Bob", "Bob@gmail.com"}
	sendNotification(&u) // because interface defined for *user, here we must take address

	/*
		Rule: not all values are addressable...
		The operand must be addressable, that is, either a variable, pointer indirection,
		or slice indexing operation; or a field selector of an addressable struct operand;
		or an array indexing operation of an addressable array.
	*/
	// address := &Duration(42) // .\notifier.go:35:13: cannot take the address of Duration(42)
	// fmt.Println(address)
}
