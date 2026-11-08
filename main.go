package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Aldyn-Say/newyear-days/internal/countdown"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Получаем дату из параметра ?date=YYYY-MM-DD
	dateStr := r.URL.Query().Get("date")
	var from time.Time

	if dateStr == "" {
		from = time.Now() // Если даты нет, берем сегодня
	} else {
		var err error
		from, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			// Если дата некорректна, возвращаем 400 Bad Request
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"invalid date format, use YYYY-MM-DD"}`))
			return
		}
	}

	// Вызываем нашу функцию из файла newyear.go
	days := countdown.DaysUntilNewYear(from)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"days": days})
}

func main() {
	http.HandleFunc("/days", handler)
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
