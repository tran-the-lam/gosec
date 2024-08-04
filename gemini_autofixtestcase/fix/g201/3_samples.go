// Format string without proper quoting with transaction
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
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	// q := fmt.Sprintf("select * from foo where name = '%s'", os.Args[1])
	rows, err := tx.QueryContext(context.Background(), "select * from foo where name = ?", os.Args[1])
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
