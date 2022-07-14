package dithering

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// TO DO: разобраться с расширениями файлов

var (
	m2 = [][]int{{0, 2}, {3, 1}} // Минимальная матрица размытия
)

// Функция рассчета матриц размытия больших размеров
func ditheringMatrix(order int) (result [][]int, errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	switch {
	case order <= 1:
		return nil, errors.New("размерность матрицы размытия должна быть больше 1")
	case order == 2:
		return m2, nil
	case order%2 != 0:
		return nil, errors.New("порядок матрицы размытия должен быть степенью двойки")
	}

	result = make([][]int, order)
	for i := 0; i < order; i++ {
		result[i] = make([]int, order)
	}

	parentD, err := ditheringMatrix(order / 2)
	if err != nil {
		return nil, err
	}

	for i := 0; i < order; i++ {
		for j := 0; j < order; j++ {
			switch {
			case i < order/2:
				switch {
				case j < order/2:
					result[i][j] = 4 * parentD[i][j]
				case j >= order/2:
					result[i][j] = 4*parentD[i][j-order/2] + 2
				}
			case i >= order/2:
				switch {
				case j < order/2:
					result[i][j] = 4*parentD[i-order/2][j] + 3
				case j >= order/2:
					result[i][j] = 4*parentD[i-order/2][j-order/2] + 1
				}
			}
		}
	}

	return result, nil
}

// Функция, реализующая алгоритм упорядоченного размытия
func OrderedDithering(srcImgPath string, dstImgPath string, order int) (errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			errF = fmt.Errorf("%v", msg)
		}
	}()

	D, err := ditheringMatrix(order)
	if err != nil {
		return err
	}

	srcImgReader, err := os.Open(srcImgPath)
	if err != nil {
		return err
	}
	defer srcImgReader.Close()
	m, _, err := image.Decode(srcImgReader)
	if err != nil {
		return err
	}
	bounds := m.Bounds()

	dstImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i := (y - bounds.Min.Y) % order
			j := (x - bounds.Min.X) % order
			r, g, b, a := m.At(x, y).RGBA()
			intens := int(math.Max(float64(r/256), math.Max(float64(g/256), float64(b/256))))
			if (intens*order*order+1)/256 > D[i][j] {
				dstImg.Set(x, y, color.RGBA{255, 255, 255, uint8(a / 256)})
			} else {
				dstImg.Set(x, y, color.RGBA{0, 0, 0, uint8(a / 256)})
			}
		}
	}

	dstImgFile, err := os.Create(dstImgPath)
	if err != nil {
		return err
	}
	png.Encode(dstImgFile, dstImg)

	return nil
}
