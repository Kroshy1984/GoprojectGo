package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Long-lived global variable
var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error

	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

// Adds Person by first and last name
// Returns pair <object_id, error>
func AddPerson(firstName string, lastName string) (int64, error) {
	result, err := db.Exec("insert into Persons (first_name, last_name) values ($1, $2)",
		firstName, lastName)
	if err != nil {
		panic(err)
	}
	return result.LastInsertId()
}

// Describes Person object
type person struct {
	id         int
	first_name string
	last_name  string
}

// Select all Person objects from DB
// Returns an array of Person
func SelectAllPersons(dbPath string, query string) ([]person, error) {

	if strings.Contains(query, "drop") || strings.Contains(query, "delete") {
		return []person{}, errors.New("cannot execute query")
	}

	InitDB(dbPath)

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	persons := []person{}
	// Parse rows into a Persons
	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.id, &p.first_name, &p.last_name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p)
	}
	return persons, nil
}

// Select specified Person by id
// Returns an Person object
func SelectPersonById(id int64) person {
	row := db.QueryRow("select * from Persons where id = $1", id)
	p := person{}
	err := row.Scan(&p.id, &p.first_name, &p.last_name)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

// Updates Person with your id first and last name
// Returns pair <count of updated rows, error>
func UpdatePersonById(id int64, first_name string, last_name string) (int64, error) {
	result, err := db.Exec("update Persons set first_name = $1, last_name = $2 where id = $3", first_name, last_name, id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}

// Deletes Person by selected id
// Returns pair <count of deleted rows, error>
func DeletePersonById(id int64) (int64, error) {
	result, err := db.Exec("delete from Persons where id = $1", id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}

// Prints Person's data from an array
func PrintPersons(persons []person) {
	for _, p := range persons {
		fmt.Println(p.id, p.first_name, p.last_name)
	}
}

func main() {
	dataSourceName := "test.db"
	InitDB(dataSourceName)

	query := "select * from Persons"
	// query := "drop table"
	persons, error := SelectAllPersons(dataSourceName, query)
	if error != nil {
		fmt.Println(error)
		return
	}
	PrintPersons(persons)
}
