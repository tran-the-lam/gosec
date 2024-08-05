// context match
package main

import (
	"context"
	"database/sql"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	rows, err := db.QueryContext(context.Background(), "select * from foo where name = ?", os.Args[1])
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
