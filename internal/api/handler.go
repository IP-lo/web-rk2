package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type HabitHandler struct {
	db *sql.DB
}

func NewHabitHandler(db *sql.DB) *HabitHandler {
	return &HabitHandler{db: db}
}

func (h *HabitHandler) CreateHabit(c echo.Context) error {
	var habit Habit
	if err := c.Bind(&habit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	query := `INSERT INTO habits (title, description) VALUES ($1, $2) RETURNING id`
	err := h.db.QueryRow(query, habit.Title, habit.Description).Scan(&habit.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create habit"})
	}

	return c.JSON(http.StatusCreated, habit)
}

func (h *HabitHandler) GetHabits(c echo.Context) error {
	rows, err := h.db.Query(`SELECT id, title, description, created_at FROM habits`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch habits"})
	}
	defer rows.Close()

	var habits []Habit
	for rows.Next() {
		var habit Habit
		if err := rows.Scan(&habit.ID, &habit.Title, &habit.Description, &habit.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error parsing habits"})
		}
		habits = append(habits, habit)
	}

	return c.JSON(http.StatusOK, habits)
}

func (h *HabitHandler) GetHabitByID(c echo.Context) error {
	habitID := c.Param("id")

	var habit Habit
	query := `SELECT id, title, description, created_at FROM habits WHERE id = $1`
	err := h.db.QueryRow(query, habitID).Scan(&habit.ID, &habit.Title, &habit.Description, &habit.CreatedAt)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Habit not found"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch habit"})
	}

	return c.JSON(http.StatusOK, habit)
}

func (h *HabitHandler) UpdateHabit(c echo.Context) error {
	habitID := c.Param("id")

	var habit Habit
	if err := c.Bind(&habit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	query := `UPDATE habits SET title = $1, description = $2 WHERE id = $3`
	res, err := h.db.Exec(query, habit.Title, habit.Description, habitID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update habit"})
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Habit not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Habit updated successfully"})
}

func (h *HabitHandler) DeleteHabit(c echo.Context) error {
	habitID := c.Param("id")

	query := `DELETE FROM habits WHERE id = $1`
	res, err := h.db.Exec(query, habitID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete habit"})
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Habit not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Habit deleted successfully"})
}
