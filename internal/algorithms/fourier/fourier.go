package fourier

import (
	"errors"
	"fmt"
	"math"
)

func DirectTransform(values *[]float64, borders ...int) (result []complex128, errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	N := len(*values)
	if N == 0 {
		return nil, errors.New("набор значений пуст")
	}

	left := 0
	right := N

	if len(borders) != 0 {
		if len(borders) != 2 {
			return nil, fmt.Errorf("границы должны быть указаны двумя значениями, получено: %v", len(borders))
		}
		if borders[0] >= N || borders[0] < 0 || borders[1] > N || borders[1] < 0 || borders[0] > borders[1] {
			return nil, fmt.Errorf("границы за пределами размерности набора значений (%v): от %v до %v", N, borders[0], borders[1]-1)
		}

		left = borders[0]
		right = borders[1]
	}

	result = make([]complex128, right-left)

	for k := left; k < right; k++ {
		for n, val := range *values {
			result[k-left] += complex(val, 0) * complex(math.Cos((2.0*math.Pi*float64(k*n))/float64(N)), -math.Sin((2.0*math.Pi*float64(k*n))/float64(N)))
		}
	}

	return result, nil
}

func InverseTransform(values *[]complex128, borders ...int) (result []float64, errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	N := len(*values)
	if N == 0 {
		return nil, errors.New("набор значений пуст")
	}

	left := 0
	right := N

	if len(borders) != 0 {
		if len(borders) != 2 {
			return nil, fmt.Errorf("границы должны быть указаны двумя значениями, получено: %v", len(borders))
		}
		if borders[0] >= N || borders[0] < 0 || borders[1] > N || borders[1] <= 0 {
			return nil, fmt.Errorf("границы за пределами размерности набора значений (%v): от %v до %v", N, borders[0], borders[1]-1)
		}

		left = borders[0]
		right = borders[1]
	}

	result = make([]float64, right-left)

	for n := left; n < right; n++ {
		for k, value := range *values {
			result[n-left] += real(value * complex(math.Cos(2*math.Pi*float64(n*k)/float64(N)), math.Sin(2*math.Pi*float64(n*k)/float64(N))))
		}
		result[n-left] /= float64(N)
	}

	return result, nil
}
