package fourier

import (
	"errors"
	"fmt"
	"math"
)

// Cтруктура для передачи результата рассчета прямого ДПФ по каналу
type directRetValues struct {
	result *[]complex128
	err    *error
}

// Cтруктура для передачи результата рассчета обратного ДПФ по каналу
type inverseRetValues struct {
	result *[]float64
	err    *error
}

// Оболочка для параллельного выполнения рассчета прямого ДПФ
func DirectTransformParallel(values *[]float64, threadCount int) (result []complex128, errF error) {
	defer func() { // Функция отлова непредвиденной паники
		if msg := recover(); msg != nil {
			result = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	if threadCount == 0 {
		return nil, errors.New("число потоков, должно быть больше 0")
	}

	N := len(*values) // Число элементов в наборе значений
	if N == 0 {
		return nil, errors.New("набор значений пуст")
	}

	result = make([]complex128, 0, N) // Инициализация слайса с результатом работы функции

	step := int(math.Max(math.Ceil(float64(N)/float64(threadCount)), 1)) // Минимальный интревал, обрабатываемый горутиной - 1

	channels := make([]chan directRetValues, 0, threadCount) // Слайс каналов для получения результатов рассчета от горутин
	defer func() {                                           // Отложенная функция освобождения каналов (чтобы горутины могли завершиться, если все еще заблокированы записью)
		for _, channel := range channels {
			<-channel
		}
	}()

	for i := 0; i < threadCount; i++ {
		left := i * step // Левая граница рассчета для горутины
		if left < N {    // Рассчет выполняется, только если хватает значений, иначе горутина не нужна
			channel := make(chan directRetValues)                  // Создаем канал взаимодействия
			channels = append(channels, channel)                   // Помещаем канал в слайс
			right := int(math.Min(float64(left+step), float64(N))) // Вычисляем правую границу (она не может быть больше числа значений)
			go func() {                                            // Запуск горутины
				S, err := DirectTransform(values, left, right) // Рассчет прямого ДПФ для заданных границ
				channel <- directRetValues{&S, &err}           // Передача результатов в канал взаимодействия
				close(channel)                                 // Закрытие канала
			}()
		}
	}

	for _, channel := range channels {
		goResult := <-channel     // Читаем содержимое канала (канал в любом случае будет открыт, так что проверка не нужна)
		if *goResult.err != nil { // Если была получена ошибка, то завершаем работу (каналы будут освобождены в отложенной функции)
			return nil, *goResult.err
		}
		result = append(result, *goResult.result...) // Добавляем результаты прямого ДПФ в результируйщий слайт
	}

	return result, nil
}

// Оболочка для параллельного выполнения рассчета обратного ДПФ
func InverseTransformParallel(values *[]complex128, threadCount int) (result []float64, errF error) {
	defer func() { // Функция отлова непредвиденной паники
		if msg := recover(); msg != nil {
			result = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	if threadCount == 0 {
		return nil, errors.New("число потоков, должно быть больше 0")
	}

	N := len(*values) // Число элементов в наборе значений
	if N == 0 {
		return nil, errors.New("набор значений пуст")
	}

	result = make([]float64, 0, N) // Инициализация слайса с результатом работы функции

	step := int(math.Max(math.Ceil(float64(N)/float64(threadCount)), 1)) // Минимальный интревал, обрабатываемый горутиной - 1

	channels := make([]chan inverseRetValues, 0, threadCount) // Слайс каналов для получения результатов рассчета от горутин
	defer func() {                                            // Отложенная функция освобождения каналов (чтобы горутины могли завершиться при возникновении ошибок)
		for _, channel := range channels {
			<-channel
		}
	}()

	for i := 0; i < threadCount; i++ {

		left := i * step // Левая граница рассчета для горутины
		if left < N {    // Рассчет выполняется, только если хватает значений, иначе горутина не нужна
			channel := make(chan inverseRetValues)                 // Создаем канал взаимодействия
			channels = append(channels, channel)                   // Помещаем канал в слайс
			right := int(math.Min(float64(left+step), float64(N))) // Вычисляем правую границу (она не может быть больше числа значений)
			go func() {                                            // Запуск горутины
				s, err := InverseTransform(values, left, right) // Ррассчет обратного ДПФ для заданных границ
				channel <- inverseRetValues{&s, &err}           // Передача результатов в канал взаимодействия
				close(channel)                                  // Закрытие канала
			}()
		}
	}

	for _, channel := range channels {
		goResult := <-channel     // Читаем содержимое канала (канал в любом случае будет открыт, так что проверка не нужна)
		if *goResult.err != nil { // Если была получена ошибка, то завершаем работу (каналы будут освобождены в отложенной функции)
			return nil, *goResult.err
		}
		result = append(result, *goResult.result...) // Добавляем результаты обратного ДПФ в результируйщий слайт
	}

	return result, nil
}
