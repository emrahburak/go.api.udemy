package main

import "fmt"
import "bytes"

func main() {
	var x bytes.Buffer
	for i := 0; i < 10; i++ {
		x.WriteString(rastgeleStiring())
	}
	fmt.Println(x.String())

}

func rastgeleStiring() string {
	return "$xyz-3642"
}
