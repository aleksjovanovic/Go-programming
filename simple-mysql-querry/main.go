package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	PersonId  int
	LastName  string
	FirstName string
	Address   string
	City      string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/test_db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Succesfully connected to MySQL database")

	results, err := db.Query("SELECT * FROM persons")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var person Person
		err := results.Scan(&person.PersonId, &person.LastName, &person.FirstName, &person.Address, &person.City)

		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%v\n", person)

	}
}
