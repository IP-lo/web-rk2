package api

import (
	"github.com/ValeryBMSTU/web-rk2/internal/entities"
	"time"
)

type Usecase interface {
	CreateUser(entities.User) (*entities.User, error)
	ListUsers() ([]*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	UpdateUserByID(id int, user entities.User) (*entities.User, error)
	DeleteUserByID(id int) error
}
type Habit struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
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
