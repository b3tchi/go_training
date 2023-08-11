package db

import (
	"database/sql"
	"log"
	"os"
)

var conn *sql.DB

func InitDB(dbString string) (*sql.DB, error) {
	var err error
	// conn, err = sql.Open("postgres", dbString)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// db, err := sql.Open("postgres", cfg.dsn)
	conn, err = sql.Open("postgres", dbString)
	if err != nil {
		logger.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("database connection pool estabilished")
	return conn, err
}

func GetDB() *sql.DB {
	return conn
}
