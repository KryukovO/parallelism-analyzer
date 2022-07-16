package analyzer

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/xuri/excelize/v2"
)

// Тип, хранящий результаты анализа длительности и предоставляющий методы их обработки
type Analyzer struct {
	name     string
	duration map[int][]time.Duration
}

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
func Analyze(f func(int) error, name string, threadCounts []int, testCount int) (anl Analyzer, errF error) {
	defer func() { // Функция отлова непредвиденной паники
		if msg := recover(); msg != nil {
			anl = Analyzer{}
			errF = fmt.Errorf("%v", msg)
		}
	}()

	if name == "" {
		name = "Временные характеристики"
	}

	anl = Analyzer{name: name, duration: make(map[int][]time.Duration)} // Аллоцируем структуру, хранящую результаты
	for _, threadCount := range threadCounts {                          // Выполняем функцию analyzeOne для заданной функции для каждого числа потоков из множества
		anl.duration[threadCount] = make([]time.Duration, testCount)
		log.Printf("Запуск тестов для %v потока(-ов)", threadCount)
		for testNum := 0; testNum < testCount; testNum++ {
			log.Printf("Запуск теста №%v/%v", testNum+1, testCount)
			if duration, err := analyzeOne(f, threadCount); err != nil {
				return Analyzer{}, err
			} else {
				anl.duration[threadCount][testNum] = duration
			}
		}
	}

	return anl, nil
}

// Функция выгрузки данных в MS Excel
func (anl *Analyzer) SaveXLSX(dstFilePath string) (errF error) {
	defer func() { // Функция отлова непредвиденной паники
		if msg := recover(); msg != nil {
			errF = fmt.Errorf("%v", msg)
		}
	}()

	threads := make([]int, 0, len(anl.duration)) // Слайс для упорядочивания мапы
	for threadCount := range anl.duration {
		threads = append(threads, threadCount)
	}
	sort.Ints(threads)

	// Генерируем xlsx документ и определяем стили в нем
	xlsx := excelize.NewFile()
	sheetId := xlsx.NewSheet(anl.name)
	xlsx.SetActiveSheet(sheetId)
	xlsx.SetCellValue(anl.name, "A1", "№ теста")
	style, err := xlsx.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		return err
	}

	// Заполняем документ данными
	threadColumn := 'B'
	testCount := len(anl.duration[threads[0]])
	for testNum := 1; testNum <= testCount; testNum++ {
		xlsx.SetCellValue(anl.name, fmt.Sprintf("A%v", testNum+1), fmt.Sprintf("%v", testNum))
	}
	for _, threadCount := range threads {
		durations := anl.duration[threadCount]
		xlsx.SetCellValue(anl.name, fmt.Sprintf("%v1", string(threadColumn)), fmt.Sprintf("%v поток(-а), мкс", threadCount))
		for testNum, duration := range durations {
			xlsx.SetCellValue(anl.name, fmt.Sprintf("%v%v", string(threadColumn), testNum+2), duration.Microseconds())
		}
		threadColumn = rune(threadColumn + 1)
	}
	xlsx.SetColWidth(anl.name, "A", string(threadColumn-1), 20)
	xlsx.SetCellStyle(anl.name, "A1", fmt.Sprintf("%v%v", string(threadColumn-1), testCount+1), style)
	xlsx.DeleteSheet("Sheet1")

	// Добавляем в документ диаграмму-график
	// TO DO: https://xuri.me/excelize/en/chart/line.html

	// Формируем файл
	if err := xlsx.SaveAs(dstFilePath); err != nil {
		return err
	}

	return nil
}
