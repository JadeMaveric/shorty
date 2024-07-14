package database

import (
	"database/sql"
	"fmt"
	"log"
)

func OpenSqliteDb() *sql.DB {
	log.Println("Connecting to database...")
	db, err := sql.Open("sqlite3", "file:links.db")
	if err != nil {
		panic(fmt.Errorf("OpenSqliteDb: %v", err))
	}
	return db
}
