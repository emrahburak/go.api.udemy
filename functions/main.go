package main

import "fmt"

func main() {

	message := "Hi"
	sayHello(&message)
	fmt.Println(message)

}

func sayHello(message *string) {
	fmt.Println(*message)
	*message = "Hı Go!"

}
