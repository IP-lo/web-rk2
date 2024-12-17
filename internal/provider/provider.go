package provider

import (
	"database/sql"
	"log"
)

type Provider struct {
	conn *sql.DB
}

func InitDB() *sql.DB {
	connStr := "user=postgres password=yourpassword dbname=habit_tracker sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database connection error:", err)
	}
	return db
}
