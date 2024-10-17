package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var database *sql.DB

// DsnString returns a string that is used to connect to the database
func DsnString(host, user, pass, name string, port int, ssl, timezone string) string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s",
		host, user, pass, name, port, ssl, timezone)
	return dsn
}

// NewDbConn create and returns a new database connection.
//
// Example:
//
//	db, err := database.NewDbConn(dsnString)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer db.Close()
func NewDbConn(dsnString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsnString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	database = db
	return db, err
}

func GetDB() *sql.DB {
	return database
}
