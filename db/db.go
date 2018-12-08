package db

import (
	"database/sql"
    _"github.com/lib/pq"
	)

var (
	Conn *sql.DB
)

func Connect() error {
	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=postgres sslmode=disable")
	if err != nil {
		return err
	}
	Conn = db
	err = Conn.Ping()
	if err != nil {
		return err
	}
	return nil
}
