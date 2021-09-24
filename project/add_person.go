package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", ".\\DataBase\\test_db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into persons (first_name, last_name) values ($1, $2)",
		"Dmitry", "Chernikov")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего добавленного объекта
	fmt.Println(result.RowsAffected()) // количество добавленных строк

}
