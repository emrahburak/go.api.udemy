package main

import (
	"bufio"
	"fmt"
	"os"
	// 	"strconv"
	// 	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

}
