package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("Openning file error: ", err)
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var c Company
	xml.Unmarshal(xmlData, &c)

	fmt.Println(c.Persons)

	// Convert to Json

	var person JsonPerson
	var persons []JsonPerson

	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.UserName = value.UserName

		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	// ekrana yaz
	fmt.Println(string(jsonData))

	//dosyaya yaz
	jsonFile, err := os.Create("./Employees.json")
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

}

type JsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string {
	return fmt.Sprintf("\t ID: %d - FirstName : %s - LastName : %s - UserName : %s \n", p.ID, p.FirstName, p.LastName, p.UserName)

}
