package main

import "fmt"

type Human struct {
	firstName string
	lastName  string
	age       int
}

// Varsayılan ve boş yapıcı method
func newHuman() *Human {
	h := new(Human)
	return h
}

func main() {
	// Nesne Öreneği oluşturma

	//v1
	// x := Human{firstName: "Emrah"}
	// fmt.Println(x.firstName)

	//v2
	// x := new(Human)
	// x.firstName = "ali"
	// fmt.Println(x.firstName)

	//v3

	x := newHuman()
	x.age = 36
	fmt.Println(x.age)
}
