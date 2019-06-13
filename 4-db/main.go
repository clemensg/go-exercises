package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "CREATE TABLE IF NOT EXISTS test (id INTEGER NOT NULL PRIMARY KEY, foo TEXT)"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("SQL statement %s failed: %q", sqlStmt, err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO TEST (id, foo) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 3; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("Random Content %d", rand.Int()))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
