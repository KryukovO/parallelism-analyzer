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

// Функция, читающая исходный файл изображения
func readSourceImg(srcImgPath string) (img image.Image, errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			img = nil
			errF = fmt.Errorf("%v", msg)
		}
	}()

	srcImgReader, err := os.Open(srcImgPath)
	if err != nil {
		return nil, err
	}
	defer srcImgReader.Close()
	m, _, err := image.Decode(srcImgReader)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Функция, реализующая алгоритм усечения по порогу
func ThresholdDithering(srcImgPath string, dstImgPath string, threshold int) (errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			errF = fmt.Errorf("%v", msg)
		}
	}()

	if threshold < 0 || threshold > 255 {
		return errors.New("пороговое значение должно быть в пределах 0-255")
	}

	srcImg, err := readSourceImg(srcImgPath)
	if err != nil {
		return err
	}
	bounds := srcImg.Bounds()

	dstImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := srcImg.At(x, y).RGBA()
			intens := int(math.Max(float64(r/256), math.Max(float64(g/256), float64(b/256))))
			if intens > threshold {
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

	err = png.Encode(dstImgFile, dstImg)
	if err != nil {
		return err
	}

	return nil
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

	srcImg, err := readSourceImg(srcImgPath)
	if err != nil {
		return err
	}
	bounds := srcImg.Bounds()

	dstImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i := (y - bounds.Min.Y) % order
			j := (x - bounds.Min.X) % order
			r, g, b, a := srcImg.At(x, y).RGBA()
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

	err = png.Encode(dstImgFile, dstImg)
	if err != nil {
		return err
	}

	return nil
}

// Функция, реализующая алгоритм рассеивания ошибок Флойда-Стейнберга
func FloydErrDithering(srcImgPath string, dstImgPath string, threshold int) (errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			errF = fmt.Errorf("%v", msg)
		}
	}()

	if threshold < 0 || threshold > 255 {
		return errors.New("пороговое значение должно быть в пределах 0-255")
	}

	srcImg, err := readSourceImg(srcImgPath)
	if err != nil {
		return err
	}
	bounds := srcImg.Bounds()

	intensMatrix := make([][]uint8, bounds.Max.Y-bounds.Min.Y) // Двумерный слайс для хранения рассчитываемой интенсивности пикселей нового изображения
	for y := 0; y < bounds.Max.Y-bounds.Min.Y; y++ {
		intensMatrix[y] = make([]uint8, bounds.Max.X-bounds.Min.X)
		for x := 0; x < bounds.Max.X-bounds.Min.X; x++ {
			r, g, b, _ := srcImg.At(x, y).RGBA()
			intensMatrix[y][x] = uint8(math.Max(float64(r/256), math.Max(float64(g/256), float64(b/256))))
		}
	}

	for y := 0; y < bounds.Max.Y-bounds.Min.Y; y++ {
		for x := 0; x < bounds.Max.X-bounds.Min.X; x++ {
			intens := intensMatrix[y][x] // запоминаем интенсивность
			errPix := 0                  // распространяемая ошибка
			if intens > uint8(threshold) {
				intensMatrix[y][x] = 255
				errPix = int(intens) - 255
			} else {
				intensMatrix[y][x] = 0
				errPix = int(intens)
			}

			if x < bounds.Max.X-bounds.Min.X-1 {
				// Распространяем ошибку на пиксель справа
				intensMatrix[y][x+1] = uint8(math.Min(math.Max(float64(intensMatrix[y][x+1])+7.0/16.0*float64(errPix), 0), 255)) // 0 <= интенсивность <= 255
			}

			if y < bounds.Max.Y-bounds.Min.Y-1 {
				// Распространяем ошибку на снизу
				intensMatrix[y+1][x] = uint8(math.Min(math.Max(float64(intensMatrix[y+1][x])+5.0/16.0*float64(errPix), 0), 255)) // 0 <= интенсивность <= 255

				if x < bounds.Max.X-bounds.Min.X-1 {
					// Распространяем ошибку на пиксель справа снизу
					intensMatrix[y+1][x+1] = uint8(math.Min(math.Max(float64(intensMatrix[y+1][x+1])+1.0/16.0*float64(errPix), 0), 255)) // 0 <= интенсивность <= 255
				}

				if x > 0 {
					// Распространяем ошибку на снизу слева
					intensMatrix[y+1][x-1] = uint8(math.Min(math.Max(float64(intensMatrix[y+1][x-1])+3.0/16.0*float64(errPix), 0), 255)) // 0 <= интенсивность <= 255
				}
			}
		}
	}

	dstImg := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y-bounds.Min.Y; y++ {
		for x := 0; x < bounds.Max.X-bounds.Min.X; x++ {
			intens := intensMatrix[y][x]
			_, _, _, a := srcImg.At(bounds.Min.X+x, bounds.Min.Y+y).RGBA()
			dstImg.Set(x, y, color.RGBA{intens, intens, intens, uint8(a / 256)})
		}
	}

	dstImgFile, err := os.Create(dstImgPath)
	if err != nil {
		return err
	}

	err = png.Encode(dstImgFile, dstImg)
	if err != nil {
		return err
	}

	return nil
}
