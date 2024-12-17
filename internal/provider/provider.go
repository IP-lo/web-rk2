package provider

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "user=postgres password=yourpassword dbname=habit_tracker sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Создание таблиц, если они не существуют
	createTables(db)

	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS habits (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS habit_logs (
		id SERIAL PRIMARY KEY,
		habit_id INT REFERENCES habits(id) ON DELETE CASCADE,
		date DATE NOT NULL,
		completed BOOLEAN DEFAULT FALSE
	);
	`)
	if err != nil {
		log.Fatal("Ошибка создания таблиц:", err)
	}
}
