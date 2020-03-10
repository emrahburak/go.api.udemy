package main

import (
	"io/ioutil"

	//. "./models" // (.) import nesnesi önünde kullanılırsa paket içindekileri load eder..

	//	"gopkg.in/yaml.v2"
	"fmt"
)

func main() {

	// os.Argv[1]
	// go run main.go config.yaml
	// Araştırma ödevi

	fileName := "./config.yaml"
	//	var config Config
	source, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(source))

	///yaml.Unmarshal(...) ödev

}
