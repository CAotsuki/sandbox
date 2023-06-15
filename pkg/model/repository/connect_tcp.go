package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func connctTCPSocket() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error environment variable %s not set", k)
		}
		return v
	}

	var (
		dbHost = mustGetenv("DB_HOST")
		dbPort = mustGetenv("DB_PORT")
		dbUser = mustGetenv("DB_USER")
		dbPass = mustGetenv("DB_PASS")
		dbName = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		dbHost, dbPort, dbUser, dbPass, dbName)

	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	configureConnectionPool(dbPool)

	return dbPool, nil
}
