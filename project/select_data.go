package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	id         int
	first_name string
	last_name  string
}

func main() {

	db, err := sql.Open("sqlite3", ".\\DataBase\\test_db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Persons")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	persons := []person{}

	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.id, &p.first_name, &p.last_name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p)
	}
	for _, p := range persons {
		fmt.Println(p.id, p.first_name, p.last_name)
	}
}
