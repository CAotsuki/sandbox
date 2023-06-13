package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	Db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
}
