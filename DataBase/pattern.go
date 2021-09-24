package main

import (
	"database/sql"
	"fmt"

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
func SelectAllPersons() []person {
	rows, err := db.Query("select * from Persons")
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
	return persons
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
	dataSourceName := ".\\DataBase\\test.db"
	InitDB(dataSourceName)
	id, err := AddPerson("Danil", "Komarov")
	if err != nil {
		panic(err)
	}
	fmt.Println("Added new Person: ")
	PrintPersons([]person{SelectPersonById(id)})
	fmt.Println("==========")
	PrintPersons(SelectAllPersons())
	count, err := UpdatePersonById(id, "Mikhail", "Dirin")
	if err != nil {
		panic(err)
	}
	fmt.Print("Updated ", count, " rows\n")
	PrintPersons(SelectAllPersons())
	count, err = DeletePersonById(id)
	if err != nil {
		panic(err)
	}
	fmt.Print("Deleted ", count, " rows\n")
	PrintPersons(SelectAllPersons())
}
