// Format string without proper quoting with context
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
	// q := fmt.Sprintf("select * from foo where name = ''", os.Args[1])
	rows, err := db.QueryContext(context.Background(), "select * from foo where name = ?", os.Args[1])
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
