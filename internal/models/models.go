package models

import "time"

type Habit struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type HabitLog struct {
	ID        int       `json:"id"`
	HabitID   int       `json:"habit_id"`
	Date      time.Time `json:"date"`
	Completed bool      `json:"completed"`
}
