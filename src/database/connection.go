package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@/todos_db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
