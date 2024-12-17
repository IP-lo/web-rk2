package api

import (
	"database/sql"
	"net/http"

	"github.com/IP-lo/web-rk2/internal/models"
	"github.com/labstack/echo/v4"
)

type HabitHandler struct {
	db *sql.DB
}

type LogHandler struct {
	db *sql.DB
}

func NewHabitHandler(db *sql.DB) *HabitHandler {
	return &HabitHandler{db: db}
}

func NewLogHandler(db *sql.DB) *LogHandler {
	return &LogHandler{db: db}
}

func (h *HabitHandler) CreateHabit(c echo.Context) error {
	var habit models.Habit
	if err := c.Bind(&habit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	query := `INSERT INTO habits (title, description) VALUES ($1, $2) RETURNING id`
	err := h.db.QueryRow(query, habit.Title, habit.Description).Scan(&habit.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка создания привычки"})
	}

	return c.JSON(http.StatusCreated, habit)
}

func (h *HabitHandler) GetHabits(c echo.Context) error {
	rows, err := h.db.Query(`SELECT id, title, description, created_at FROM habits`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка получения привычек"})
	}
	defer rows.Close()

	var habits []models.Habit
	for rows.Next() {
		var habit models.Habit
		if err := rows.Scan(&habit.ID, &habit.Title, &habit.Description, &habit.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка обработки данных"})
		}
		habits = append(habits, habit)
	}

	return c.JSON(http.StatusOK, habits)
}

func (h *HabitHandler) GetHabitByID(c echo.Context) error {
	habitID := c.Param("id")

	var habit models.Habit
	query := `SELECT id, title, description, created_at FROM habits WHERE id = $1`
	err := h.db.QueryRow(query, habitID).Scan(&habit.ID, &habit.Title, &habit.Description, &habit.CreatedAt)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Привычка не найдена"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка получения привычки"})
	}

	return c.JSON(http.StatusOK, habit)
}

func (h *HabitHandler) UpdateHabit(c echo.Context) error {
	habitID := c.Param("id")
	var habit models.Habit
	if err := c.Bind(&habit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	query := `UPDATE habits SET title = $1, description = $2 WHERE id = $3`
	_, err := h.db.Exec(query, habit.Title, habit.Description, habitID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка обновления привычки"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Привычка успешно обновлена"})
}

func (h *HabitHandler) DeleteHabit(c echo.Context) error {
	habitID := c.Param("id")

	query := `DELETE FROM habits WHERE id = $1`
	_, err := h.db.Exec(query, habitID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка удаления привычки"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Привычка успешно удалена"})
}

func (l *LogHandler) AddLog(c echo.Context) error {
	habitID := c.Param("id") // Получение ID привычки из маршрута

	var log models.HabitLog
	if err := c.Bind(&log); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ошибка обработки данных"})
	}

	// Проверка наличия даты
	if log.Date.IsZero() {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Дата должна быть указана"})
	}

	// SQL-запрос для вставки записи
	query := `INSERT INTO habit_logs (habit_id, date, completed) VALUES ($1, $2, $3)`
	_, err := l.db.Exec(query, habitID, log.Date, log.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка добавления записи"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Лог добавлен успешно"})
}

func (l *LogHandler) GetLogs(c echo.Context) error {
	habitID := c.Param("id") // Получение ID привычки из маршрута

	// SQL-запрос для получения логов
	rows, err := l.db.Query(`SELECT id, date, completed FROM habit_logs WHERE habit_id = $1 ORDER BY date`, habitID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка получения логов"})
	}
	defer rows.Close()

	var logs []models.HabitLog
	for rows.Next() {
		var log models.HabitLog
		if err := rows.Scan(&log.ID, &log.Date, &log.Completed); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка обработки данных"})
		}
		logs = append(logs, log)
	}

	return c.JSON(http.StatusOK, logs)
}
