package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Car struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Price float64 `json:"price_thousands"`
}

func SearchCarsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	field := r.URL.Query().Get("field")

	if query == "" || field == "" {
		http.Error(w, "Missing query or field", http.StatusBadRequest)
		return
	}

	var rows *sql.Rows
	var err error

	switch field {
	case "brand":
		rows, err = DB.Query(`SELECT b.name, m.name, c.year, c.price_thousands
                              FROM cars c
                              JOIN models m ON c.model_id = m.id
                              JOIN brands b ON m.brand_id = b.id
                              WHERE b.name ILIKE '%' || $1 || '%'`, query)
	case "model":
		rows, err = DB.Query(`SELECT b.name, m.name, c.year, c.price_thousands
                              FROM cars c
                              JOIN models m ON c.model_id = m.id
                              JOIN brands b ON m.brand_id = b.id
                              WHERE m.name ILIKE '%' || $1 || '%'`, query)
	default:
		http.Error(w, "Invalid field", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.Brand, &car.Model, &car.Year, &car.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
