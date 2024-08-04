// Format string without proper quoting
package main

import (
	"database/sql"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	q := "SELECT * FROM foo where name = ?"
	rows, err := db.Query(q, os.Args[1])
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
