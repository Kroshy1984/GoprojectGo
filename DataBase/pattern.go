package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
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
	Id         int
	First_name string
	Last_name  string
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
		err := rows.Scan(&p.Id, &p.First_name, &p.Last_name)
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
	err := row.Scan(&p.Id, &p.First_name, &p.Last_name)
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
		if p.Id != 0 {
			fmt.Print(p.Id)
		}
		fmt.Println(p.First_name, p.Last_name)
	}
}

func Select(dbPath string, query string) ([]person, error) {

	if strings.Contains(query, "drop") || strings.Contains(query, "delete") {
		return []person{}, errors.New("cannot execute query")
	}

	db, err := sqlx.Connect("sqlite3", dbPath)

	if err != nil {
		fmt.Println(err)
		return []person{}, err
	}

	persons := []person{}
	er := db.Select(&persons, query)

	if er != nil {
		fmt.Println(err)
		return []person{}, err
	}

	return persons, nil
}

func SelectFromFile(dbPath string, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	result := ""

	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		result += string(data[:n])
		/*
			поскольку данные представляют срез байтов,
			хотя файл *.txt хранит текстовую информацию,
			то преобразуем срез байтов в строку:
			string(data[:n])
		*/
	}
	fmt.Print("query:\n" + result + "\n")
	fmt.Print("result:\n")
	p, err := Select(dbPath, result)
	if err != nil {
		fmt.Println(err)
		return
	}

	PrintPersons(p)
}

func main() {
	dataSourceName := "test.db"
	InitDB(dataSourceName)

	query := "select First_name, Last_name from Persons order by id desc"
	// // query := "drop table"
	// persons, error := SelectAllPersons(dataSourceName, query)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// PrintPersons(persons)

	_, err := Select(dataSourceName, query)
	if err != nil {
		fmt.Println(err)
		return
	}

	//PrintPersons(p)

	// reading from text file
	queryPath := "test_query.txt"
	SelectFromFile(dataSourceName, queryPath)

}
