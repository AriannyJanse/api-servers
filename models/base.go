package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	username := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", username, dbHost, dbPort, dbName)
	fmt.Println(dbURI)

	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		fmt.Print("Error connecting to the database: ", err)
	}

	db = conn
	fmt.Println("Connected to the database...")
}

func GetDB() *sql.DB {
	return db
}
