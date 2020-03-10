package main

import "fmt"
import "strings"

// String builder ile string birleştirme..

func main() {
	builder := strings.Builder{}

	for i := 0; i < 10; i++ {
		builder.WriteString("Data ")
	}

	result := builder.String()
	fmt.Println(result)
}
