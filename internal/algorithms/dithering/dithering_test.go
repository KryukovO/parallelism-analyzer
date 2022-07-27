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
			name:    "dithering.OrderedDithering(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 4},
			isError: true,
		},
		{
			name:    "dithering.OrderedDithering(): Wrong destination file",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "", order: 4},
			isError: true,
		},
		{
			name:    "dithering.OrderedDithering(): Wrong order №1",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 0},
			isError: true,
		},
		{
			name:    "dithering.OrderedDithering(): Wrong order №2",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 6},
			isError: true,
		},
		{
			name:    "dithering.OrderedDithering(): Check result",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 4},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := OrderedDithering(test.args.srcImgPath, test.args.dstImgPath, test.args.order)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.args.dstImgPath)
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
			name:    "dithering.OrderedDitheringParallel(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", order: 4, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong destination file",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "", order: 4, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong order №1",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 0, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong order №2",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 6, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Wrong thread count",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 4, threadCount: 0},
			isError: true,
		},
		{
			name:    "dithering.OrderedDitheringParallel(): Check result",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", order: 4, threadCount: 2},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := OrderedDitheringParallel(test.args.srcImgPath, test.args.dstImgPath, test.args.order, test.args.threadCount)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.args.dstImgPath)
			}
		})
	}
}

func TestThresholdDithering(t *testing.T) {
	type args struct {
		srcImgPath string
		dstImgPath string
		threshold  int
	}

	table := []struct {
		name    string
		args    args
		isError bool
	}{
		{
			name:    "dithering.ThresholdDithering(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", threshold: 100},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDithering(): Wrong destination file",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "", threshold: 100},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDithering(): Wrong threshold №1",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: -100},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDithering(): Wrong threshold №2",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 300},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDithering(): Check result",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 100},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := ThresholdDithering(test.args.srcImgPath, test.args.dstImgPath, test.args.threshold)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.args.dstImgPath)
			}
		})
	}
}

func TestThresholdDitheringParallel(t *testing.T) {
	type args struct {
		srcImgPath  string
		dstImgPath  string
		threshold   int
		threadCount int
	}

	table := []struct {
		name    string
		args    args
		isError bool
	}{
		{
			name:    "dithering.ThresholdDitheringParallel(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", threshold: 100, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDitheringParallel(): Wrong destination file",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "", threshold: 100, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDitheringParallel(): Wrong threshold №1",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: -100, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDitheringParallel(): Wrong threshold №2",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 300, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDitheringParallel(): Wrong thread count",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 100, threadCount: 0},
			isError: true,
		},
		{
			name:    "dithering.ThresholdDitheringParallel(): Check result",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 100, threadCount: 2},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := ThresholdDitheringParallel(test.args.srcImgPath, test.args.dstImgPath, test.args.threshold, test.args.threadCount)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.args.dstImgPath)
			}
		})
	}
}

func TestFloydErrDithering(t *testing.T) {
	type args struct {
		srcImgPath string
		dstImgPath string
		threshold  int
	}

	table := []struct {
		name    string
		args    args
		isError bool
	}{
		{
			name:    "dithering.FloydErrDithering(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", threshold: 100},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDithering(): Wrong destination file",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "", threshold: 100},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDithering(): Wrong threshold №1",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: -100},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDithering(): Wrong threshold №2",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 300},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDithering(): Check result",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 100},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := FloydErrDithering(test.args.srcImgPath, test.args.dstImgPath, test.args.threshold)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.args.dstImgPath)
			}
		})
	}
}

func TestFloydErrDitheringParallel(t *testing.T) {
	type args struct {
		srcImgPath  string
		dstImgPath  string
		threshold   int
		threadCount int
	}

	table := []struct {
		name    string
		args    args
		isError bool
	}{
		{
			name:    "dithering.FloydErrDitheringParallel(): Source file does not exists",
			args:    args{srcImgPath: "source_file.png", dstImgPath: "", threshold: 100, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDitheringParallel(): Wrong destination file",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "", threshold: 100, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDitheringParallel(): Wrong threshold №1",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: -100, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDitheringParallel(): Wrong threshold №2",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 300, threadCount: 2},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDitheringParallel(): Wrong thread count",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 100, threadCount: 0},
			isError: true,
		},
		{
			name:    "dithering.FloydErrDitheringParallel(): Check result",
			args:    args{srcImgPath: "../../../assets/tiger.png", dstImgPath: "../../../results/test_result.png", threshold: 100, threadCount: 2},
			isError: false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			err := FloydErrDitheringParallel(test.args.srcImgPath, test.args.dstImgPath, test.args.threshold, test.args.threadCount)
			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.args.dstImgPath)
			}
		})
	}
}
