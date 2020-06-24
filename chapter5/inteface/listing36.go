//sample program to show how to use an interface in Go.
package main

import "fmt"

//notifier is an interface that defined notification
//type behavior
type notifier interface {
	notify()
}

//user defines a user in the program
type user struct {
	name  string
	email string
}

//notify implements a method with a pointer receiver
func (self *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", self.name, self.email)
}

//main is the entry point for the application
func main() {
	//creat a value of type user and send a notification
	u := user{"Bill", "bill@email.com"}
	//u.notify()
	sendNotification(&u)
}

//sendNotification accepts values that implements the notifier
//interface and sends notifications
func sendNotification(n notifier) {
	n.notify()
}
