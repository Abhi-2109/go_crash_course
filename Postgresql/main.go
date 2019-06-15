package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "*******"
	DB_NAME     = "Learning"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM link")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var link_id int
		var new_title_name string
		var url string

		err = rows.Scan(&link_id, &new_title_name, &url)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%v %v %v\n", link_id, new_title_name, url)
	}

	fmt.Println("Mission SucccessFul")

	defer db.Close()
}
