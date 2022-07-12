package dithering

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
)

// Рекомендуемые матрицы размытия
var (
	m2 = [][]int{{0, 2}, {3, 1}}
	m3 = [][]int{{0, 7, 3}, {6, 5, 2}, {4, 1, 8}}
)

// Функция рассчета матриц размытия больших размеров
func ditheringMatrix(n int) (result [][]int, errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	switch {
	case n <= 1:
		return nil, errors.New("размерность матрицы размытия должна быть больше 1")
	case n == 2:
		return m2, nil
	case n == 3:
		return m3, nil
	}

	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	parentD, err := ditheringMatrix(n / 2)
	if err != nil {
		return nil, err
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i < n/2:
				switch {
				case j < n/2:
					result[i][j] = 4 * parentD[i][j]
				case j >= n/2:
					result[i][j] = 4*parentD[i][j-n/2] + 2
				}
			case i >= n/2:
				switch {
				case j < n/2:
					result[i][j] = 4*parentD[i-n/2][j] + 3
				case j >= n/2:
					result[i][j] = 4*parentD[i-n/2][j-n/2] + 1
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
