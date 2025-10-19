package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует случайные элементы.
func generateRandomElements(size int) []int {
	// Обработка крайнего случая: размер слайса равен 0
	if size <= 0 {
		return []int{}
	}

	// Создаем слайс нужного размера
	data := make([]int, size)

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Заполняем слайс случайными положительными числами
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) + 1 // числа от 1 до 1000000
	}

	return data
}

// maximum возвращает максимальное число элементов.
func maximum(data []int) int {
	// Обработка крайнего случая: пустой слайс
	if len(data) == 0 {
		return 0
	}

	// Обработка крайнего случая: слайс с одним элементом
	if len(data) == 1 {
		return data[0]
	}

	// Инициализируем максимум первым элементом
	maxValue := data[0]

	// Проходим по всем элементам слайса
	for i := 1; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}

	return maxValue
}

// maxChunks возвращает максимальное число элементов в чанках.
func maxChunks(data []int) int {
	// Обработка крайнего случая: пустой слайс
	if len(data) == 0 {
		return 0
	}

	// Обработка крайнего случая: слайс с одним элементом
	if len(data) == 1 {
		return data[0]
	}

	// Создаем слайс для хранения максимумов каждого чанка
	chunkMaxes := make([]int, CHUNKS)

	// Создаем WaitGroup для синхронизации горутин
	var wg sync.WaitGroup

	// Вычисляем размер каждого чанка
	chunkSize := len(data) / CHUNKS

	// Запускаем горутины для обработки каждого чанка
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func(chunkIndex int) {
			defer wg.Done()

			// Вычисляем индексы для текущего чанка
			startIndex := chunkIndex * chunkSize
			endIndex := startIndex + chunkSize

			// Обработка последнего чанка - он может быть больше остальных
			if chunkIndex == CHUNKS-1 {
				endIndex = len(data)
			}

			// Если чанк пустой, пропускаем его
			if startIndex >= endIndex {
				chunkMaxes[chunkIndex] = 0
				return
			}

			// Находим максимум в текущем чанке
			maxValue := data[startIndex]
			for j := startIndex + 1; j < endIndex; j++ {
				if data[j] > maxValue {
					maxValue = data[j]
				}
			}

			chunkMaxes[chunkIndex] = maxValue
		}(i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Находим максимум среди максимумов чанков
	return maximum(chunkMaxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	// Генерируем данные
	data := generateRandomElements(SIZE)
	fmt.Printf("Сгенерировано %d элементов\n", len(data))

	fmt.Println("Ищем максимальное значение в один поток")

	// Измеряем время выполнения однопоточного поиска
	startTime := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startTime)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d микросекунд\n", maxSingle, elapsedSingle.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	// Измеряем время выполнения многопоточного поиска
	startTime = time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(startTime)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d микросекунд\n", maxMulti, elapsedMulti.Microseconds())

	// Проверяем, что результаты совпадают
	if maxSingle == maxMulti {
		fmt.Println("Результаты совпадают!")
	} else {
		fmt.Printf("ОШИБКА: Результаты не совпадают! Однопоточный: %d, Многопоточный: %d\n", maxSingle, maxMulti)
	}

	// Выводим информацию об ускорении
	if elapsedMulti.Microseconds() > 0 {
		speedup := float64(elapsedSingle.Microseconds()) / float64(elapsedMulti.Microseconds())
		fmt.Printf("Ускорение: %.2fx\n", speedup)
	}
}
