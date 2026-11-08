package countdown

import (
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	// Табличный тест: список дат и ожидаемых результатов
	tests := []struct {
		name string
		date time.Time
		want int
	}{
		{"Конец года", time.Date(2024, time.December, 31, 12, 0, 0, 0, time.UTC), 1},
		{"Начало года (високосный 2024)", time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), 366},
		{"До 29 февраля (високосный)", time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC), 308},
		{"Обычный день", time.Date(2023, time.June, 15, 0, 0, 0, 0, time.UTC), 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DaysUntilNewYear(tt.date)
			if got != tt.want {
				t.Errorf("DaysUntilNewYear(%v) = %d, want %d", tt.date, got, tt.want)
			}
		})
	}
}
