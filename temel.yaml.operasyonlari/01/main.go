package main

// yml uzanılı dosya olarak yaz..

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

func main() {

	p := Person{"emrah", "Burak", 36}

	y, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(y))
}

type Person struct {
	FirstName string `yaml:"firs_name"`
	LastName  string `yaml:"last_name"`
	Age       int    `yaml:"age"`
}
