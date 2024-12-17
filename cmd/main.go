package main

import (
	"log"

	"github.com/IP-lo/web-rk2/internal/api"
	"github.com/IP-lo/web-rk2/internal/provider"
	"github.com/labstack/echo/v4"
)

func main() {
	// Инициализация базы данных
	db := provider.InitDB()

	// Инициализация Echo
	e := echo.New()

	// Регистрация маршрутов
	api.RegisterRoutes(e, db)

	// Запуск сервера
	log.Fatal(e.Start(":8080"))
}

/* разработать трекер привычек - календарь куда записываются привычки и отмечаются
дни выполнения */
