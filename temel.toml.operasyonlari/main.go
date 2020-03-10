package main

import (
	"database/sql"
	"fmt"

	. "./models"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

//postgres://username:password@localhost/db_name?sslmode-disable

func main() {

	// example -1
	var conf Config
	if _, err := toml.DecodeFile("./configurations/config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%#v\n", conf)

	// example -2

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode-disable", conf.Database.User, conf.Database.Password, conf.Database.Server, conf.Database.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT 5 + 5").Scan(&count)
	if err != nil {
		panic(err)
	}

	fmt.Println(count)

}
