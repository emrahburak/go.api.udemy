package main

import "fmt"
import "strconv"

func main() {
	/*
		var massage string
		massage = "Merhaba Go"

	*/
	var a, b, c int = 3, 4, 5
	d := "Metin"
	e := 'K'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	var mystring string = "2"

	number, _ := strconv.Atoi(mystring)
	fmt.Println(number)
	result := number + 2
	fmt.Println(result)
}
