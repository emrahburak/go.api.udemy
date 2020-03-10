package main

import "fmt"
import "os"

func main() {

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	uName := os.Getenv("GOPATH")
	uDomain := os.Getenv("MYPATH")

	fmt.Println(uName, uDomain)
}
