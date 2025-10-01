package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Строка подключения к PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		// Используем Fatalf, чтобы остановить программу, если подключение невозможно
		log.Fatalf("Ошибка при открытии подключения к БД: %v", err)
	}

	if err = DB.Ping(); err != nil {
		// Используем Fatalf, если база данных недоступна
		log.Fatalf("БД недоступна (Ping failed): %v", err)
	}
	log.Println("Подключение к БД успешно")
}
