package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Инициализация подключения к базе данных
	InitDB()

	// Настройка маршрутизации
	http.HandleFunc("/search", SearchCarsHandler)

	// Запуск HTTP-сервера. Используем переменную окружения PORT, если она установлена.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Порт по умолчанию
	}

	log.Printf("Сервер запущен на :%s", port)
	// Fatalf вызовет os.Exit(1), если сервер не сможет запуститься
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
