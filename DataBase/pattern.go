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

// car struct
type car struct {
	id           int
	manufacturer string
	model        string
	product_year int
}

// Adds Car by all needed info
// Returns pair <object_id, error>
func AddCar(id int64, manufacturer string, model string, product_year int64) (int64, error) {
	result, err := db.Exec("insert into CARS (id, manufacturer, model, production_year) values ($1, $2, $3, $4)",
		id, manufacturer, model, product_year)
	if err != nil {
		panic(err)
	}
	return result.LastInsertId()
}

// Select all Car objects from DB
// Returns an array of Car
func SelectAllCars() []car {
	rows, err := db.Query("select * from CARS")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	cars := []car{}
	// Parse rows into a Cars
	for rows.Next() {
		c := car{}
		err := rows.Scan(&c.id, &c.manufacturer, &c.model, &c.product_year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cars = append(cars, c)
	}
	return cars
}

// Select specified Car by id
// Returns an Car object
func SelectCarById(id int64) car {
	row := db.QueryRow("select * from CARS where id = $1", id)
	c := car{}
	err := row.Scan(&c.id, &c.manufacturer, &c.model, &c.product_year)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

// Updates Car with input data
// Returns pair <count of updated rows, error>
func UpdateCarById(id int64, manufacturer string, model string, product_year int) (int64, error) {
	result, err := db.Exec("update CARS set manufacturer = $1, model = $2, production_year = $3 where id = $4", manufacturer, model, product_year, id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}

// Deletes Car by id
// Returns pair <count of deleted rows, error>
func DeleteCarById(id int64) (int64, error) {
	result, err := db.Exec("delete from CARS where id = $1", id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}

// Prints Car data from an array
func PrintCars(cars []car) {
	for _, c := range cars {
		fmt.Println(c.id, c.manufacturer, c.model, c.product_year)
	}
}

func main() {
	dataSourceName := "cars.db"
	InitDB(dataSourceName)

	//add demo
	id, err := AddCar(3, "McLaren", "MP35M", 2021)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added new Car: ")
	PrintCars([]car{SelectCarById(id)})

	fmt.Println()

	//select all demo
	PrintCars(SelectAllCars())

	fmt.Println()

	//update demo
	count, err := UpdateCarById(id, "Aston Martin", "AMR21", 2021)
	if err != nil {
		panic(err)
	}
	fmt.Print("Updated ", count, " rows\n")
	PrintCars(SelectAllCars())

	fmt.Println()

	//delete demo
	count, err = DeleteCarById(id)
	if err != nil {
		panic(err)
	}
	fmt.Print("Deleted ", count, " rows\n")
	PrintCars(SelectAllCars())
}
