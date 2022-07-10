package main

import (
	"fmt"
	"math/rand"
	"parallelism-analyzer/internal/algorithms/fourier"
	"parallelism-analyzer/internal/analyzer"
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

	for threads := 1; threads <= 16; threads *= 2 {
		dur, err := analyzer.Analyze(
			func(threadCount int) error {
				_, err := fourier.DirectTransformParallel(&test, threadCount)
				return err
			},
			threads,
		)

		if err != nil {
			fmt.Printf("Ошибка (%v): %v\n", threads, err)
		} else {
			fmt.Printf("Длительность (%v): %v\n", threads, dur.String())
		}
	}

	testCompl := make([]complex128, 16384)
	for i := 0; i < len(testCompl); i++ {
		testCompl[i] = complex(rand.Float64()*float64(rand.Intn(1000000)), 0)
	}

	for threads := 1; threads <= 16; threads *= 2 {
		dur, err := analyzer.Analyze(
			func(threadCount int) error {
				_, err := fourier.InverseTransformParallel(&testCompl, threadCount)
				return err
			},
			threads,
		)

		if err != nil {
			fmt.Printf("Ошибка (%v): %v\n", threads, err)
		} else {
			fmt.Printf("Длительность (%v): %v\n", threads, dur.String())
		}
	}
}
