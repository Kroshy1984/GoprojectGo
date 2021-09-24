package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type student struct {
	id        int
	name      string
	course    int
	group_num int
}

func insertIntoDB(n string, c int, g int) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Student (name, course, group_num) values ($1, $2, $3)", n, c, g)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}

func checkConnection(dbPATH string) {
	db, err := sql.Open("sqlite3", dbPATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func selectFromDB(dbPATH string) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Student")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	students := []student{}

	for rows.Next() {
		st := student{}
		err := rows.Scan(&st.id, &st.name, &st.course, &st.group_num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		students = append(students, st)
	}
	for _, st := range students {
		fmt.Println(st.id, st.name, st.course, st.group_num)
	}
}

func main() {

	checkConnection("test.db")

	//insertIntoDB("Name3", 4, 7)

	selectFromDB("test.db")
}
