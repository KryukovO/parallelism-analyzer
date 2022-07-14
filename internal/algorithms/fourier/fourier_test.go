package fourier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectTransform(t *testing.T) {
	type args struct { // Структура для описания аргументов функции DirectTransform
		values []float64
		left   int
		right  int
	}

	table := []struct { // Таблица тестов
		name    string       // Наименование теста
		args    args         // Аргумент функции
		result  []complex128 // Ожидаемый результат
		isError bool         // Должна ли появиться ошибка
	}{
		{
			name:    "fourier.DirectTransform(): Zero length",
			args:    args{values: []float64{}, left: -1, right: -1},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Left border out of range №1",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: -1, right: 10},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Left border out of range №2",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 10, right: 10},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Right border out of range №1",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 0, right: -1},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Right border out of range №2",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 0, right: 11},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Borders out of range",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: -1, right: 11},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Left border is larger than right",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 10, right: 0},
			result:  nil,
			isError: true,
		},
		{
			name: "fourier.DirectTransform(): Check result",
			args: args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 0, right: 10},
			result: []complex128{
				(45 + 0i),
				(-5.0000000000000036 + 15.388417685876266i),
				(-5.000000000000002 + 6.881909602355868i),
				(-5.000000000000003 + 3.6327126400268i),
				(-5.000000000000002 + 1.6245984811645275i),
				(-5 - 5.510910596163082e-15i),
				(-4.999999999999998 - 1.6245984811645382i),
				(-5.000000000000038 - 3.6327126400267824i),
				(-4.999999999999987 - 6.881909602355874i),
				(-4.999999999999973 - 15.388417685876234i),
			},
			isError: false,
		},
		{
			name: "fourier.DirectTransform(): Check result with borders",
			args: args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 3, right: 7},
			result: []complex128{
				(-5.000000000000003 + 3.6327126400268i),
				(-5.000000000000002 + 1.6245984811645275i),
				(-5 - 5.510910596163082e-15i),
				(-4.999999999999998 - 1.6245984811645382i),
			},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			S, err := DirectTransform(&test.args.values, test.args.left, test.args.right)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			assert.EqualValues(t, test.result, S)
		})
	}
}

func TestDirectTransformWrongAmountArgs(t *testing.T) {
	S, err := DirectTransform(&[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 10, 1234567)
	assert.EqualError(t, err, "границы должны быть указаны двумя значениями, получено: 3")
	assert.EqualValues(t, []complex128(nil), S)
}

func TestInverseTransform(t *testing.T) {
	type args struct { // Структура для описания аргументов функции InverseTransform
		values []complex128
		left   int
		right  int
	}

	table := []struct {
		name    string    // Наименование теста
		args    args      // Аргумент функции
		result  []float64 // Ожидаемый результат
		isError bool      // Должна ли появиться ошибка
	}{
		{
			name:    "fourier.DirectTransform(): Zero length",
			args:    args{values: []complex128{}, left: -1, right: -1},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Left border out of range №1",
			args:    args{values: []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: -1, right: 10},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Left border out of range №2",
			args:    args{values: []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 10, right: 10},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Right border out of range №1",
			args:    args{values: []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 0, right: 0},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Right border out of range №2",
			args:    args{values: []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: 0, right: 11},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransform(): Borders out of range",
			args:    args{values: []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, left: -1, right: 11},
			result:  nil,
			isError: true,
		},
		{
			name: "fourier.DirectTransform(): Check result",
			args: args{
				values: []complex128{
					(45 + 0i),
					(-5.0000000000000036 + 15.388417685876266i),
					(-5.000000000000002 + 6.881909602355868i),
					(-5.000000000000003 + 3.6327126400268i),
					(-5.000000000000002 + 1.6245984811645275i),
					(-5 - 5.510910596163082e-15i),
					(-4.999999999999998 - 1.6245984811645382i),
					(-5.000000000000038 - 3.6327126400267824i),
					(-4.999999999999987 - 6.881909602355874i),
					(-4.999999999999973 - 15.388417685876234i),
				},
				left:  0,
				right: 10,
			},
			result: []float64{
				-2.6645352591003756e-16,
				1.0000000000000069,
				2.000000000000005,
				2.9999999999999973,
				4.000000000000003,
				5.000000000000004,
				5.999999999999993,
				6.999999999999995,
				8,
				8.999999999999991,
			},
			isError: false,
		},
		{
			name: "fourier.DirectTransform(): Check result with borders",
			args: args{
				values: []complex128{
					(45 + 0i),
					(-5.0000000000000036 + 15.388417685876266i),
					(-5.000000000000002 + 6.881909602355868i),
					(-5.000000000000003 + 3.6327126400268i),
					(-5.000000000000002 + 1.6245984811645275i),
					(-5 - 5.510910596163082e-15i),
					(-4.999999999999998 - 1.6245984811645382i),
					(-5.000000000000038 - 3.6327126400267824i),
					(-4.999999999999987 - 6.881909602355874i),
					(-4.999999999999973 - 15.388417685876234i),
				},
				left:  3,
				right: 7,
			},
			result: []float64{
				2.9999999999999973,
				4.000000000000003,
				5.000000000000004,
				5.999999999999993,
			},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			sInv, err := InverseTransform(&test.args.values, test.args.left, test.args.right)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			assert.EqualValues(t, test.result, sInv)
		})
	}
}

func TestInverseTransformWrongAmountArgs(t *testing.T) {
	s, err := InverseTransform(&[]complex128{
		(45 + 0i),
		(-5.0000000000000036 + 15.388417685876266i),
		(-5.000000000000002 + 6.881909602355868i),
		(-5.000000000000003 + 3.6327126400268i),
		(-5.000000000000002 + 1.6245984811645275i),
		(-5 - 5.510910596163082e-15i),
		(-4.999999999999998 - 1.6245984811645382i),
		(-5.000000000000038 - 3.6327126400267824i),
		(-4.999999999999987 - 6.881909602355874i),
		(-4.999999999999973 - 15.388417685876234i),
	},
		0,
		10,
		1234567,
	)
	assert.EqualError(t, err, "границы должны быть указаны двумя значениями, получено: 3")
	assert.EqualValues(t, []float64(nil), s)
}

func TestDirectTransformParallel(t *testing.T) {
	type args struct { // Структура для описания аргументов функции DirectTransform
		values      []float64
		threadCount int
	}

	table := []struct { // Таблица тестов
		name    string       // Наименование теста
		args    args         // Аргумент функции
		result  []complex128 // Ожидаемый результат
		isError bool         // Должна ли появиться ошибка
	}{
		{
			name:    "fourier.DirectTransformParallel(): Zero length",
			args:    args{values: []float64{}, threadCount: 2},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.DirectTransformParallel(): Zero threads",
			args:    args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, threadCount: 0},
			result:  nil,
			isError: true,
		},
		{
			name: "fourier.DirectTransformParallel(): Normal relation N/threads",
			args: args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, threadCount: 2},
			result: []complex128{
				(45 + 0i),
				(-5.0000000000000036 + 15.388417685876266i),
				(-5.000000000000002 + 6.881909602355868i),
				(-5.000000000000003 + 3.6327126400268i),
				(-5.000000000000002 + 1.6245984811645275i),
				(-5 - 5.510910596163082e-15i),
				(-4.999999999999998 - 1.6245984811645382i),
				(-5.000000000000038 - 3.6327126400267824i),
				(-4.999999999999987 - 6.881909602355874i),
				(-4.999999999999973 - 15.388417685876234i),
			},
			isError: false,
		},
		{
			name: "fourier.DirectTransformParallel(): Too many threads",
			args: args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, threadCount: 16},
			result: []complex128{
				(45 + 0i),
				(-5.0000000000000036 + 15.388417685876266i),
				(-5.000000000000002 + 6.881909602355868i),
				(-5.000000000000003 + 3.6327126400268i),
				(-5.000000000000002 + 1.6245984811645275i),
				(-5 - 5.510910596163082e-15i),
				(-4.999999999999998 - 1.6245984811645382i),
				(-5.000000000000038 - 3.6327126400267824i),
				(-4.999999999999987 - 6.881909602355874i),
				(-4.999999999999973 - 15.388417685876234i),
			},
			isError: false,
		},
		{
			name: "fourier.DirectTransformParallel(): Irrational relation N/threads",
			args: args{values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, threadCount: 3},
			result: []complex128{
				(45 + 0i),
				(-5.0000000000000036 + 15.388417685876266i),
				(-5.000000000000002 + 6.881909602355868i),
				(-5.000000000000003 + 3.6327126400268i),
				(-5.000000000000002 + 1.6245984811645275i),
				(-5 - 5.510910596163082e-15i),
				(-4.999999999999998 - 1.6245984811645382i),
				(-5.000000000000038 - 3.6327126400267824i),
				(-4.999999999999987 - 6.881909602355874i),
				(-4.999999999999973 - 15.388417685876234i),
			},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			S, err := DirectTransformParallel(&test.args.values, test.args.threadCount)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			assert.EqualValues(t, test.result, S)
		})
	}
}

func TestInverseTransformParallel(t *testing.T) {
	type args struct { // Структура для описания аргументов функции InverseTransform
		values      []complex128
		threadCount int
	}

	table := []struct {
		name    string    // Наименование теста
		args    args      // Аргумент функции
		result  []float64 // Ожидаемый результат
		isError bool      // Должна ли появиться ошибка
	}{
		{
			name:    "fourier.InverseTransformParallel(): Zero length",
			args:    args{values: []complex128{}, threadCount: 2},
			result:  nil,
			isError: true,
		},
		{
			name:    "fourier.InverseTransformParallel(): Zero threads",
			args:    args{values: []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, threadCount: 0},
			result:  nil,
			isError: true,
		},
		{
			name: "fourier.InverseTransformParallel(): Normal relation N/threads",
			args: args{
				values: []complex128{
					(45 + 0i),
					(-5.0000000000000036 + 15.388417685876266i),
					(-5.000000000000002 + 6.881909602355868i),
					(-5.000000000000003 + 3.6327126400268i),
					(-5.000000000000002 + 1.6245984811645275i),
					(-5 - 5.510910596163082e-15i),
					(-4.999999999999998 - 1.6245984811645382i),
					(-5.000000000000038 - 3.6327126400267824i),
					(-4.999999999999987 - 6.881909602355874i),
					(-4.999999999999973 - 15.388417685876234i),
				},
				threadCount: 2,
			},
			result: []float64{
				-2.6645352591003756e-16,
				1.0000000000000069,
				2.000000000000005,
				2.9999999999999973,
				4.000000000000003,
				5.000000000000004,
				5.999999999999993,
				6.999999999999995,
				8,
				8.999999999999991,
			},
			isError: false,
		},
		{
			name: "fourier.InverseTransformParallel(): Too many threads",
			args: args{
				values: []complex128{
					(45 + 0i),
					(-5.0000000000000036 + 15.388417685876266i),
					(-5.000000000000002 + 6.881909602355868i),
					(-5.000000000000003 + 3.6327126400268i),
					(-5.000000000000002 + 1.6245984811645275i),
					(-5 - 5.510910596163082e-15i),
					(-4.999999999999998 - 1.6245984811645382i),
					(-5.000000000000038 - 3.6327126400267824i),
					(-4.999999999999987 - 6.881909602355874i),
					(-4.999999999999973 - 15.388417685876234i),
				},
				threadCount: 16,
			},
			result: []float64{
				-2.6645352591003756e-16,
				1.0000000000000069,
				2.000000000000005,
				2.9999999999999973,
				4.000000000000003,
				5.000000000000004,
				5.999999999999993,
				6.999999999999995,
				8,
				8.999999999999991,
			},
			isError: false,
		},
		{
			name: "fourier.InverseTransformParallel(): Irrational relation N/threads",
			args: args{
				values: []complex128{
					(45 + 0i),
					(-5.0000000000000036 + 15.388417685876266i),
					(-5.000000000000002 + 6.881909602355868i),
					(-5.000000000000003 + 3.6327126400268i),
					(-5.000000000000002 + 1.6245984811645275i),
					(-5 - 5.510910596163082e-15i),
					(-4.999999999999998 - 1.6245984811645382i),
					(-5.000000000000038 - 3.6327126400267824i),
					(-4.999999999999987 - 6.881909602355874i),
					(-4.999999999999973 - 15.388417685876234i),
				},
				threadCount: 3,
			},
			result: []float64{
				-2.6645352591003756e-16,
				1.0000000000000069,
				2.000000000000005,
				2.9999999999999973,
				4.000000000000003,
				5.000000000000004,
				5.999999999999993,
				6.999999999999995,
				8,
				8.999999999999991,
			},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			sInv, err := InverseTransformParallel(&test.args.values, test.args.threadCount)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			assert.EqualValues(t, test.result, sInv)
		})
	}
}
