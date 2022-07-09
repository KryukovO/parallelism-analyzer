package main

import (
	"fmt"
	"parallelism-analyzer/internal/fourier"
)

func main() {
	s := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	S, err := fourier.DirectTransformParallel(&s, 12)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Прямое преобразование:", S)

		sInv, err := fourier.InverseTransformParallel(&S, 12)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Обратное преобразование:", sInv)
		}
	}
}
