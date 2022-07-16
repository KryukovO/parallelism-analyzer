package main

import (
	"fmt"
	"math/rand"
	"parallelism-analyzer/internal/algorithms/fourier"
	"parallelism-analyzer/pkg/analyzer"
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
		"Прямое ДПФ",
		[]int{1, 2, 4, 6},
		20,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = anl.SaveXLSX("results/analyze.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// dithering.OrderedDithering("assets/tiger.png", "results/tiger_ord_dith_2.png", 2)
	// dithering.OrderedDithering("assets/tiger.png", "results/tiger_ord_dith_4.png", 4)
	// dithering.OrderedDithering("assets/tiger.png", "results/tiger_ord_dith_8.png", 8)
	// dithering.OrderedDithering("assets/tiger.png", "results/tiger_ord_dith_16.png", 16)
	// dithering.OrderedDithering("assets/tiger.png", "results/tiger_ord_dith_32.png", 32)

	// dithering.OrderedDitheringParallel("assets/tiger.png", "results/tiger_ord_dith_p_2.png", 2, 6)
	// dithering.OrderedDitheringParallel("assets/tiger.png", "results/tiger_ord_dith_p_4.png", 4, 6)
	// dithering.OrderedDitheringParallel("assets/tiger.png", "results/tiger_ord_dith_p_8.png", 8, 6)
	// dithering.OrderedDitheringParallel("assets/tiger.png", "results/tiger_ord_dith_p_16.png", 16, 6)
	// dithering.OrderedDitheringParallel("assets/tiger.png", "results/tiger_ord_dith_p_32.png", 32, 6)

	// dithering.ThresholdDithering("assets/tiger.png", "results/tiger_trhld_dith_50.png", 50)
	// dithering.ThresholdDithering("assets/tiger.png", "results/tiger_trhld_dith_100.png", 100)
	// dithering.ThresholdDithering("assets/tiger.png", "results/tiger_trhld_dith_150.png", 150)
	// dithering.ThresholdDithering("assets/tiger.png", "results/tiger_trhld_dith_200.png", 200)

	// dithering.FloydErrDithering("assets/tiger.png", "results/tiger_floyd_dith_50.png", 50)
	// dithering.FloydErrDithering("assets/tiger.png", "results/tiger_floyd_dith_100.png", 100)
	// dithering.FloydErrDithering("assets/tiger.png", "results/tiger_floyd_dith_150.png", 150)
	// dithering.FloydErrDithering("assets/tiger.png", "results/tiger_floyd_dith_200.png", 200)
}
