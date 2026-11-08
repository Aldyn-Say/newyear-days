package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerSuccess(t *testing.T) {
	// Создаем запрос с корректной датой (31 декабря 2024)
	req := httptest.NewRequest(http.MethodGet, "/days?date=2024-12-31", nil)
	w := httptest.NewRecorder()

	// Вызываем обработчик напрямую
	handler(w, req)

	// 1. Проверяем статус код
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус 200, получили %d", w.Code)
	}

	// 2. Проверяем, что пришел JSON с правильным ответом
	var resp map[string]int
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Ошибка чтения JSON: %v", err)
	}
	if resp["days"] != 1 {
		t.Errorf("Ожидалось 1 день, получили %d", resp["days"])
	}
}

func TestHandlerBadRequest(t *testing.T) {
	// Создаем запрос с НЕКОРРЕКТНОЙ датой (неправильный формат)
	req := httptest.NewRequest(http.MethodGet, "/days?date=invalid-date", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Проверяем, что вернулся статус 400
	if w.Code != http.StatusBadRequest {
		t.Errorf("Ожидался статус 400, получили %d", w.Code)
	}
}
