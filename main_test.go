package main

import (
	"testing"
)

// Тесты для функции generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "Нормальный размер",
			size:     100,
			expected: 100,
		},
		{
			name:     "Размер равен 0",
			size:     0,
			expected: 0,
		},
		{
			name:     "Отрицательный размер",
			size:     -5,
			expected: 0,
		},
		{
			name:     "Размер равен 1",
			size:     1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = длина %d, ожидалось %d", tt.size, len(result), tt.expected)
			}

			// Проверяем, что все числа положительные (для ненулевого размера)
			if tt.size > 0 {
				for i, val := range result {
					if val <= 0 {
						t.Errorf("Элемент %d равен %d, ожидалось положительное число", i, val)
					}
				}
			}
		})
	}
}

// Тесты для функции maximum
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Пустой слайс",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "Слайс с одним элементом",
			data:     []int{42},
			expected: 42,
		},
		{
			name:     "Слайс с несколькими элементами",
			data:     []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		{
			name:     "Все элементы одинаковые",
			data:     []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "Максимум в начале",
			data:     []int{10, 1, 2, 3},
			expected: 10,
		},
		{
			name:     "Максимум в конце",
			data:     []int{1, 2, 3, 10},
			expected: 10,
		},
		{
			name:     "Отрицательные числа",
			data:     []int{-5, -1, -10, -3},
			expected: -1,
		},
		{
			name:     "Смешанные положительные и отрицательные",
			data:     []int{-5, 10, -1, 3},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, result, tt.expected)
			}
		})
	}
}

// Тесты для функции maxChunks
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Пустой слайс",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "Слайс с одним элементом",
			data:     []int{42},
			expected: 42,
		},
		{
			name:     "Слайс с несколькими элементами",
			data:     []int{1, 5, 3, 9, 2, 7, 4, 6},
			expected: 9,
		},
		{
			name:     "Все элементы одинаковые",
			data:     []int{7, 7, 7, 7, 7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "Большой слайс",
			data:     make([]int, 1000),
			expected: 0, // будет 0, так как все элементы инициализированы нулями
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Для теста с большим слайсом заполняем его данными
			if tt.name == "Большой слайс" {
				for i := range tt.data {
					tt.data[i] = i + 1
				}
				tt.expected = len(tt.data)
			}

			result := maxChunks(tt.data)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, result, tt.expected)
			}
		})
	}
}

// Тест для проверки согласованности результатов между maximum и maxChunks
func TestConsistency(t *testing.T) {
	// Создаем тестовые данные
	testData := []int{1, 5, 3, 9, 2, 7, 4, 6, 8, 10, 12, 11}

	maxSingle := maximum(testData)
	maxMulti := maxChunks(testData)

	if maxSingle != maxMulti {
		t.Errorf("Результаты не совпадают: maximum = %d, maxChunks = %d", maxSingle, maxMulti)
	}
}
