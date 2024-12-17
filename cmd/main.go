package main

import (
	"github.com/IP-lo/web-rk2/internal/api"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// Инициализация базы данных
	db := database.InitDB()

	// Инициализация Echo
	e := echo.New()

	// Регистрация маршрутов
	routes.RegisterRoutes(e, db)

	// Запуск сервера
	log.Fatal(e.Start(":8080"))
}

/* разработать трекер привычек - календарь куда записываются привычки и отмечаются
дни выполнения */
