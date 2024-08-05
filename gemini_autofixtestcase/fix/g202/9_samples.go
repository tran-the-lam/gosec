// ExecContext match
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	result, err := db.ExecContext(context.Background(), "select * from foo where name ?", os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
