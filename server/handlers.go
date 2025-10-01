package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

// Структура для возврата данных об автомобиле
type Car struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Price float64 `json:"price_thousands"`
}

func SearchCarsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	field := r.URL.Query().Get("field")

	// 1. Проверка на обязательные параметры
	if query == "" || field == "" {
		http.Error(w, "Missing query or field", http.StatusBadRequest)
		log.Printf("Error: Missing query or field in request")
		return
	}

	var rows *sql.Rows
	var err error

	// 2. Выбор SQL-запроса в зависимости от поля
	switch field {
	case "brand":
		// Запрос по марке
		rows, err = DB.Query(`SELECT b.name, m.name, c.year, c.price_thousands
							  FROM cars c
							  JOIN models m ON c.model_id = m.id
							  JOIN brands b ON m.brand_id = b.id
							  WHERE b.name ILIKE '%' || $1 || '%'`, query)
	case "model":
		// Запрос по модели
		rows, err = DB.Query(`SELECT b.name, m.name, c.year, c.price_thousands
							  FROM cars c
							  JOIN models m ON c.model_id = m.id
							  JOIN brands b ON m.brand_id = b.id
							  WHERE m.name ILIKE '%' || $1 || '%'`, query)
	default:
		// Неверное поле поиска
		http.Error(w, "Invalid field. Must be 'brand' or 'model'.", http.StatusBadRequest)
		log.Printf("Error: Invalid field '%s' provided", field)
		return
	}

	// 3. Обработка ошибок выполнения запроса к БД
	if err != nil {
		http.Error(w, "Database query failed.", http.StatusInternalServerError)
		log.Printf("Database Query Error: %v", err)
		return
	}
	defer rows.Close()

	var cars []Car
	// 4. Сканирование результатов
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.Brand, &car.Model, &car.Year, &car.Price); err != nil {
			// Ошибка при сканировании строки - логируем и отправляем ошибку 500
			http.Error(w, "Error reading database results.", http.StatusInternalServerError)
			log.Printf("Row Scan Error: %v", err)
			return
		}
		cars = append(cars, car)
	}

	// 5. Проверка ошибок после цикла (например, ошибка сети при извлечении данных)
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating over results.", http.StatusInternalServerError)
		log.Printf("Rows Iteration Error: %v", err)
		return
	}

	// 6. Отправка успешного ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Явно устанавливаем 200 OK
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
