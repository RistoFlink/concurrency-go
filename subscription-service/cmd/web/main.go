package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
	"time"
)

const webPort = "80"

func main() {
	// connect to the database
	db := initDB()
	db.Ping()
	// create sessions

	// create channels

	// create waitgroup

	// set up the appconfig

	// setup email sending

	// listen for connections
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("unable to connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready for connection")
		} else {
			log.Println("successfully connected to database")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("waiting for database to become available")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
