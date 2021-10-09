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

// Adds Product by first and last name
// Returns pair <object_id, error>

func AddProduct(id int, product string, expiration_date string) (int64, error) {
	result, err := db.Exec("insert into Products (id, product, expiration_date) values ($1, $2, $3)",
		id, product, expiration_date)
	if err != nil {
		panic(err)
	}
	return result.LastInsertId()
}

// Describes Product object
type product struct {
	id              int
	product         string
	expiration_date string
}

// Select all Product objects from DB
// Returns an array of Product
func SelectAllProducts() []product {
	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}
	// Parse rows into a products
	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.product, &p.expiration_date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	return products
}

// Select specified Product by id
// Returns an Product object
func SelectProductById(id int64) product {
	row := db.QueryRow("select * from Products where id = $1", id)
	p := product{}
	err := row.Scan(&p.id, &p.product, &p.expiration_date)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

// Updates Product with your id first and last name
// Returns pair <count of updated rows, error>
func UpdateProductById(id int64, product string, expiration_date string) (int64, error) {
	result, err := db.Exec("update Products set product = $1, expiration_date = $2 where id = $3", product, expiration_date, id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}

// Deletes Product by selected id
// Returns pair <count of deleted rows, error>
func DeleteProductById(id int64) (int64, error) {
	result, err := db.Exec("delete from Products where id = $1", id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}

// Prints Product's data from an array
func PrintProducts(products []product) {
	for _, p := range products {
		fmt.Println(p.id, p.product, p.expiration_date)
	}
}

func main() {
	dataSourceName := "test.db"
	InitDB(dataSourceName)

	id, err := AddProduct(6, "Water", "2022-10-12")
	if err != nil {
		panic(err)
	}
	fmt.Println("Added new Product: ")
	PrintProducts([]product{SelectProductById(id)})
	fmt.Println("==========")
	PrintProducts(SelectAllProducts())
	count, err := UpdateProductById(id, "Orange", "2021-12-30")
	if err != nil {
		panic(err)
	}
	fmt.Print("Updated ", count, " rows\n")
	PrintProducts(SelectAllProducts())

}
