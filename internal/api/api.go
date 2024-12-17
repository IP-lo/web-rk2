package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	habitHandler := NewHabitHandler(db)
	logHandler := NewLogHandler(db)

	e.POST("/habits", habitHandler.CreateHabit)
	e.GET("/habits", habitHandler.GetHabits)
	e.GET("/habits/:id", habitHandler.GetHabitByID)
	e.PUT("/habits/:id", habitHandler.UpdateHabit)
	e.DELETE("/habits/:id", habitHandler.DeleteHabit)

	e.POST("/habits/:id/logs", logHandler.AddLog)
	e.GET("/habits/:id/logs", logHandler.GetLogs)
}
