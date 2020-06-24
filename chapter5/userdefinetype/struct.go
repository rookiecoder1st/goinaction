package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

//declare admin composite user and level
type admin struct {
	person user
	level  string
}

func main() {
	//declare a variable of type user
	var bill user

	//declare a variable of type user and initialize all the field
	lisa := user{
		"lisa",
		"lisa@email.com",
		123,
		true,
	}
	fmt.Printf("%+v\n %+v\n", bill, lisa)
	//declare a variable of type admin
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@eamil.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)
}
