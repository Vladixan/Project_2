package main

// Пишите тесты в этом файле

import (
	"testing"
)

// Тест для функции генерации случайных чисел
func TestGenerateRandomElements(t *testing.T) {
	// Таблица тестовых случаев
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{
			name:    "положительный размер",
			size:    100,
			wantLen: 100,
		},
		{
			name:    "нулевой размер",
			size:    0,
			wantLen: 0,
		},
		{
			name:    "отрицательный размер",
			size:    -10,
			wantLen: 0,
		},
		{
			name:    "размер 1",
			size:    1,
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			// Проверяем, что длина слайса соответствует ожидаемой
			if len(got) != tt.wantLen {
				t.Errorf("generateRandomElements(%d) вернул слайс длиной %d, ожидалось %d",
					tt.size, len(got), tt.wantLen)
			}

			// Если размер положительный, проверяем, что слайс не пустой и содержит данные
			if tt.size > 0 && len(got) > 0 && got[0] == 0 && got[len(got)-1] == 0 {
				// Очень маловероятно, что rand.Int() сгенерирует два нуля подряд
				t.Errorf("generateRandomElements(%d) вернул слайс нулей, ожидались случайные числа", tt.size)
			}
		})
	}
}

// Тест для функции поиска максимума в одном потоке
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "обычный массив",
			data: []int{1, 5, 3, 8, 2},
			want: 8,
		},
		{
			name: "максимум в начале",
			data: []int{10, 2, 3, 4},
			want: 10,
		},
		{
			name: "максимум в конце",
			data: []int{1, 2, 3, 10},
			want: 10,
		},
		{
			name: "все элементы одинаковые",
			data: []int{5, 5, 5, 5},
			want: 5,
		},
		{
			name: "один элемент",
			data: []int{7},
			want: 7,
		},
		{
			name: "пустой массив",
			data: []int{},
			want: 0,
		},
		{
			name: "отрицательные числа",
			data: []int{-10, -5, -20, -1},
			want: -1,
		},
		{
			name: "отрицательные числа кроме первого",
			data: []int{10, -5, -20, -1},
			want: 10,
		},
		{
			name: "отрицательные числа кроме последнего",
			data: []int{-10, -5, -20, 1},
			want: 1,
		},
		{
			name: "отрицательные числа кроме элемента в середине",
			data: []int{-10, -5, 20, -1},
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) вернул %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

// Тест для функции многопоточного поиска максимума
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "обычный большой массив (кратный 8), максимум в конце",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100},
			want: 100,
		},
		{
			name: "обычный большой массив (кратный 8), максимум в начале",
			data: []int{200, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100},
			want: 200,
		},
		{
			name: "обычный большой массив (кратный 8), максимум в середине первого среза",
			data: []int{1, 2, 3, 1000, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100},
			want: 1000,
		},
		{
			name: "обычный большой массив (кратный 8), максимум в середине второго среза",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000, 11, 12, 13, 14, 15, 100},
			want: 1000,
		},
		{
			name: "массив не кратный 8 (проверка остатка)",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 99, 100, 50},
			want: 100, // 19 элементов
		},
		{
			name: "маленький массив (меньше 8)",
			data: []int{4, 8, 2, 1},
			want: 8,
		},
		{
			name: "пустой массив",
			data: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) вернул %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
