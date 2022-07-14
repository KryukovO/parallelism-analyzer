package dithering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDitheringMatrix(t *testing.T) {
	table := []struct {
		name    string
		arg     int
		result  [][]int
		isError bool
	}{
		{
			name:    "dithering.ditheringMatrix(): Wrong order №1",
			arg:     0,
			result:  nil,
			isError: true,
		},
		{
			name:    "dithering.ditheringMatrix(): Wrong order №2",
			arg:     6,
			result:  nil,
			isError: true,
		},
		{
			name:    "dithering.ditheringMatrix(): Check result",
			arg:     4,
			result:  [][]int{{0, 8, 2, 10}, {12, 4, 14, 6}, {3, 11, 1, 9}, {15, 7, 13, 5}},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			D, err := ditheringMatrix(test.arg)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			assert.EqualValues(t, test.result, D)
		})
	}
}

func TestOrderedDithering(t *testing.T) {
	type args struct {
		srcImgPath string
		dstImgPath string
		order      int
	}

	table := []struct {
		name    string
		args    args
		isError bool
	}{
		{
			name:    "dithering.OrderedDithering(): Wrong order №1",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 0},
			isError: true,
		},
		{
			name:    "dithering.OrderedDithering(): Wrong order №2",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 6},
			isError: true,
		},
		{
			name:    "dithering.OrderedDithering(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 4},
			isError: true,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := OrderedDithering(test.args.srcImgPath, test.args.dstImgPath, test.args.order)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
		})
	}
}

func TestOrderedDitheringParallel(t *testing.T) {
	type args struct {
		srcImgPath  string
		dstImgPath  string
		order       int
		threadCount int
	}

	table := []struct {
		name    string
		args    args
		isError bool
	}{
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong order №1",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 0, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong order №2",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 6, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong thread count",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 4, threadCount: 0},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 4, threadCount: 2},
			isError: true,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := OrderedDitheringParallel(test.args.srcImgPath, test.args.dstImgPath, test.args.order, test.args.threadCount)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
		})
	}
}
