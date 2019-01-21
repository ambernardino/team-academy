package main

import (
"database/sql"
"fmt"
"strconv"

_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "./clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec()
	statement, err = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec("John", "Doe")
	rows, err := database.Query("SELECT id, firstname, lastname FROM people")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var id int
	var firstname, lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}

