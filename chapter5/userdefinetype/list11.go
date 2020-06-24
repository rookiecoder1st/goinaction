//sample program to show how to declare methods and
//how the Go compiler support them.

package main

import "fmt"

//user defines a user int the program
type userInfo struct {
	name  string
	email string
}

//notify implements a method with a value receiver
func (u userInfo) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

//changeEmail implements a method with a pointer receiver
func (u *userInfo) changeEmail(email string) {
	u.email = email
}

//main is the entry point for the applications

func main() {
	//Values of type user can be used to call methods
	//declared with a value receiver
	bill := userInfo{
		"Bill",
		"bill@email.com",
	}
	bill.notify()

	//pointers of type user can also be used to call methods
	//declare with a value reciever
	lisa := &userInfo{
		"Lisa",
		"lisa@email.com",
	}
	lisa.notify()
	//values of type user can be used to call methods
	//declared with a value receiver
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	//pointers of type user acn be used to call methods
	//declared with a pointer receiver
	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()
}
