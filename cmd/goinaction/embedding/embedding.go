// this package demonstrate embedding and export
package main

// import "fmt"

import (
	"github.com/jzpin/learn/pkg/entities"
)

type notifier interface {
	notify()
}

func sendNotification(u notifier) {
	u.notify()
}

type admin struct {
	entities.User
	level string
}

func main() {
	u := admin{User: entities.User{Name: "Bob", Email: "Bob@gmail.com"}, level: "High"}
	sendNotification(&u) // because interface defined for *user, here we must take address
}
