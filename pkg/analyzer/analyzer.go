package analyzer

import (
	"fmt"
	"time"
)

// Тип, хранящий результаты анализа длительности ([число потоков] = Duration)
type Analyzer = map[int]time.Duration

// Функция однократного анализа длительности выполнения функции для единичного числа потоков
func analyzeOne(f func(int) error, threadCount int) (duration time.Duration, errF error) {
	defer func() { // Функция отлова непредвиденной паники
		if msg := recover(); msg != nil {
			duration = 0
			errF = fmt.Errorf("%v", msg)
		}
	}()

	start := time.Now()   // Засекаем время начала выполнения
	err := f(threadCount) // Дергаем функцию
	now := time.Now()     // Засекаем время завершения выполнения
	if err != nil {       // Что-то пошло не так
		return 0, fmt.Errorf("ошибка при %v потоков: %v", threadCount, err)
	}
	duration = now.Sub(start) // Вычисляем длительность выполнения функции

	return duration, nil
}

// Функция многократного анализа длительности выполнения функции на заданном множестве потоков
func Analyze(f func(int) error, threadCounts []int) (anl Analyzer, errF error) {
	defer func() { // Функция отлова непредвиденной паники
		if msg := recover(); msg != nil {
			anl = Analyzer{}
			errF = fmt.Errorf("%v", msg)
		}
	}()

	anl = make(map[int]time.Duration) // Аллоцируем возвращаемое значение как мапу
	for _, n := range threadCounts {  // Выполняем функцию analyzeOne для заданной функции для каждого числа потоков из множества
		if duration, err := analyzeOne(f, n); err != nil {
			return Analyzer{}, err
		} else {
			anl[n] = duration
		}
	}

	return anl, nil
}
