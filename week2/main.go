package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Customer struct {
	CustomerId string
	Name       string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
}

func QueryCustomerById(id string) (Customer, error) {
	var customer Customer
	row := Db.QueryRow("select id ,name from customer where id = ?", id)
	err := row.Scan(&customer.CustomerId, &customer.Name)
	if err != nil {
		return customer, errors.Wrap(err, "Error! QueryCustomerById has an error")
	}
	return customer, nil
}

func main() {
	defer Db.Close()
	customer, err := QueryCustomerById("123456")
	if err != nil {
		fmt.Printf("query customer err : %+v", err)
		return
	}
	fmt.Println("query customer : ", customer)
}
