package main

import (
	"fmt"
	"math/rand"
	"parallelism-analyzer/internal/algorithms/fourier"
	"parallelism-analyzer/pkg/analyzer"
	"sort"
)

func main() {
	// s := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// S, err := fourier.DirectTransformParallel(&s, 12)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Прямое преобразование:", S)

	// 	sInv, err := fourier.InverseTransformParallel(&S, 12)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("Обратное преобразование:", sInv)
	// 	}
	// }

	test := make([]float64, 16384)
	for i := 0; i < len(test); i++ {
		test[i] = rand.Float64() * float64(rand.Intn(1000000))
	}

	anl, err := analyzer.Analyze(
		func(threadCount int) error {
			_, err := fourier.DirectTransformParallel(&test, threadCount)
			return err
		},
		[]int{1, 2, 4, 8, 16},
	)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		keys := make([]int, 0, len(anl))
		for k := range anl {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, threads := range keys {
			dur := anl[threads]
			fmt.Printf("Длительность (%v): %v\n", threads, dur.String())
		}
	}
}
