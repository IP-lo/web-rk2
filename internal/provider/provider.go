package provider

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Provider struct {
	conn *sql.DB
}

func InitDB() *sql.DB {
	connStr := "user=postgres password=BKMZ661248082 dbname=habit_tracker sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS habits (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT
    );`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection error:", err)
	}
	return db
}
