package repository

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		db = mustConnect()
	})
	return db
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error cannot load .env: %s", err)
	}
}

func migrateDB(db *sql.DB) error {
	createTodos := `CREATE TABLE IF NOT EXISTS todo (
		id SERIAL NOT NULL,
		title VARCHAR(40) NOT NULL,
		content VARCHAR(100) NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`
	_, err := db.Exec(createTodos)
	return err
}

func mustConnect() *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	db, err = connctTCPSocket()
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	if err := migrateDB(db); err != nil {
		log.Fatalf("unable to create table: %s", err)
	}

	return db
}

// configureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
func configureConnectionPool(db *sql.DB) {
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(7)
	db.SetConnMaxLifetime(1800 * time.Second)
}
