package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func GetConnection() (db *sql.DB) {
	connStr := "user=postgres password=postgres dbname=lazinessdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
